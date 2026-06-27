package cimgui

// #include <stdlib.h>
// #include <GLFW/glfw3.h>
// #include "backend.h"
import "C"

import "unsafe"

// Window is an opaque handle to a GLFW window.
type Window = unsafe.Pointer

// GLFWInit initializes GLFW. Returns false on failure.
func GLFWInit() bool { return C.glfwInit() != 0 }

// GLFWTerminate shuts GLFW down.
func GLFWTerminate() { C.glfwTerminate() }

// GLFWDefaultWindowHints resets window creation hints to their defaults.
func GLFWDefaultWindowHints() { C.glfwDefaultWindowHints() }

// GLFWApplyGLHints requests the OpenGL context version and profile ImGui needs.
// Call after [GLFWInit] and before [CreateWindow].
func GLFWApplyGLHints() { C.backendGLFWHints() }

// CreateWindow creates a window and its OpenGL context. Returns nil on failure.
func CreateWindow(width, height int, title string) Window {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	return Window(C.glfwCreateWindow(C.int(width), C.int(height), ctitle, nil, nil))
}

// DestroyWindow destroys a window created by [CreateWindow].
func DestroyWindow(win Window) { C.glfwDestroyWindow((*C.GLFWwindow)(win)) }

// MakeContextCurrent makes the window's GL context current on the calling thread.
func MakeContextCurrent(win Window) { C.glfwMakeContextCurrent((*C.GLFWwindow)(win)) }

// SwapInterval sets the buffer swap interval (1 enables vsync).
func SwapInterval(interval int) { C.glfwSwapInterval(C.int(interval)) }

// WindowShouldClose reports whether the window has been asked to close.
func WindowShouldClose(win Window) bool {
	return C.glfwWindowShouldClose((*C.GLFWwindow)(win)) != 0
}

// PollEvents processes pending window-system events.
func PollEvents() { C.glfwPollEvents() }

// SwapBuffers presents the back buffer.
func SwapBuffers(win Window) { C.glfwSwapBuffers((*C.GLFWwindow)(win)) }

// FramebufferSize returns the window's framebuffer size in pixels.
func FramebufferSize(win Window) (width, height int) {
	var w, h C.int
	C.glfwGetFramebufferSize((*C.GLFWwindow)(win), &w, &h)
	return int(w), int(h)
}
