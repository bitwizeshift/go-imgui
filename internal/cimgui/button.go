package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// Dir is a cardinal direction. Mirrors ImGuiDir.
type Dir int32

const (
	DirNone  Dir = C.ImGuiDir_None
	DirLeft  Dir = C.ImGuiDir_Left
	DirRight Dir = C.ImGuiDir_Right
	DirUp    Dir = C.ImGuiDir_Up
	DirDown  Dir = C.ImGuiDir_Down
)

// ButtonFlags configures [InvisibleButton]. Mirrors the public ImGuiButtonFlags_.
type ButtonFlags int32

const (
	ButtonFlagsNone              ButtonFlags = C.ImGuiButtonFlags_None
	ButtonFlagsMouseButtonLeft   ButtonFlags = C.ImGuiButtonFlags_MouseButtonLeft
	ButtonFlagsMouseButtonRight  ButtonFlags = C.ImGuiButtonFlags_MouseButtonRight
	ButtonFlagsMouseButtonMiddle ButtonFlags = C.ImGuiButtonFlags_MouseButtonMiddle
	ButtonFlagsEnableNav         ButtonFlags = C.ImGuiButtonFlags_EnableNav
)

// Button draws a button and reports whether it was clicked. A zero size
// auto-fits the label.
func Button(label string, size Vec2) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igButton(clabel, size.c()))
}

// SmallButton draws a button with no frame padding.
func SmallButton(label string) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igSmallButton(clabel))
}

// InvisibleButton draws a sizeable behavior-only button with no visuals.
func InvisibleButton(strID string, size Vec2, flags ButtonFlags) bool {
	cid := C.CString(strID)
	defer C.free(unsafe.Pointer(cid))
	return bool(C.igInvisibleButton(cid, size.c(), C.ImGuiButtonFlags(flags)))
}

// ArrowButton draws a square button containing an arrow in the given direction.
func ArrowButton(strID string, dir Dir) bool {
	cid := C.CString(strID)
	defer C.free(unsafe.Pointer(cid))
	return bool(C.igArrowButton(cid, C.ImGuiDir(dir)))
}

// Bullet draws a small bullet and advances the cursor onto the same line.
func Bullet() { C.igBullet() }

// Checkbox draws a checkbox bound to v and reports whether it changed.
func Checkbox(label string, v *bool) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cv := C.bool(*v)
	ret := bool(C.igCheckbox(clabel, &cv))
	*v = bool(cv)
	return ret
}

// CheckboxFlags_IntPtr toggles flagsValue within flags and reports whether it changed.
func CheckboxFlags_IntPtr(label string, flags *int32, flagsValue int32) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cf := C.int(*flags)
	ret := bool(C.igCheckboxFlags_IntPtr(clabel, &cf, C.int(flagsValue)))
	*flags = int32(cf)
	return ret
}

// CheckboxFlags_UintPtr toggles flagsValue within flags and reports whether it changed.
func CheckboxFlags_UintPtr(label string, flags *uint32, flagsValue uint32) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cf := C.uint(*flags)
	ret := bool(C.igCheckboxFlags_UintPtr(clabel, &cf, C.uint(flagsValue)))
	*flags = uint32(cf)
	return ret
}

// RadioButton_Bool draws a radio button and reports whether it was clicked.
func RadioButton_Bool(label string, active bool) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igRadioButton_Bool(clabel, C.bool(active)))
}

// RadioButton_IntPtr draws a radio button that sets v to vButton when clicked,
// and reports whether it changed.
func RadioButton_IntPtr(label string, v *int32, vButton int32) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cv := C.int(*v)
	ret := bool(C.igRadioButton_IntPtr(clabel, &cv, C.int(vButton)))
	*v = int32(cv)
	return ret
}

// ProgressBar draws a progress bar filled to fraction (0..1). A zero size
// auto-fits; overlay, when non-empty, is drawn centered over the bar.
func ProgressBar(fraction float32, size Vec2, overlay string) {
	var coverlay *C.char
	if overlay != "" {
		coverlay = C.CString(overlay)
		defer C.free(unsafe.Pointer(coverlay))
	}
	C.igProgressBar(C.float(fraction), size.c(), coverlay)
}
