package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// TabBarOptions are the optional inputs to [TabBar]. A nil *TabBarOptions uses
// Dear ImGui's defaults; each field maps to an ImGuiTabBarFlags_ bit.
type TabBarOptions struct {
	Reorderable                  bool // ImGuiTabBarFlags_Reorderable
	AutoSelectNewTabs            bool // ImGuiTabBarFlags_AutoSelectNewTabs
	TabListPopupButton           bool // ImGuiTabBarFlags_TabListPopupButton
	NoCloseWithMiddleMouseButton bool // ImGuiTabBarFlags_NoCloseWithMiddleMouseButton
	NoTabListScrollingButtons    bool // ImGuiTabBarFlags_NoTabListScrollingButtons
	NoTooltip                    bool // ImGuiTabBarFlags_NoTooltip
	DrawSelectedOverline         bool // ImGuiTabBarFlags_DrawSelectedOverline
	FittingPolicyShrink          bool // ImGuiTabBarFlags_FittingPolicyShrink
	FittingPolicyScroll          bool // ImGuiTabBarFlags_FittingPolicyScroll
	FittingPolicyMixed           bool // ImGuiTabBarFlags_FittingPolicyMixed
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *TabBarOptions) flags() cimgui.TabBarFlags {
	if o == nil {
		return cimgui.TabBarFlagsNone
	}
	var f cimgui.TabBarFlags
	if o.Reorderable {
		f |= cimgui.TabBarFlagsReorderable
	}
	if o.AutoSelectNewTabs {
		f |= cimgui.TabBarFlagsAutoSelectNewTabs
	}
	if o.TabListPopupButton {
		f |= cimgui.TabBarFlagsTabListPopupButton
	}
	if o.NoCloseWithMiddleMouseButton {
		f |= cimgui.TabBarFlagsNoCloseWithMiddleMouseButton
	}
	if o.NoTabListScrollingButtons {
		f |= cimgui.TabBarFlagsNoTabListScrollingButtons
	}
	if o.NoTooltip {
		f |= cimgui.TabBarFlagsNoTooltip
	}
	if o.DrawSelectedOverline {
		f |= cimgui.TabBarFlagsDrawSelectedOverline
	}
	if o.FittingPolicyShrink {
		f |= cimgui.TabBarFlagsFittingPolicyShrink
	}
	if o.FittingPolicyScroll {
		f |= cimgui.TabBarFlagsFittingPolicyScroll
	}
	if o.FittingPolicyMixed {
		f |= cimgui.TabBarFlagsFittingPolicyMixed
	}
	return f
}

// TabItemOptions are the optional inputs to [TabItem] and [TabItemButton]. A nil
// *TabItemOptions uses Dear ImGui's defaults; each field maps to an
// ImGuiTabItemFlags_ bit.
type TabItemOptions struct {
	UnsavedDocument              bool // ImGuiTabItemFlags_UnsavedDocument
	SetSelected                  bool // ImGuiTabItemFlags_SetSelected
	NoCloseWithMiddleMouseButton bool // ImGuiTabItemFlags_NoCloseWithMiddleMouseButton
	NoPushID                     bool // ImGuiTabItemFlags_NoPushId
	NoTooltip                    bool // ImGuiTabItemFlags_NoTooltip
	NoReorder                    bool // ImGuiTabItemFlags_NoReorder
	Leading                      bool // ImGuiTabItemFlags_Leading
	Trailing                     bool // ImGuiTabItemFlags_Trailing
	NoAssumedClosure             bool // ImGuiTabItemFlags_NoAssumedClosure
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *TabItemOptions) flags() cimgui.TabItemFlags {
	if o == nil {
		return cimgui.TabItemFlagsNone
	}
	var f cimgui.TabItemFlags
	if o.UnsavedDocument {
		f |= cimgui.TabItemFlagsUnsavedDocument
	}
	if o.SetSelected {
		f |= cimgui.TabItemFlagsSetSelected
	}
	if o.NoCloseWithMiddleMouseButton {
		f |= cimgui.TabItemFlagsNoCloseWithMiddleMouseButton
	}
	if o.NoPushID {
		f |= cimgui.TabItemFlagsNoPushId
	}
	if o.NoTooltip {
		f |= cimgui.TabItemFlagsNoTooltip
	}
	if o.NoReorder {
		f |= cimgui.TabItemFlagsNoReorder
	}
	if o.Leading {
		f |= cimgui.TabItemFlagsLeading
	}
	if o.Trailing {
		f |= cimgui.TabItemFlagsTrailing
	}
	if o.NoAssumedClosure {
		f |= cimgui.TabItemFlagsNoAssumedClosure
	}
	return f
}

// TabBar begins a tab bar identified by strID. It models ImGui::BeginTabBar.
// open reports whether the bar is visible; the returned [EndFunc]
// (ImGui::EndTabBar) ends it only when open.
func TabBar(strID string, opts *TabBarOptions) (open bool, end EndFunc) {
	open = cimgui.BeginTabBar(strID, opts.flags())
	if !open {
		return open, func() {}
	}
	return open, cimgui.EndTabBar
}

// TabItem begins a tab within the current tab bar. It models ImGui::BeginTabItem.
// When open is non-nil a close button is shown and *open is updated; selected
// reports whether the tab's contents should be drawn. The returned [EndFunc]
// (ImGui::EndTabItem) ends it only when selected.
func TabItem(label string, open *bool, opts *TabItemOptions) (selected bool, end EndFunc) {
	selected = cimgui.BeginTabItem(label, open, opts.flags())
	if !selected {
		return selected, func() {}
	}
	return selected, cimgui.EndTabItem
}

// TabItemButton draws a tab that behaves like a button and reports whether it was
// clicked. It models ImGui::TabItemButton.
func TabItemButton(label string, opts *TabItemOptions) bool {
	return cimgui.TabItemButton(label, opts.flags())
}

// SetTabItemClosed notifies the tab bar that the named tab or window was closed
// externally this frame. It models ImGui::SetTabItemClosed.
func SetTabItemClosed(label string) {
	cimgui.SetTabItemClosed(label)
}
