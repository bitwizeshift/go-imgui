#include "imgui.h"

#include "imgui_impl_glfw.h"
#include "imgui_impl_opengl3.h"

#include <GLFW/glfw3.h>

#if defined(__APPLE__)
#include <OpenGL/gl3.h>
#elif defined(_WIN32)
#include <windows.h>
#include <GL/gl.h>
#else
#include <GL/gl.h>
#endif

#include "backend.h"

extern "C" {

void backendGLFWHints(void) {
	glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 3);
	glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 2);
	glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE);
#if defined(__APPLE__)
	glfwWindowHint(GLFW_OPENGL_FORWARD_COMPAT, GLFW_TRUE);
#endif
}

bool backendInit(GLFWwindow *window) {
	if (!ImGui_ImplGlfw_InitForOpenGL(window, true)) {
		return false;
	}
#if defined(__APPLE__)
	const char *glsl_version = "#version 150";
#else
	const char *glsl_version = "#version 130";
#endif
	return ImGui_ImplOpenGL3_Init(glsl_version);
}

void backendNewFrame(void) {
	ImGui_ImplOpenGL3_NewFrame();
	ImGui_ImplGlfw_NewFrame();
}

void backendRender(int fb_width, int fb_height, float r, float g, float b, float a) {
	glViewport(0, 0, fb_width, fb_height);
	glClearColor(r, g, b, a);
	glClear(GL_COLOR_BUFFER_BIT);
	ImGui_ImplOpenGL3_RenderDrawData(ImGui::GetDrawData());
}

void backendShutdown(void) {
	ImGui_ImplOpenGL3_Shutdown();
	ImGui_ImplGlfw_Shutdown();
}

} // extern "C"
