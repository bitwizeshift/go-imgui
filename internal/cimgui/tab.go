package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// TabBarFlags configures a tab bar opened with [BeginTabBar]. Mirrors the public
// ImGuiTabBarFlags_.
type TabBarFlags int32

const (
	TabBarFlagsNone                         TabBarFlags = C.ImGuiTabBarFlags_None
	TabBarFlagsReorderable                  TabBarFlags = C.ImGuiTabBarFlags_Reorderable
	TabBarFlagsAutoSelectNewTabs            TabBarFlags = C.ImGuiTabBarFlags_AutoSelectNewTabs
	TabBarFlagsTabListPopupButton           TabBarFlags = C.ImGuiTabBarFlags_TabListPopupButton
	TabBarFlagsNoCloseWithMiddleMouseButton TabBarFlags = C.ImGuiTabBarFlags_NoCloseWithMiddleMouseButton
	TabBarFlagsNoTabListScrollingButtons    TabBarFlags = C.ImGuiTabBarFlags_NoTabListScrollingButtons
	TabBarFlagsNoTooltip                    TabBarFlags = C.ImGuiTabBarFlags_NoTooltip
	TabBarFlagsDrawSelectedOverline         TabBarFlags = C.ImGuiTabBarFlags_DrawSelectedOverline
	TabBarFlagsFittingPolicyShrink          TabBarFlags = C.ImGuiTabBarFlags_FittingPolicyShrink
	TabBarFlagsFittingPolicyScroll          TabBarFlags = C.ImGuiTabBarFlags_FittingPolicyScroll
	TabBarFlagsFittingPolicyMixed           TabBarFlags = C.ImGuiTabBarFlags_FittingPolicyMixed
)

// TabItemFlags configures a tab item opened with [BeginTabItem]. Mirrors the
// public ImGuiTabItemFlags_.
type TabItemFlags int32

const (
	TabItemFlagsNone                         TabItemFlags = C.ImGuiTabItemFlags_None
	TabItemFlagsUnsavedDocument              TabItemFlags = C.ImGuiTabItemFlags_UnsavedDocument
	TabItemFlagsSetSelected                  TabItemFlags = C.ImGuiTabItemFlags_SetSelected
	TabItemFlagsNoCloseWithMiddleMouseButton TabItemFlags = C.ImGuiTabItemFlags_NoCloseWithMiddleMouseButton
	TabItemFlagsNoPushId                     TabItemFlags = C.ImGuiTabItemFlags_NoPushId
	TabItemFlagsNoTooltip                    TabItemFlags = C.ImGuiTabItemFlags_NoTooltip
	TabItemFlagsNoReorder                    TabItemFlags = C.ImGuiTabItemFlags_NoReorder
	TabItemFlagsLeading                      TabItemFlags = C.ImGuiTabItemFlags_Leading
	TabItemFlagsTrailing                     TabItemFlags = C.ImGuiTabItemFlags_Trailing
	TabItemFlagsNoAssumedClosure             TabItemFlags = C.ImGuiTabItemFlags_NoAssumedClosure
)

// BeginTabBar opens a tab bar. Call [EndTabBar] only if it returns true.
func BeginTabBar(strID string, flags TabBarFlags) bool {
	cid := C.CString(strID)
	defer C.free(unsafe.Pointer(cid))
	return bool(C.igBeginTabBar(cid, C.ImGuiTabBarFlags(flags)))
}

// EndTabBar closes the tab bar opened by [BeginTabBar].
func EndTabBar() { C.igEndTabBar() }

// BeginTabItem opens a tab. pOpen may be nil; when non-nil it shows a close
// button and is updated with the tab's open state. Call [EndTabItem] only if it
// returns true.
func BeginTabItem(label string, pOpen *bool, flags TabItemFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	var ret C.bool
	withBoolPtr(pOpen, func(p *C.bool) {
		ret = C.igBeginTabItem(clabel, p, C.ImGuiTabItemFlags(flags))
	})
	return bool(ret)
}

// EndTabItem closes the tab opened by [BeginTabItem].
func EndTabItem() { C.igEndTabItem() }

// TabItemButton draws a tab that behaves like a button and reports whether it was
// clicked.
func TabItemButton(label string, flags TabItemFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igTabItemButton(clabel, C.ImGuiTabItemFlags(flags)))
}

// SetTabItemClosed notifies the tab bar that the named tab/window was closed
// externally this frame.
func SetTabItemClosed(label string) {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	C.igSetTabItemClosed(clabel)
}
