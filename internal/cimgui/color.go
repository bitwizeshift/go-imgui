package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// ColorEditFlags configures color editors and pickers. Mirrors the public
// ImGuiColorEditFlags_.
type ColorEditFlags int32

const (
	ColorEditFlagsNone             ColorEditFlags = C.ImGuiColorEditFlags_None
	ColorEditFlagsNoAlpha          ColorEditFlags = C.ImGuiColorEditFlags_NoAlpha
	ColorEditFlagsNoPicker         ColorEditFlags = C.ImGuiColorEditFlags_NoPicker
	ColorEditFlagsNoOptions        ColorEditFlags = C.ImGuiColorEditFlags_NoOptions
	ColorEditFlagsNoSmallPreview   ColorEditFlags = C.ImGuiColorEditFlags_NoSmallPreview
	ColorEditFlagsNoInputs         ColorEditFlags = C.ImGuiColorEditFlags_NoInputs
	ColorEditFlagsNoTooltip        ColorEditFlags = C.ImGuiColorEditFlags_NoTooltip
	ColorEditFlagsNoLabel          ColorEditFlags = C.ImGuiColorEditFlags_NoLabel
	ColorEditFlagsNoSidePreview    ColorEditFlags = C.ImGuiColorEditFlags_NoSidePreview
	ColorEditFlagsNoDragDrop       ColorEditFlags = C.ImGuiColorEditFlags_NoDragDrop
	ColorEditFlagsNoBorder         ColorEditFlags = C.ImGuiColorEditFlags_NoBorder
	ColorEditFlagsAlphaOpaque      ColorEditFlags = C.ImGuiColorEditFlags_AlphaOpaque
	ColorEditFlagsAlphaNoBg        ColorEditFlags = C.ImGuiColorEditFlags_AlphaNoBg
	ColorEditFlagsAlphaPreviewHalf ColorEditFlags = C.ImGuiColorEditFlags_AlphaPreviewHalf
	ColorEditFlagsAlphaBar         ColorEditFlags = C.ImGuiColorEditFlags_AlphaBar
	ColorEditFlagsHDR              ColorEditFlags = C.ImGuiColorEditFlags_HDR
	ColorEditFlagsDisplayRGB       ColorEditFlags = C.ImGuiColorEditFlags_DisplayRGB
	ColorEditFlagsDisplayHSV       ColorEditFlags = C.ImGuiColorEditFlags_DisplayHSV
	ColorEditFlagsDisplayHex       ColorEditFlags = C.ImGuiColorEditFlags_DisplayHex
	ColorEditFlagsUint8            ColorEditFlags = C.ImGuiColorEditFlags_Uint8
	ColorEditFlagsFloat            ColorEditFlags = C.ImGuiColorEditFlags_Float
	ColorEditFlagsPickerHueBar     ColorEditFlags = C.ImGuiColorEditFlags_PickerHueBar
	ColorEditFlagsPickerHueWheel   ColorEditFlags = C.ImGuiColorEditFlags_PickerHueWheel
	ColorEditFlagsInputRGB         ColorEditFlags = C.ImGuiColorEditFlags_InputRGB
	ColorEditFlagsInputHSV         ColorEditFlags = C.ImGuiColorEditFlags_InputHSV
)

// ColorEdit3 edits an RGB color in place and reports whether it changed.
func ColorEdit3(label string, col *[3]float32, flags ColorEditFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igColorEdit3(clabel, (*C.float)(&col[0]), C.ImGuiColorEditFlags(flags)))
}

// ColorEdit4 edits an RGBA color in place and reports whether it changed.
func ColorEdit4(label string, col *[4]float32, flags ColorEditFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igColorEdit4(clabel, (*C.float)(&col[0]), C.ImGuiColorEditFlags(flags)))
}

// ColorPicker3 shows an RGB color picker editing col in place.
func ColorPicker3(label string, col *[3]float32, flags ColorEditFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igColorPicker3(clabel, (*C.float)(&col[0]), C.ImGuiColorEditFlags(flags)))
}

// ColorPicker4 shows an RGBA color picker editing col in place. refCol may be nil.
func ColorPicker4(label string, col *[4]float32, flags ColorEditFlags, refCol *[4]float32) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	var cref *C.float
	if refCol != nil {
		cref = (*C.float)(&refCol[0])
	}
	return bool(C.igColorPicker4(clabel, (*C.float)(&col[0]), C.ImGuiColorEditFlags(flags), cref))
}

// ColorButton draws a color swatch button and reports whether it was clicked.
func ColorButton(descID string, col Vec4, flags ColorEditFlags, size Vec2) bool {
	cid := C.CString(descID)
	defer C.free(unsafe.Pointer(cid))
	return bool(C.igColorButton(cid, col.c(), C.ImGuiColorEditFlags(flags), size.c()))
}
