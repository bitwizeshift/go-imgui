package app

import (
	"errors"
	"runtime"

	"github.com/bitwizeshift/go-imgui"
	"github.com/bitwizeshift/go-imgui/internal/cimgui"
)

func init() {
	// Lock the main goroutine to the main OS thread as early as possible so the
	// window system and Dear ImGui run there. This runs before main.
	runtime.LockOSThread()
}

// Errors returned by [Run].
var (
	ErrGLFWInit     = errors.New("glfw: initialization failed")
	ErrWindowCreate = errors.New("glfw: window creation failed")
	ErrBackendInit  = errors.New("imgui: backend initialization failed")
)

// Config configures the application window. The zero value is usable; empty
// fields fall back to sensible defaults.
type Config struct {
	Title         string
	Width, Height int
	ClearColor    imgui.Vec4
}

func (c Config) withDefaults() Config {
	if c.Title == "" {
		c.Title = "go-imgui"
	}
	if c.Width <= 0 {
		c.Width = 1280
	}
	if c.Height <= 0 {
		c.Height = 720
	}
	if c.ClearColor == (imgui.Vec4{}) {
		c.ClearColor = imgui.Vec4{X: 0.10, Y: 0.10, Z: 0.12, W: 1.0}
	}
	return c
}

// Run opens a window per cfg and drives the frame loop until the window is
// closed, invoking frame once per frame between NewFrame and Render. The frame
// callback issues imgui widget calls. Run must be called from the program's main
// goroutine. It returns [ErrGLFWInit], [ErrWindowCreate], or [ErrBackendInit] on
// startup failure, and nil after the window closes.
func Run(cfg Config, frame func()) error {
	cfg = cfg.withDefaults()
	runtime.LockOSThread()

	if !cimgui.GLFWInit() {
		return ErrGLFWInit
	}
	defer cimgui.GLFWTerminate()

	cimgui.GLFWDefaultWindowHints()
	cimgui.GLFWApplyGLHints()

	win := cimgui.CreateWindow(cfg.Width, cfg.Height, cfg.Title)
	if win == nil {
		return ErrWindowCreate
	}
	defer cimgui.DestroyWindow(win)

	cimgui.MakeContextCurrent(win)
	cimgui.SwapInterval(1)

	ctx := imgui.CreateContext()
	defer ctx.Destroy()
	imgui.StyleColorsDark()

	if !cimgui.BackendInit(win) {
		return ErrBackendInit
	}
	defer cimgui.BackendShutdown()

	clear := cfg.ClearColor
	for !cimgui.WindowShouldClose(win) {
		cimgui.PollEvents()

		cimgui.BackendNewFrame()
		imgui.NewFrame()

		frame()

		imgui.Render()
		w, h := cimgui.FramebufferSize(win)
		cimgui.BackendRender(w, h, clear.X, clear.Y, clear.Z, clear.W)
		cimgui.SwapBuffers(win)
	}
	return nil
}
