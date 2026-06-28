package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// Context is an opaque handle to a Dear ImGui context.
type Context = unsafe.Pointer

// Vec2 mirrors Dear ImGui's ImVec2.
type Vec2 struct {
	X, Y float32
}

// c converts v to the cimgui ImVec2_c value type.
func (v Vec2) c() C.ImVec2_c {
	return C.ImVec2_c{x: C.float(v.X), y: C.float(v.Y)}
}

// Vec4 mirrors Dear ImGui's ImVec4 (also used for RGBA colors).
type Vec4 struct {
	X, Y, Z, W float32
}

// c converts v to the cimgui ImVec4_c value type.
func (v Vec4) c() C.ImVec4_c {
	return C.ImVec4_c{x: C.float(v.X), y: C.float(v.Y), z: C.float(v.Z), w: C.float(v.W)}
}

// withBoolPtr runs fn with a C bool pointer reflecting p, copying any change
// back into p. A nil p yields a nil C pointer.
func withBoolPtr(p *bool, fn func(*C.bool)) {
	if p == nil {
		fn(nil)
		return
	}
	v := C.bool(*p)
	fn(&v)
	*p = bool(v)
}
