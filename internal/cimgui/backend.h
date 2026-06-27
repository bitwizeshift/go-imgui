// Hand-written extern "C" wrappers for the GLFW platform backend and the
// OpenGL3 renderer backend. These keep all GL calls on the C side so Go never
// needs OpenGL bindings.
#ifndef CIMGUI_BACKEND_H
#define CIMGUI_BACKEND_H

#include <stdbool.h>

struct GLFWwindow;

#ifdef __cplusplus
extern "C" {
#endif

// backendGLFWHints requests the OpenGL context version/profile ImGui needs.
// Call after glfwInit and before glfwCreateWindow.
void backendGLFWHints(void);

// backendInit wires the GLFW + OpenGL3 ImGui backends to an existing window
// with a current GL context. Returns false on failure.
bool backendInit(struct GLFWwindow *window);

// backendNewFrame starts a backend frame (call before igNewFrame).
void backendNewFrame(void);

// backendRender clears the framebuffer and draws the current ImGui frame.
// Call after igRender.
void backendRender(int fb_width, int fb_height, float r, float g, float b, float a);

// backendShutdown tears down both backends.
void backendShutdown(void);

#ifdef __cplusplus
}
#endif

#endif // CIMGUI_BACKEND_H
