package cimgui

/*
#include <stdlib.h>
#include "cimgui.h"

// The exported Go trampolines, forward-declared so the helpers below can take
// their addresses. cgo forbids defining these helpers in a file that also uses
// //export, so they live here, away from callbacks.go.
extern int goInputTextCallbackTrampoline(ImGuiInputTextCallbackData *data);
extern char *goStrGetterTrampoline(void *user_data, int idx);
extern float goFloatGetterTrampoline(void *user_data, int idx);

static inline ImGuiInputTextCallback cimguiInputTextCallbackPtr(void) {
	return goInputTextCallbackTrampoline;
}

static inline const char *(*cimguiStrGetterPtr(void))(void *, int) {
	return (const char *(*)(void *, int))goStrGetterTrampoline;
}

static inline float (*cimguiFloatGetterPtr(void))(void *, int) {
	return goFloatGetterTrampoline;
}
*/
import "C"

import (
	"unsafe"

	"github.com/bitwizeshift/go-imgui/internal/handle"
)

// InputTextResizable edits buf, growing it automatically as text is entered, and
// reports whether it changed. cb, when non-nil, receives any callback events
// enabled in flags (CallbackResize is always enabled internally).
func InputTextResizable(label string, buf *TextBuffer, flags InputTextFlags, cb InputTextCallback) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	token := handle.Save(&inputTextBinding{buf: buf, callback: cb})
	defer handle.Delete(token)
	flags |= InputTextFlagsCallbackResize
	return bool(C.igInputText(clabel, (*C.char)(buf.data), C.size_t(buf.size),
		C.ImGuiInputTextFlags(flags), C.cimguiInputTextCallbackPtr(), token))
}

// InputTextMultilineResizable is [InputTextResizable] in a multi-line box of the
// given size.
func InputTextMultilineResizable(label string, buf *TextBuffer, size Vec2, flags InputTextFlags, cb InputTextCallback) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	token := handle.Save(&inputTextBinding{buf: buf, callback: cb})
	defer handle.Delete(token)
	flags |= InputTextFlagsCallbackResize
	return bool(C.igInputTextMultiline(clabel, (*C.char)(buf.data), C.size_t(buf.size), size.c(),
		C.ImGuiInputTextFlags(flags), C.cimguiInputTextCallbackPtr(), token))
}

// InputTextWithHintResizable is [InputTextResizable] showing hint when empty.
func InputTextWithHintResizable(label, hint string, buf *TextBuffer, flags InputTextFlags, cb InputTextCallback) bool {
	clabel, chint := C.CString(label), C.CString(hint)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(chint))
	token := handle.Save(&inputTextBinding{buf: buf, callback: cb})
	defer handle.Delete(token)
	flags |= InputTextFlagsCallbackResize
	return bool(C.igInputTextWithHint(clabel, chint, (*C.char)(buf.data), C.size_t(buf.size),
		C.ImGuiInputTextFlags(flags), C.cimguiInputTextCallbackPtr(), token))
}

// Combo_FnStrPtr draws a combo box whose itemsCount labels are produced lazily by
// getter, updating currentItem and reporting whether it changed. A negative
// popupMaxHeightInItems uses the default.
func Combo_FnStrPtr(label string, currentItem *int32, getter func(idx int32) string, itemsCount, popupMaxHeightInItems int32) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	binding := &strGetterBinding{getter: getter}
	token := handle.Save(binding)
	defer handle.Delete(token)
	defer binding.free()
	ci := C.int(*currentItem)
	ret := bool(C.igCombo_FnStrPtr(clabel, &ci, C.cimguiStrGetterPtr(), token, C.int(itemsCount), C.int(popupMaxHeightInItems)))
	*currentItem = int32(ci)
	return ret
}

// ListBox_FnStrPtr draws a list box whose itemsCount labels are produced lazily by
// getter, updating currentItem and reporting whether it changed. A negative
// heightInItems uses the default.
func ListBox_FnStrPtr(label string, currentItem *int32, getter func(idx int32) string, itemsCount, heightInItems int32) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	binding := &strGetterBinding{getter: getter}
	token := handle.Save(binding)
	defer handle.Delete(token)
	defer binding.free()
	ci := C.int(*currentItem)
	ret := bool(C.igListBox_FnStrPtr(clabel, &ci, C.cimguiStrGetterPtr(), token, C.int(itemsCount), C.int(heightInItems)))
	*currentItem = int32(ci)
	return ret
}

// PlotLines_FnFloatPtr draws a line plot of valuesCount samples produced lazily by
// getter. See [PlotLines_FloatPtr] for the remaining parameters.
func PlotLines_FnFloatPtr(label string, getter func(idx int32) float32, valuesCount, valuesOffset int32, overlayText string, scaleMin, scaleMax float32, graphSize Vec2) {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	var coverlay *C.char
	if overlayText != "" {
		coverlay = C.CString(overlayText)
		defer C.free(unsafe.Pointer(coverlay))
	}
	token := handle.Save(floatGetter(getter))
	defer handle.Delete(token)
	C.igPlotLines_FnFloatPtr(clabel, C.cimguiFloatGetterPtr(), token, C.int(valuesCount), C.int(valuesOffset),
		coverlay, C.float(scaleMin), C.float(scaleMax), graphSize.c())
}

// PlotHistogram_FnFloatPtr draws a histogram of valuesCount samples produced lazily
// by getter. See [PlotLines_FloatPtr] for the remaining parameters.
func PlotHistogram_FnFloatPtr(label string, getter func(idx int32) float32, valuesCount, valuesOffset int32, overlayText string, scaleMin, scaleMax float32, graphSize Vec2) {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	var coverlay *C.char
	if overlayText != "" {
		coverlay = C.CString(overlayText)
		defer C.free(unsafe.Pointer(coverlay))
	}
	token := handle.Save(floatGetter(getter))
	defer handle.Delete(token)
	C.igPlotHistogram_FnFloatPtr(clabel, C.cimguiFloatGetterPtr(), token, C.int(valuesCount), C.int(valuesOffset),
		coverlay, C.float(scaleMin), C.float(scaleMax), graphSize.c())
}
