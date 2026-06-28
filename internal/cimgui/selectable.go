package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// SelectableFlags configures selectables. Mirrors the public ImGuiSelectableFlags_.
type SelectableFlags int32

const (
	SelectableFlagsNone              SelectableFlags = C.ImGuiSelectableFlags_None
	SelectableFlagsNoAutoClosePopups SelectableFlags = C.ImGuiSelectableFlags_NoAutoClosePopups
	SelectableFlagsSpanAllColumns    SelectableFlags = C.ImGuiSelectableFlags_SpanAllColumns
	SelectableFlagsAllowDoubleClick  SelectableFlags = C.ImGuiSelectableFlags_AllowDoubleClick
	SelectableFlagsDisabled          SelectableFlags = C.ImGuiSelectableFlags_Disabled
	SelectableFlagsAllowOverlap      SelectableFlags = C.ImGuiSelectableFlags_AllowOverlap
	SelectableFlagsHighlight         SelectableFlags = C.ImGuiSelectableFlags_Highlight
)

// Selectable_Bool draws a selectable item and reports whether it was clicked. A
// zero size fits the label.
func Selectable_Bool(label string, selected bool, flags SelectableFlags, size Vec2) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igSelectable_Bool(clabel, C.bool(selected), C.ImGuiSelectableFlags(flags), size.c()))
}

// Selectable_BoolPtr draws a selectable item bound to pSelected (toggled on
// click) and reports whether it was clicked.
func Selectable_BoolPtr(label string, pSelected *bool, flags SelectableFlags, size Vec2) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	var ret C.bool
	withBoolPtr(pSelected, func(p *C.bool) {
		ret = C.igSelectable_BoolPtr(clabel, p, C.ImGuiSelectableFlags(flags), size.c())
	})
	return bool(ret)
}
