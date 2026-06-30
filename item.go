package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// HoveredOptions are the optional inputs to [IsItemHovered]. A nil
// *HoveredOptions uses Dear ImGui's defaults; each field maps to an
// ImGuiHoveredFlags_ bit.
type HoveredOptions struct {
	ChildWindows                 bool // ImGuiHoveredFlags_ChildWindows
	RootWindow                   bool // ImGuiHoveredFlags_RootWindow
	AnyWindow                    bool // ImGuiHoveredFlags_AnyWindow
	NoPopupHierarchy             bool // ImGuiHoveredFlags_NoPopupHierarchy
	AllowWhenBlockedByPopup      bool // ImGuiHoveredFlags_AllowWhenBlockedByPopup
	AllowWhenBlockedByActiveItem bool // ImGuiHoveredFlags_AllowWhenBlockedByActiveItem
	AllowWhenOverlapped          bool // ImGuiHoveredFlags_AllowWhenOverlapped
	AllowWhenDisabled            bool // ImGuiHoveredFlags_AllowWhenDisabled
	RectOnly                     bool // ImGuiHoveredFlags_RectOnly (composite)
	RootAndChildWindows          bool // ImGuiHoveredFlags_RootAndChildWindows (composite)
	ForTooltip                   bool // ImGuiHoveredFlags_ForTooltip
	Stationary                   bool // ImGuiHoveredFlags_Stationary
	DelayNone                    bool // ImGuiHoveredFlags_DelayNone
	DelayShort                   bool // ImGuiHoveredFlags_DelayShort
	DelayNormal                  bool // ImGuiHoveredFlags_DelayNormal
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *HoveredOptions) flags() cimgui.HoveredFlags {
	if o == nil {
		return cimgui.HoveredFlagsNone
	}
	var f cimgui.HoveredFlags
	if o.ChildWindows {
		f |= cimgui.HoveredFlagsChildWindows
	}
	if o.RootWindow {
		f |= cimgui.HoveredFlagsRootWindow
	}
	if o.AnyWindow {
		f |= cimgui.HoveredFlagsAnyWindow
	}
	if o.NoPopupHierarchy {
		f |= cimgui.HoveredFlagsNoPopupHierarchy
	}
	if o.AllowWhenBlockedByPopup {
		f |= cimgui.HoveredFlagsAllowWhenBlockedByPopup
	}
	if o.AllowWhenBlockedByActiveItem {
		f |= cimgui.HoveredFlagsAllowWhenBlockedByActiveItem
	}
	if o.AllowWhenOverlapped {
		f |= cimgui.HoveredFlagsAllowWhenOverlapped
	}
	if o.AllowWhenDisabled {
		f |= cimgui.HoveredFlagsAllowWhenDisabled
	}
	if o.RectOnly {
		f |= cimgui.HoveredFlagsRectOnly
	}
	if o.RootAndChildWindows {
		f |= cimgui.HoveredFlagsRootAndChildWindows
	}
	if o.ForTooltip {
		f |= cimgui.HoveredFlagsForTooltip
	}
	if o.Stationary {
		f |= cimgui.HoveredFlagsStationary
	}
	if o.DelayNone {
		f |= cimgui.HoveredFlagsDelayNone
	}
	if o.DelayShort {
		f |= cimgui.HoveredFlagsDelayShort
	}
	if o.DelayNormal {
		f |= cimgui.HoveredFlagsDelayNormal
	}
	return f
}

// IsItemHovered reports whether the previous item is hovered. It models
// ImGui::IsItemHovered.
func IsItemHovered(opts *HoveredOptions) bool {
	return cimgui.IsItemHovered(opts.flags())
}

// IsItemActive reports whether the previous item is active (e.g. held). It models
// ImGui::IsItemActive.
func IsItemActive() bool {
	return cimgui.IsItemActive()
}

// IsItemClicked reports whether the previous item was clicked with the given
// mouse button. It models ImGui::IsItemClicked.
func IsItemClicked(button MouseButton) bool {
	return cimgui.IsItemClicked(button)
}

// SetItemDefaultFocus makes the previous item the default-focused one when the
// window appears. It models ImGui::SetItemDefaultFocus.
func SetItemDefaultFocus() {
	cimgui.SetItemDefaultFocus()
}

// SetKeyboardFocusHere focuses the next item (offset 0) or a later item. It
// models ImGui::SetKeyboardFocusHere.
func SetKeyboardFocusHere(offset int32) {
	cimgui.SetKeyboardFocusHere(offset)
}

// ItemWidth pushes the width of common widgets. It models ImGui::PushItemWidth.
// The returned [EndFunc] (ImGui::PopItemWidth) restores the previous width.
func ItemWidth(itemWidth float32) (pop EndFunc) {
	cimgui.PushItemWidth(itemWidth)
	return cimgui.PopItemWidth
}

// PushItemWidth pushes the width of common widgets; balance it with
// [PopItemWidth]. It models ImGui::PushItemWidth. Prefer [ItemWidth] for scoped
// use.
func PushItemWidth(itemWidth float32) {
	cimgui.PushItemWidth(itemWidth)
}

// PopItemWidth restores the width pushed by [PushItemWidth]. It models
// ImGui::PopItemWidth.
func PopItemWidth() {
	cimgui.PopItemWidth()
}

// SetNextItemWidth sets the width of the next common widget. It models
// ImGui::SetNextItemWidth.
func SetNextItemWidth(itemWidth float32) {
	cimgui.SetNextItemWidth(itemWidth)
}

// Disabled begins a disabled block when disabled is true. It models
// ImGui::BeginDisabled. The returned [EndFunc] (ImGui::EndDisabled) ends the
// block and must always be called.
func Disabled(disabled bool) (end EndFunc) {
	cimgui.BeginDisabled(disabled)
	return cimgui.EndDisabled
}
