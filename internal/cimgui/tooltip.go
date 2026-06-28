package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
// #include "shims.h"
import "C"

import "unsafe"

// BeginTooltip opens a tooltip window. Call [EndTooltip] only if it returns true.
func BeginTooltip() bool { return bool(C.igBeginTooltip()) }

// BeginItemTooltip opens a tooltip only when the previous item is hovered. Call
// [EndTooltip] only if it returns true.
func BeginItemTooltip() bool { return bool(C.igBeginItemTooltip()) }

// EndTooltip closes the tooltip opened by [BeginTooltip] or [BeginItemTooltip].
func EndTooltip() { C.igEndTooltip() }

// SetTooltip sets the contents of a tooltip shown while the previous item is
// hovered.
func SetTooltip(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.shimSetTooltip(ctext)
}
