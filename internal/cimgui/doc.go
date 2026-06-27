// package cimgui is the single cgo boundary for the Dear ImGui wrapper.
//
// It is the only package in the module that uses `import "C"`. It compiles, from
// the vendored sources copied into csources/ (see gencsources.go): the cimgui C
// API, its bundled Dear ImGui, the GLFW platform backend, the OpenGL3 renderer
// backend, and GLFW itself. On top of that it exposes Go functions that map 1:1
// onto the cimgui C API with Go types substituted for C ones.
//
// It is deliberately internal and intentionally un-ergonomic: the curated,
// idiomatic surface lives in the public packages that import this one. Overloaded
// ImGui functions are bound to a single chosen overload here (for example
// [RadioButton] wraps cimgui's igRadioButton_Bool).
//
// CIMGUI_DEFINE_ENUMS_AND_STRUCTS is defined for C compilation (the cgo preamble)
// so cimgui.h exposes the C structs and enums, but not for the C++ translation
// units, where cimgui.cpp must see the real Dear ImGui types.
package cimgui
