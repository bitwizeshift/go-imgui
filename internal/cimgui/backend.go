package cimgui

// #include <GLFW/glfw3.h>
// #include "backend.h"
import "C"

// BackendInit wires the GLFW + OpenGL3 ImGui backends to win. Returns false on
// failure.
func BackendInit(win Window) bool { return bool(C.backendInit((*C.GLFWwindow)(win))) }

// BackendNewFrame starts a backend frame; call before [NewFrame].
func BackendNewFrame() { C.backendNewFrame() }

// BackendRender clears to the given color and draws the current frame; call
// after [Render].
func BackendRender(fbWidth, fbHeight int, r, g, b, a float32) {
	C.backendRender(C.int(fbWidth), C.int(fbHeight), C.float(r), C.float(g), C.float(b), C.float(a))
}

// BackendShutdown tears down both backends.
func BackendShutdown() { C.backendShutdown() }
