package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// SelectableOptions are the optional inputs to [Selectable].
// A nil *SelectableOptions uses Dear ImGui's defaults; each field maps to an
// ImGuiSelectableFlags_ bit.
type SelectableOptions struct {
	NoAutoClosePopups bool // ImGuiSelectableFlags_NoAutoClosePopups
	SpanAllColumns    bool // ImGuiSelectableFlags_SpanAllColumns
	AllowDoubleClick  bool // ImGuiSelectableFlags_AllowDoubleClick
	Disabled          bool // ImGuiSelectableFlags_Disabled
	AllowOverlap      bool // ImGuiSelectableFlags_AllowOverlap
	Highlight         bool // ImGuiSelectableFlags_Highlight
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *SelectableOptions) flags() cimgui.SelectableFlags {
	if o == nil {
		return cimgui.SelectableFlagsNone
	}
	var f cimgui.SelectableFlags
	if o.NoAutoClosePopups {
		f |= cimgui.SelectableFlagsNoAutoClosePopups
	}
	if o.SpanAllColumns {
		f |= cimgui.SelectableFlagsSpanAllColumns
	}
	if o.AllowDoubleClick {
		f |= cimgui.SelectableFlagsAllowDoubleClick
	}
	if o.Disabled {
		f |= cimgui.SelectableFlagsDisabled
	}
	if o.AllowOverlap {
		f |= cimgui.SelectableFlagsAllowOverlap
	}
	if o.Highlight {
		f |= cimgui.SelectableFlagsHighlight
	}
	return f
}

// Selectable draws a selectable item and reports whether it was clicked. A zero
// size fits the label. selected is either a bool (rendered selected) or a *bool
// (toggled on click); the form is inferred at the call site. It models
// ImGui::Selectable.
func Selectable[T BoolOrPtr](label string, selected T, size Vec2, opts *SelectableOptions) bool {
	switch v := any(selected).(type) {
	case bool:
		return cimgui.Selectable_Bool(label, v, opts.flags(), size)
	case *bool:
		return cimgui.Selectable_BoolPtr(label, v, opts.flags(), size)
	}
	return false
}
