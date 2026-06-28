package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

// Col identifies a styleable interface color. Mirrors the public ImGuiCol_.
type Col int32

const (
	ColText             Col = C.ImGuiCol_Text
	ColTextDisabled     Col = C.ImGuiCol_TextDisabled
	ColWindowBg         Col = C.ImGuiCol_WindowBg
	ColChildBg          Col = C.ImGuiCol_ChildBg
	ColPopupBg          Col = C.ImGuiCol_PopupBg
	ColBorder           Col = C.ImGuiCol_Border
	ColFrameBg          Col = C.ImGuiCol_FrameBg
	ColFrameBgHovered   Col = C.ImGuiCol_FrameBgHovered
	ColFrameBgActive    Col = C.ImGuiCol_FrameBgActive
	ColTitleBg          Col = C.ImGuiCol_TitleBg
	ColTitleBgActive    Col = C.ImGuiCol_TitleBgActive
	ColTitleBgCollapsed Col = C.ImGuiCol_TitleBgCollapsed
	ColMenuBarBg        Col = C.ImGuiCol_MenuBarBg
	ColCheckMark        Col = C.ImGuiCol_CheckMark
	ColSliderGrab       Col = C.ImGuiCol_SliderGrab
	ColSliderGrabActive Col = C.ImGuiCol_SliderGrabActive
	ColButton           Col = C.ImGuiCol_Button
	ColButtonHovered    Col = C.ImGuiCol_ButtonHovered
	ColButtonActive     Col = C.ImGuiCol_ButtonActive
	ColHeader           Col = C.ImGuiCol_Header
	ColHeaderHovered    Col = C.ImGuiCol_HeaderHovered
	ColHeaderActive     Col = C.ImGuiCol_HeaderActive
	ColSeparator        Col = C.ImGuiCol_Separator
	ColTab              Col = C.ImGuiCol_Tab
	ColTabHovered       Col = C.ImGuiCol_TabHovered
	ColTabSelected      Col = C.ImGuiCol_TabSelected
	ColPlotLines        Col = C.ImGuiCol_PlotLines
	ColPlotHistogram    Col = C.ImGuiCol_PlotHistogram
	ColTextSelectedBg   Col = C.ImGuiCol_TextSelectedBg
	ColModalWindowDimBg Col = C.ImGuiCol_ModalWindowDimBg
)

// StyleVar identifies a styleable layout variable. Mirrors the public
// ImGuiStyleVar_.
type StyleVar int32

const (
	StyleVarAlpha            StyleVar = C.ImGuiStyleVar_Alpha
	StyleVarDisabledAlpha    StyleVar = C.ImGuiStyleVar_DisabledAlpha
	StyleVarWindowPadding    StyleVar = C.ImGuiStyleVar_WindowPadding
	StyleVarWindowRounding   StyleVar = C.ImGuiStyleVar_WindowRounding
	StyleVarWindowBorderSize StyleVar = C.ImGuiStyleVar_WindowBorderSize
	StyleVarFramePadding     StyleVar = C.ImGuiStyleVar_FramePadding
	StyleVarFrameRounding    StyleVar = C.ImGuiStyleVar_FrameRounding
	StyleVarFrameBorderSize  StyleVar = C.ImGuiStyleVar_FrameBorderSize
	StyleVarItemSpacing      StyleVar = C.ImGuiStyleVar_ItemSpacing
	StyleVarItemInnerSpacing StyleVar = C.ImGuiStyleVar_ItemInnerSpacing
	StyleVarIndentSpacing    StyleVar = C.ImGuiStyleVar_IndentSpacing
	StyleVarScrollbarSize    StyleVar = C.ImGuiStyleVar_ScrollbarSize
	StyleVarGrabMinSize      StyleVar = C.ImGuiStyleVar_GrabMinSize
	StyleVarGrabRounding     StyleVar = C.ImGuiStyleVar_GrabRounding
	StyleVarTabRounding      StyleVar = C.ImGuiStyleVar_TabRounding
)

// PushStyleColor_Vec4 pushes col onto the style-color stack for idx. Balance it
// with [PopStyleColor].
func PushStyleColor_Vec4(idx Col, col Vec4) {
	C.igPushStyleColor_Vec4(C.ImGuiCol(idx), col.c())
}

// PopStyleColor pops count entries from the style-color stack.
func PopStyleColor(count int32) {
	C.igPopStyleColor(C.int(count))
}

// PushStyleVar_Float pushes a float style variable. Balance it with [PopStyleVar].
func PushStyleVar_Float(idx StyleVar, val float32) {
	C.igPushStyleVar_Float(C.ImGuiStyleVar(idx), C.float(val))
}

// PushStyleVar_Vec2 pushes a [Vec2] style variable. Balance it with [PopStyleVar].
func PushStyleVar_Vec2(idx StyleVar, val Vec2) {
	C.igPushStyleVar_Vec2(C.ImGuiStyleVar(idx), val.c())
}

// PopStyleVar pops count entries from the style-variable stack.
func PopStyleVar(count int32) {
	C.igPopStyleVar(C.int(count))
}

// GetStyleColorVec4 returns the current style color for idx.
func GetStyleColorVec4(idx Col) Vec4 {
	return vec4(*C.igGetStyleColorVec4(C.ImGuiCol(idx)))
}

// Style is a view over the live ImGuiStyle. Mutating it changes the global style
// in place.
type Style struct {
	c *C.ImGuiStyle
}

// GetStyle returns the current global [Style].
func GetStyle() *Style {
	return &Style{c: C.igGetStyle()}
}

// Alpha is the global opacity applied to everything.
func (s *Style) Alpha() float32 {
	return float32(s.c.Alpha)
}

// SetAlpha sets the global opacity.
func (s *Style) SetAlpha(v float32) {
	s.c.Alpha = C.float(v)
}

// DisabledAlpha is the opacity applied to disabled items.
func (s *Style) DisabledAlpha() float32 {
	return float32(s.c.DisabledAlpha)
}

// SetDisabledAlpha sets the opacity for disabled items.
func (s *Style) SetDisabledAlpha(v float32) {
	s.c.DisabledAlpha = C.float(v)
}

// WindowRounding is the radius of window corners.
func (s *Style) WindowRounding() float32 {
	return float32(s.c.WindowRounding)
}

// SetWindowRounding sets the radius of window corners.
func (s *Style) SetWindowRounding(v float32) {
	s.c.WindowRounding = C.float(v)
}

// FrameRounding is the radius of frame corners (buttons, inputs, …).
func (s *Style) FrameRounding() float32 {
	return float32(s.c.FrameRounding)
}

// SetFrameRounding sets the radius of frame corners.
func (s *Style) SetFrameRounding(v float32) {
	s.c.FrameRounding = C.float(v)
}

// GrabRounding is the radius of slider/grab corners.
func (s *Style) GrabRounding() float32 {
	return float32(s.c.GrabRounding)
}

// SetGrabRounding sets the radius of slider/grab corners.
func (s *Style) SetGrabRounding(v float32) {
	s.c.GrabRounding = C.float(v)
}

// TabRounding is the radius of tab corners.
func (s *Style) TabRounding() float32 {
	return float32(s.c.TabRounding)
}

// SetTabRounding sets the radius of tab corners.
func (s *Style) SetTabRounding(v float32) {
	s.c.TabRounding = C.float(v)
}

// WindowPadding is the padding within a window.
func (s *Style) WindowPadding() Vec2 {
	return vec2(s.c.WindowPadding)
}

// SetWindowPadding sets the padding within a window.
func (s *Style) SetWindowPadding(v Vec2) {
	s.c.WindowPadding = v.c()
}

// FramePadding is the padding within framed items.
func (s *Style) FramePadding() Vec2 {
	return vec2(s.c.FramePadding)
}

// SetFramePadding sets the padding within framed items.
func (s *Style) SetFramePadding(v Vec2) {
	s.c.FramePadding = v.c()
}

// ItemSpacing is the spacing between successive items.
func (s *Style) ItemSpacing() Vec2 {
	return vec2(s.c.ItemSpacing)
}

// SetItemSpacing sets the spacing between successive items.
func (s *Style) SetItemSpacing(v Vec2) {
	s.c.ItemSpacing = v.c()
}

// Color returns the style color for idx.
func (s *Style) Color(idx Col) Vec4 {
	return vec4(s.c.Colors[int(idx)])
}

// SetColor sets the style color for idx.
func (s *Style) SetColor(idx Col, col Vec4) {
	s.c.Colors[int(idx)] = col.c()
}
