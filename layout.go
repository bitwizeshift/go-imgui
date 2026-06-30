package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// Separator draws a horizontal line. It models ImGui::Separator.
func Separator() {
	cimgui.Separator()
}

// SameLine continues the current line. Pass zeros for default spacing. It models
// ImGui::SameLine.
func SameLine(offsetFromStartX, spacing float32) {
	cimgui.SameLine(offsetFromStartX, spacing)
}

// NewLine moves to the next line. It models ImGui::NewLine.
func NewLine() {
	cimgui.NewLine()
}

// Spacing adds vertical spacing. It models ImGui::Spacing.
func Spacing() {
	cimgui.Spacing()
}

// Dummy adds an empty item of the given size. It models ImGui::Dummy.
func Dummy(size Vec2) {
	cimgui.Dummy(size)
}

// Indent increases the indent. A zero width uses the default. It models
// ImGui::Indent.
func Indent(indentW float32) {
	cimgui.Indent(indentW)
}

// Unindent decreases the indent. A zero width uses the default. It models
// ImGui::Unindent.
func Unindent(indentW float32) {
	cimgui.Unindent(indentW)
}

// Group begins a group; layout queries treat the group as one item until the
// scope ends. It models ImGui::BeginGroup. The returned [EndFunc]
// (ImGui::EndGroup) ends the group.
func Group() (end EndFunc) {
	cimgui.BeginGroup()
	return cimgui.EndGroup
}

// GetCursorScreenPos returns the cursor position in absolute screen coordinates,
// the origin used by the draw-list primitives. It models
// ImGui::GetCursorScreenPos.
func GetCursorScreenPos() Vec2 {
	return cimgui.GetCursorScreenPos()
}

// SetCursorScreenPos moves the cursor to pos in absolute screen coordinates. It
// models ImGui::SetCursorScreenPos.
func SetCursorScreenPos(pos Vec2) {
	cimgui.SetCursorScreenPos(pos)
}

// GetCursorPos returns the cursor position in window-local coordinates. It models
// ImGui::GetCursorPos.
func GetCursorPos() Vec2 {
	return cimgui.GetCursorPos()
}

// GetContentRegionAvail returns the space remaining from the cursor to the edge
// of the current content region. It models ImGui::GetContentRegionAvail.
func GetContentRegionAvail() Vec2 {
	return cimgui.GetContentRegionAvail()
}
