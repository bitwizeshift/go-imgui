package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// ComboFlags configures a combo box opened with [BeginCombo]. Mirrors the public
// ImGuiComboFlags_.
type ComboFlags int32

const (
	ComboFlagsNone           ComboFlags = C.ImGuiComboFlags_None
	ComboFlagsPopupAlignLeft ComboFlags = C.ImGuiComboFlags_PopupAlignLeft
	ComboFlagsHeightSmall    ComboFlags = C.ImGuiComboFlags_HeightSmall
	ComboFlagsHeightRegular  ComboFlags = C.ImGuiComboFlags_HeightRegular
	ComboFlagsHeightLarge    ComboFlags = C.ImGuiComboFlags_HeightLarge
	ComboFlagsHeightLargest  ComboFlags = C.ImGuiComboFlags_HeightLargest
	ComboFlagsNoArrowButton  ComboFlags = C.ImGuiComboFlags_NoArrowButton
	ComboFlagsNoPreview      ComboFlags = C.ImGuiComboFlags_NoPreview
	ComboFlagsWidthFitPreview ComboFlags = C.ImGuiComboFlags_WidthFitPreview
)

// cStringArray builds a C array of NUL-terminated strings from items. The
// returned free releases every allocation; call it after the C call returns.
func cStringArray(items []string) (arr **C.char, free func()) {
	if len(items) == 0 {
		return nil, func() {}
	}
	n := len(items)
	block := C.malloc(C.size_t(uintptr(n) * unsafe.Sizeof((*C.char)(nil))))
	view := unsafe.Slice((**C.char)(block), n)
	cs := make([]*C.char, n)
	for i, s := range items {
		cs[i] = C.CString(s)
		view[i] = cs[i]
	}
	return (**C.char)(block), func() {
		for _, p := range cs {
			C.free(unsafe.Pointer(p))
		}
		C.free(block)
	}
}

// BeginCombo opens a combo box. Call [EndCombo] only if it returns true.
func BeginCombo(label, previewValue string, flags ComboFlags) bool {
	clabel, cpreview := C.CString(label), C.CString(previewValue)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cpreview))
	return bool(C.igBeginCombo(clabel, cpreview, C.ImGuiComboFlags(flags)))
}

// EndCombo closes the combo box opened by [BeginCombo].
func EndCombo() { C.igEndCombo() }

// Combo_Str_arr draws a combo box selecting an index within items, updating
// currentItem and reporting whether it changed. A negative popupMaxHeightInItems
// uses the default.
func Combo_Str_arr(label string, currentItem *int32, items []string, popupMaxHeightInItems int32) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	arr, free := cStringArray(items)
	defer free()
	ci := C.int(*currentItem)
	ret := bool(C.igCombo_Str_arr(clabel, &ci, arr, C.int(len(items)), C.int(popupMaxHeightInItems)))
	*currentItem = int32(ci)
	return ret
}

// Combo_Str draws a combo box whose items come from a single string of
// NUL-separated, double-NUL-terminated entries (e.g. "a\x00b\x00c\x00").
func Combo_Str(label string, currentItem *int32, itemsSeparatedByZeros string, popupMaxHeightInItems int32) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	citems := C.CString(itemsSeparatedByZeros)
	defer C.free(unsafe.Pointer(citems))
	ci := C.int(*currentItem)
	ret := bool(C.igCombo_Str(clabel, &ci, citems, C.int(popupMaxHeightInItems)))
	*currentItem = int32(ci)
	return ret
}

// BeginListBox opens a scrolling list box of the given size. Call [EndListBox]
// only if it returns true.
func BeginListBox(label string, size Vec2) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igBeginListBox(clabel, size.c()))
}

// EndListBox closes the list box opened by [BeginListBox].
func EndListBox() { C.igEndListBox() }

// ListBox_Str_arr draws a list box selecting an index within items, updating
// currentItem and reporting whether it changed. A negative heightInItems uses the
// default.
func ListBox_Str_arr(label string, currentItem *int32, items []string, heightInItems int32) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	arr, free := cStringArray(items)
	defer free()
	ci := C.int(*currentItem)
	ret := bool(C.igListBox_Str_arr(clabel, &ci, arr, C.int(len(items)), C.int(heightInItems)))
	*currentItem = int32(ci)
	return ret
}
