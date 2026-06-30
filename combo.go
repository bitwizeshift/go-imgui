package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// ComboOptions are the optional inputs to [Combo]. A nil *ComboOptions uses Dear
// ImGui's defaults; each field maps to an ImGuiComboFlags_ bit.
type ComboOptions struct {
	PopupAlignLeft  bool // ImGuiComboFlags_PopupAlignLeft
	HeightSmall     bool // ImGuiComboFlags_HeightSmall
	HeightRegular   bool // ImGuiComboFlags_HeightRegular
	HeightLarge     bool // ImGuiComboFlags_HeightLarge
	HeightLargest   bool // ImGuiComboFlags_HeightLargest
	NoArrowButton   bool // ImGuiComboFlags_NoArrowButton
	NoPreview       bool // ImGuiComboFlags_NoPreview
	WidthFitPreview bool // ImGuiComboFlags_WidthFitPreview
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *ComboOptions) flags() cimgui.ComboFlags {
	if o == nil {
		return cimgui.ComboFlagsNone
	}
	var f cimgui.ComboFlags
	if o.PopupAlignLeft {
		f |= cimgui.ComboFlagsPopupAlignLeft
	}
	if o.HeightSmall {
		f |= cimgui.ComboFlagsHeightSmall
	}
	if o.HeightRegular {
		f |= cimgui.ComboFlagsHeightRegular
	}
	if o.HeightLarge {
		f |= cimgui.ComboFlagsHeightLarge
	}
	if o.HeightLargest {
		f |= cimgui.ComboFlagsHeightLargest
	}
	if o.NoArrowButton {
		f |= cimgui.ComboFlagsNoArrowButton
	}
	if o.NoPreview {
		f |= cimgui.ComboFlagsNoPreview
	}
	if o.WidthFitPreview {
		f |= cimgui.ComboFlagsWidthFitPreview
	}
	return f
}

// Combo begins a combo box showing previewValue, into which arbitrary selectable
// content is drawn. It models ImGui::BeginCombo. open reports whether the popup
// is open; the returned [EndFunc] (ImGui::EndCombo) ends it only when open. For
// the simple list-selection forms see [ComboItems], [ComboZeroSep] and
// [ComboFunc].
func Combo(label, previewValue string, opts *ComboOptions) (open bool, end EndFunc) {
	open = cimgui.BeginCombo(label, previewValue, opts.flags())
	if !open {
		return open, func() {}
	}
	return open, cimgui.EndCombo
}

// ComboItems draws a combo box selecting an index within items, updating
// currentItem and reporting whether it changed. A negative popupMaxHeightInItems
// uses the default. It models ImGui::Combo (the items-array overload).
func ComboItems(label string, currentItem *int32, items []string, popupMaxHeightInItems int32) bool {
	return cimgui.Combo_Str_arr(label, currentItem, items, popupMaxHeightInItems)
}

// ComboZeroSep draws a combo box whose items come from a single string of
// NUL-separated, double-NUL-terminated entries (e.g. "a\x00b\x00c\x00"). It
// models ImGui::Combo (the zero-separated-string overload).
func ComboZeroSep(label string, currentItem *int32, itemsSeparatedByZeros string, popupMaxHeightInItems int32) bool {
	return cimgui.Combo_Str(label, currentItem, itemsSeparatedByZeros, popupMaxHeightInItems)
}

// ComboFunc draws a combo box whose itemsCount labels are produced lazily by
// getter, updating currentItem and reporting whether it changed. A negative
// popupMaxHeightInItems uses the default. It models ImGui::Combo (the getter
// overload).
func ComboFunc(label string, currentItem *int32, getter func(idx int32) string, itemsCount, popupMaxHeightInItems int32) bool {
	return cimgui.Combo_FnStrPtr(label, currentItem, getter, itemsCount, popupMaxHeightInItems)
}

// ListBox begins a scrolling list box of the given size, into which selectable
// content is drawn. It models ImGui::BeginListBox. open reports whether the box
// is visible; the returned [EndFunc] (ImGui::EndListBox) ends it only when open.
// For the simple list-selection forms see [ListBoxItems] and [ListBoxFunc].
func ListBox(label string, size Vec2) (open bool, end EndFunc) {
	open = cimgui.BeginListBox(label, size)
	if !open {
		return open, func() {}
	}
	return open, cimgui.EndListBox
}

// ListBoxItems draws a list box selecting an index within items, updating
// currentItem and reporting whether it changed. A negative heightInItems uses
// the default. It models ImGui::ListBox (the items-array overload).
func ListBoxItems(label string, currentItem *int32, items []string, heightInItems int32) bool {
	return cimgui.ListBox_Str_arr(label, currentItem, items, heightInItems)
}

// ListBoxFunc draws a list box whose itemsCount labels are produced lazily by
// getter, updating currentItem and reporting whether it changed. A negative
// heightInItems uses the default. It models ImGui::ListBox (the getter overload).
func ListBoxFunc(label string, currentItem *int32, getter func(idx int32) string, itemsCount, heightInItems int32) bool {
	return cimgui.ListBox_FnStrPtr(label, currentItem, getter, itemsCount, heightInItems)
}
