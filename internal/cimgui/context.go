package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

// CreateContext creates and activates a new Dear ImGui context.
func CreateContext() Context { return Context(C.igCreateContext(nil)) }

// DestroyContext destroys a context created by [CreateContext].
func DestroyContext(ctx Context) { C.igDestroyContext((*C.ImGuiContext)(ctx)) }

// NewFrame begins a new Dear ImGui frame.
func NewFrame() { C.igNewFrame() }

// EndFrame ends the current frame; usually called implicitly by [Render].
func EndFrame() { C.igEndFrame() }

// Render finalizes the current frame's draw data.
func Render() { C.igRender() }

// StyleColorsDark applies the built-in dark style.
func StyleColorsDark() { C.igStyleColorsDark(nil) }

// StyleColorsLight applies the built-in light style.
func StyleColorsLight() { C.igStyleColorsLight(nil) }

// StyleColorsClassic applies the built-in classic style.
func StyleColorsClassic() { C.igStyleColorsClassic(nil) }
