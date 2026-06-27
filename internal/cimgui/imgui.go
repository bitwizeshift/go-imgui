package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// Context is an opaque handle to a Dear ImGui context.
type Context = unsafe.Pointer

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

// ---- Context / frame lifecycle -------------------------------------------

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

// ---- Windows -------------------------------------------------------------

// Begin opens a window. pOpen may be nil; when non-nil it is updated with the
// window's open state. Returns whether the window's contents are visible.
func Begin(name string, pOpen *bool, flags int) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var ret C.bool
	withBoolPtr(pOpen, func(p *C.bool) {
		ret = C.igBegin(cname, p, C.ImGuiWindowFlags(flags))
	})
	return bool(ret)
}

// End closes the window opened by the matching [Begin].
func End() { C.igEnd() }

// BeginChild opens a child region. A zero size fills the available space.
func BeginChild(strID string, sizeX, sizeY float32, childFlags, windowFlags int) bool {
	cid := C.CString(strID)
	defer C.free(unsafe.Pointer(cid))
	return bool(C.igBeginChild_Str(cid, vec2(sizeX, sizeY),
		C.ImGuiChildFlags(childFlags), C.ImGuiWindowFlags(windowFlags)))
}

// EndChild closes the child region opened by [BeginChild].
func EndChild() { C.igEndChild() }

// SetNextWindowSize sets the size of the next window.
func SetNextWindowSize(w, h float32, cond int) {
	C.igSetNextWindowSize(vec2(w, h), C.ImGuiCond(cond))
}

// SetNextWindowPos sets the position of the next window, with an optional pivot.
func SetNextWindowPos(x, y float32, cond int, pivotX, pivotY float32) {
	C.igSetNextWindowPos(vec2(x, y), C.ImGuiCond(cond), vec2(pivotX, pivotY))
}

// ---- Text ----------------------------------------------------------------

// TextUnformatted draws text verbatim (no printf formatting).
func TextUnformatted(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.igTextUnformatted(ctext, nil)
}

// SeparatorText draws a horizontal separator with a centered label.
func SeparatorText(label string) {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	C.igSeparatorText(clabel)
}

// ---- Layout --------------------------------------------------------------

// Separator draws a horizontal line.
func Separator() { C.igSeparator() }

// SameLine continues the current line. Pass zeros for default spacing.
func SameLine(offsetFromStartX, spacing float32) {
	C.igSameLine(C.float(offsetFromStartX), C.float(spacing))
}

// NewLine moves to the next line.
func NewLine() { C.igNewLine() }

// Spacing adds vertical spacing.
func Spacing() { C.igSpacing() }

// Dummy adds an empty item of the given size.
func Dummy(w, h float32) { C.igDummy(vec2(w, h)) }

// Indent increases the indent. A zero width uses the default.
func Indent(indentW float32) { C.igIndent(C.float(indentW)) }

// Unindent decreases the indent. A zero width uses the default.
func Unindent(indentW float32) { C.igUnindent(C.float(indentW)) }

// ---- Basic widgets -------------------------------------------------------

// Button draws a button and reports whether it was clicked. A zero size
// auto-fits the label.
func Button(label string, sizeX, sizeY float32) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igButton(clabel, vec2(sizeX, sizeY)))
}

// SmallButton draws a button with no frame padding.
func SmallButton(label string) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igSmallButton(clabel))
}

// Checkbox draws a checkbox bound to v and reports whether it changed.
func Checkbox(label string, v *bool) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cv := C.bool(*v)
	ret := bool(C.igCheckbox(clabel, &cv))
	*v = bool(cv)
	return ret
}

// RadioButton draws a radio button and reports whether it was clicked.
func RadioButton(label string, active bool) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igRadioButton_Bool(clabel, C.bool(active)))
}

// ---- Sliders / inputs / drags --------------------------------------------

// SliderFloat draws a float slider bound to v and reports whether it changed.
func SliderFloat(label string, v *float32, vMin, vMax float32, format string, flags int) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	cv := C.float(*v)
	ret := bool(C.igSliderFloat(clabel, &cv, C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
	*v = float32(cv)
	return ret
}

// SliderInt draws an int slider bound to v and reports whether it changed.
func SliderInt(label string, v *int32, vMin, vMax int32, format string, flags int) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	cv := C.int(*v)
	ret := bool(C.igSliderInt(clabel, &cv, C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
	*v = int32(cv)
	return ret
}

// InputFloat draws a float input bound to v and reports whether it changed.
func InputFloat(label string, v *float32, step, stepFast float32, format string, flags int) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	cv := C.float(*v)
	ret := bool(C.igInputFloat(clabel, &cv, C.float(step), C.float(stepFast), cformat, C.ImGuiInputTextFlags(flags)))
	*v = float32(cv)
	return ret
}

// InputInt draws an int input bound to v and reports whether it changed.
func InputInt(label string, v *int32, step, stepFast int32, flags int) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cv := C.int(*v)
	ret := bool(C.igInputInt(clabel, &cv, C.int(step), C.int(stepFast), C.ImGuiInputTextFlags(flags)))
	*v = int32(cv)
	return ret
}

// DragFloat draws a draggable float bound to v and reports whether it changed.
func DragFloat(label string, v *float32, speed, vMin, vMax float32, format string, flags int) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	cv := C.float(*v)
	ret := bool(C.igDragFloat(clabel, &cv, C.float(speed), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
	*v = float32(cv)
	return ret
}

// DragInt draws a draggable int bound to v and reports whether it changed.
func DragInt(label string, v *int32, speed float32, vMin, vMax int32, format string, flags int) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cformat := C.CString(format)
	defer C.free(unsafe.Pointer(cformat))
	cv := C.int(*v)
	ret := bool(C.igDragInt(clabel, &cv, C.float(speed), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
	*v = int32(cv)
	return ret
}

// ColorEdit3 edits an RGB color in place and reports whether it changed.
func ColorEdit3(label string, col *[3]float32, flags int) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igColorEdit3(clabel, (*C.float)(unsafe.Pointer(&col[0])), C.ImGuiColorEditFlags(flags)))
}

// ColorEdit4 edits an RGBA color in place and reports whether it changed.
func ColorEdit4(label string, col *[4]float32, flags int) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igColorEdit4(clabel, (*C.float)(unsafe.Pointer(&col[0])), C.ImGuiColorEditFlags(flags)))
}

// ---- Trees / headers -----------------------------------------------------

// TreeNode opens a tree node labeled by label. Call [TreePop] if it returns true.
func TreeNode(label string) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igTreeNode_Str(clabel))
}

// TreePop closes the tree node opened by [TreeNode].
func TreePop() { C.igTreePop() }

// CollapsingHeader draws a collapsing header and reports whether it is open.
func CollapsingHeader(label string, flags int) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igCollapsingHeader_TreeNodeFlags(clabel, C.ImGuiTreeNodeFlags(flags)))
}

// ---- Combo / selectable --------------------------------------------------

// BeginCombo opens a combo box. Call [EndCombo] if it returns true.
func BeginCombo(label, previewValue string, flags int) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cpreview := C.CString(previewValue)
	defer C.free(unsafe.Pointer(cpreview))
	return bool(C.igBeginCombo(clabel, cpreview, C.ImGuiComboFlags(flags)))
}

// EndCombo closes the combo box opened by [BeginCombo].
func EndCombo() { C.igEndCombo() }

// Selectable draws a selectable item and reports whether it was clicked. A zero
// size fits the label.
func Selectable(label string, selected bool, flags int, sizeX, sizeY float32) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igSelectable_Bool(clabel, C.bool(selected), C.ImGuiSelectableFlags(flags), vec2(sizeX, sizeY)))
}

// ---- Tooltip -------------------------------------------------------------

// BeginTooltip opens a tooltip. Call [EndTooltip] if it returns true.
func BeginTooltip() bool { return bool(C.igBeginTooltip()) }

// EndTooltip closes the tooltip opened by [BeginTooltip].
func EndTooltip() { C.igEndTooltip() }

// ---- Demo / debug windows ------------------------------------------------

// ShowDemoWindow displays the Dear ImGui demo window. pOpen may be nil.
func ShowDemoWindow(pOpen *bool) {
	withBoolPtr(pOpen, func(p *C.bool) { C.igShowDemoWindow(p) })
}

// ShowMetricsWindow displays the metrics/debug window. pOpen may be nil.
func ShowMetricsWindow(pOpen *bool) {
	withBoolPtr(pOpen, func(p *C.bool) { C.igShowMetricsWindow(p) })
}

// ShowAboutWindow displays the about window. pOpen may be nil.
func ShowAboutWindow(pOpen *bool) {
	withBoolPtr(pOpen, func(p *C.bool) { C.igShowAboutWindow(p) })
}

// vec2 builds a cimgui ImVec2 from x/y.
func vec2(x, y float32) C.ImVec2_c { return C.ImVec2_c{x: C.float(x), y: C.float(y)} }
