// package app runs a Dear ImGui application: it owns the GLFW window, the
// OpenGL3 + GLFW backends, the Dear ImGui context, and the per-frame loop, so
// callers only write per-frame UI code.
//
// The window system and Dear ImGui must run on the process main thread (a hard
// requirement on macOS). This package handles that with runtime.LockOSThread,
// so callers must not; the only expectation is that [Run] is called from the
// program's main goroutine (i.e. from main).
package app
