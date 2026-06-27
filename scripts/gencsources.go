//go:build ignore

// Command gencsources copies and "massages" the vendored cimgui (the generated
// C API over Dear ImGui), its bundled Dear ImGui, the GLFW and OpenGL3
// backends, and GLFW itself out of the git submodules in vendor/ into
// internal/cimgui/csources/ so that cgo can compile them from source.
//
// It also (re)writes the small "unity" translation units in internal/cimgui/
// that #include those copied sources, plus the synthesized glfw_config.h that
// GLFW's internal.h hard-requires.
//
// The output is committed to the repository: consumers build it without the
// submodules and without running this generator. Run it only when bumping the
// vendored versions:
//
//	go generate ./...
//
// This file uses only the standard library on purpose: nothing it emits, and
// nothing it imports, may leak into the published module.
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root, err := repoRoot()
	if err != nil {
		fatal(err)
	}

	var (
		cimguiSrc = filepath.Join(root, "vendor", "cimgui")
		imguiSrc  = filepath.Join(cimguiSrc, "imgui")
		glfwSrc   = filepath.Join(root, "vendor", "glfw")
		pkgDir    = filepath.Join(root, "internal", "cimgui")
		dst       = filepath.Join(pkgDir, "csources")
	)

	if err := assertSubmodule(cimguiSrc, "cimgui.h"); err != nil {
		fatal(err)
	}
	if err := assertSubmodule(imguiSrc, "imgui.h"); err != nil {
		fatal(err)
	}
	if err := assertSubmodule(glfwSrc, filepath.Join("include", "GLFW", "glfw3.h")); err != nil {
		fatal(err)
	}

	// Start from a clean csources/ so removed upstream files don't linger.
	if err := os.RemoveAll(dst); err != nil {
		fatal(err)
	}

	// 1. cimgui's bundled Dear ImGui. The directory structure is preserved
	//    under csources/cimgui/imgui so cimgui.cpp's relative "./imgui/..."
	//    includes resolve unchanged.
	for _, name := range imguiCoreFiles {
		copyInto(filepath.Join(imguiSrc, name), filepath.Join(dst, "cimgui", "imgui", name))
	}

	// imconfig.h is included near the top of imgui.h. Appending our user config
	// here applies compile-time overrides (imconfig_user.h lives in the package
	// dir, found via -I${SRCDIR}).
	appendFile(filepath.Join(dst, "cimgui", "imgui", "imconfig.h"),
		"\n// Appended by gencsources.go: pull in our local overrides.\n"+
			"#include \"imconfig_user.h\"\n")

	// 2. cimgui itself (the generated C API).
	for _, name := range cimguiFiles {
		copyInto(filepath.Join(cimguiSrc, name), filepath.Join(dst, "cimgui", name))
	}

	// 3. Backends we support: GLFW platform + OpenGL3 renderer.
	for _, name := range backendFiles {
		copyInto(filepath.Join(imguiSrc, "backends", name), filepath.Join(dst, "backends", name))
	}

	// 4. GLFW: preserve src/ + include/ structure so its relative #includes
	//    ("internal.h", "../include/GLFW/glfw3.h") resolve unchanged.
	copyTree(filepath.Join(glfwSrc, "src"), filepath.Join(dst, "glfw", "src"), isGLFWSource)
	copyTree(filepath.Join(glfwSrc, "include"), filepath.Join(dst, "glfw", "include"), isHeader)

	// GLFW's internal.h does `#include "glfw_config.h"`, a file CMake normally
	// generates. We select the platform via cgo CFLAGS instead, so an empty stub
	// is all that is required.
	writeFile(filepath.Join(dst, "glfw", "src", "glfw_config.h"), glfwConfigStub)

	// 5. Unity translation units in the package dir that pull in the copied
	//    sources. cgo compiles only files in the package dir (not csources/),
	//    so these are the actual compilation entry points.
	for name, content := range unityFiles {
		writeFile(filepath.Join(pkgDir, name), content)
	}

	fmt.Println("gencsources: wrote", rel(root, dst), "and unity translation units")
}

// imguiCoreFiles are the Dear ImGui sources compiled into unity_core.cpp.
// imgui_demo.cpp is included so ShowDemoWindow is available.
var imguiCoreFiles = []string{
	"imgui.h",
	"imgui_internal.h",
	"imconfig.h",
	"imstb_rectpack.h",
	"imstb_textedit.h",
	"imstb_truetype.h",
	"imgui.cpp",
	"imgui_draw.cpp",
	"imgui_tables.cpp",
	"imgui_widgets.cpp",
	"imgui_demo.cpp",
}

// cimguiFiles are the generated C API and its config header.
var cimguiFiles = []string{
	"cimgui.h",
	"cimgui.cpp",
	"cimconfig.h",
}

var backendFiles = []string{
	"imgui_impl_glfw.h",
	"imgui_impl_glfw.cpp",
	"imgui_impl_opengl3.h",
	"imgui_impl_opengl3.cpp",
	"imgui_impl_opengl3_loader.h",
}

// GLFW source files grouped the way GLFW's own CMakeLists.txt groups them.
// Shared files build on every platform; the platform groups are gated by the
// matching GOOS-suffixed unity file.
var (
	glfwCommon = []string{
		"context.c", "init.c", "input.c", "monitor.c", "platform.c",
		"vulkan.c", "window.c", "egl_context.c", "osmesa_context.c",
		"null_init.c", "null_monitor.c", "null_window.c", "null_joystick.c",
	}
	glfwCocoa = []string{
		"cocoa_init.m", "cocoa_joystick.m", "cocoa_monitor.m", "cocoa_window.m",
		"cocoa_time.c", "nsgl_context.m", "posix_module.c", "posix_thread.c",
	}
	glfwX11 = []string{
		"x11_init.c", "x11_monitor.c", "x11_window.c", "xkb_unicode.c",
		"glx_context.c", "posix_time.c", "posix_thread.c", "posix_module.c",
		"posix_poll.c", "linux_joystick.c",
	}
	glfwWin32 = []string{
		"win32_init.c", "win32_module.c", "win32_joystick.c", "win32_monitor.c",
		"win32_time.c", "win32_thread.c", "win32_window.c", "wgl_context.c",
	}
)

// unityFiles maps a file name in internal/cimgui/ to its generated contents.
var unityFiles = map[string]string{
	"unity_core.cpp": genHeader("//") + joinIncludes(
		"csources/cimgui/imgui/imgui.cpp",
		"csources/cimgui/imgui/imgui_draw.cpp",
		"csources/cimgui/imgui/imgui_tables.cpp",
		"csources/cimgui/imgui/imgui_widgets.cpp",
		"csources/cimgui/imgui/imgui_demo.cpp",
	),

	"unity_cimgui.cpp": genHeader("//") + joinIncludes(
		"csources/cimgui/cimgui.cpp",
	),

	"unity_backends.cpp": genHeader("//") + joinIncludes(
		"csources/backends/imgui_impl_glfw.cpp",
		"csources/backends/imgui_impl_opengl3.cpp",
	),

	"unity_glfw_common.c":  genHeader("//") + glfwIncludes(glfwCommon),
	"unity_glfw_darwin.m":  genHeader("//") + glfwIncludes(glfwCocoa),
	"unity_glfw_linux.c":   genHeader("//") + glfwIncludes(glfwX11),
	"unity_glfw_windows.c": genHeader("//") + glfwIncludes(glfwWin32),
}

const glfwConfigStub = `// Code generated by gencsources.go. DO NOT EDIT.
//
// GLFW's internal.h requires this file. The platform backend is selected via
// cgo CFLAGS (-D_GLFW_COCOA / -D_GLFW_X11 / -D_GLFW_WIN32), so nothing is
// defined here.
`

func glfwIncludes(files []string) string {
	paths := make([]string, len(files))
	for i, f := range files {
		paths[i] = "csources/glfw/src/" + f
	}
	return joinIncludes(paths...)
}

func joinIncludes(paths ...string) string {
	var b strings.Builder
	for _, p := range paths {
		fmt.Fprintf(&b, "#include %q\n", p)
	}
	return b.String()
}

func genHeader(comment string) string {
	return comment + " Code generated by gencsources.go. DO NOT EDIT.\n\n"
}

func isGLFWSource(name string) bool {
	switch filepath.Ext(name) {
	case ".c", ".m", ".h":
		return true
	default:
		return false
	}
}

func isHeader(name string) bool { return filepath.Ext(name) == ".h" }

// --- filesystem helpers ---------------------------------------------------

func copyTree(srcDir, dstDir string, keep func(name string) bool) {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		fatal(err)
	}
	for _, e := range entries {
		src := filepath.Join(srcDir, e.Name())
		dst := filepath.Join(dstDir, e.Name())
		if e.IsDir() {
			copyTree(src, dst, keep)
			continue
		}
		if keep(e.Name()) {
			copyInto(src, dst)
		}
	}
}

func copyInto(src, dst string) {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		fatal(err)
	}
	in, err := os.Open(src)
	if err != nil {
		fatal(fmt.Errorf("copy %s: %w", src, err))
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		fatal(err)
	}
	if _, err := io.Copy(out, in); err != nil {
		out.Close()
		fatal(err)
	}
	if err := out.Close(); err != nil {
		fatal(err)
	}
}

func appendFile(path, content string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		fatal(err)
	}
}

func writeFile(path, content string) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		fatal(err)
	}
}

// repoRoot walks up from the working directory until it finds go.work, so the
// generator behaves the same whether invoked via `go generate` from the module
// root or run directly from scripts/.
func repoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.work")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not locate go.work above %s", dir)
		}
		dir = parent
	}
}

func assertSubmodule(dir, sentinel string) error {
	if _, err := os.Stat(filepath.Join(dir, sentinel)); err != nil {
		return fmt.Errorf("missing %s in %s; run `git submodule update --init --recursive`: %w", sentinel, dir, err)
	}
	return nil
}

func rel(root, p string) string {
	if r, err := filepath.Rel(root, p); err == nil {
		return r
	}
	return p
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "gencsources:", err)
	os.Exit(1)
}
