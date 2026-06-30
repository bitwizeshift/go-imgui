package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// Style is a view over the live ImGuiStyle. Mutating it changes the global style
// in place. It models ImGuiStyle.
type Style = cimgui.Style

// GetStyle returns the current global [Style]. It models ImGui::GetStyle.
func GetStyle() *Style {
	return cimgui.GetStyle()
}

// GetStyleColorVec4 returns the current style color for idx. It models
// ImGui::GetStyleColorVec4.
func GetStyleColorVec4(idx Col) Vec4 {
	return cimgui.GetStyleColorVec4(idx)
}

// PushStyleColor pushes col onto the style-color stack for idx; balance it with
// [PopStyleColor]. It models ImGui::PushStyleColor. Prefer [StyleColor] for
// scoped use.
func PushStyleColor(idx Col, col Vec4) {
	cimgui.PushStyleColor_Vec4(idx, col)
}

// PopStyleColor pops count entries from the style-color stack. It models
// ImGui::PopStyleColor.
func PopStyleColor(count int32) {
	cimgui.PopStyleColor(count)
}

// StyleColor pushes col onto the style-color stack for idx. It models
// ImGui::PushStyleColor. The returned [EndFunc] pops the single entry
// (ImGui::PopStyleColor).
func StyleColor(idx Col, col Vec4) (pop EndFunc) {
	cimgui.PushStyleColor_Vec4(idx, col)
	return func() {
		cimgui.PopStyleColor(1)
	}
}

// PushStyleVar pushes val onto the style-variable stack for idx; balance it with
// [PopStyleVar]. val may be a float32 or a [Vec2], modelling the float and
// ImVec2 overloads of ImGui::PushStyleVar. Prefer [StyleVarScope] for scoped use.
func PushStyleVar[T StyleVarValue](idx StyleVar, val T) {
	switch v := any(val).(type) {
	case float32:
		cimgui.PushStyleVar_Float(idx, v)
	case Vec2:
		cimgui.PushStyleVar_Vec2(idx, v)
	}
}

// PopStyleVar pops count entries from the style-variable stack. It models
// ImGui::PopStyleVar.
func PopStyleVar(count int32) {
	cimgui.PopStyleVar(count)
}

// StyleVarScope pushes val onto the style-variable stack for idx. val may be a
// float32 or a [Vec2], modelling the overloads of ImGui::PushStyleVar. The
// returned [EndFunc] pops the single entry (ImGui::PopStyleVar).
func StyleVarScope[T StyleVarValue](idx StyleVar, val T) (pop EndFunc) {
	PushStyleVar(idx, val)
	return func() {
		cimgui.PopStyleVar(1)
	}
}
