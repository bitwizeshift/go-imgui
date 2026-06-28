package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

// MouseButton identifies a mouse button. Mirrors ImGuiMouseButton_.
type MouseButton int32

const (
	MouseButtonLeft   MouseButton = C.ImGuiMouseButton_Left
	MouseButtonRight  MouseButton = C.ImGuiMouseButton_Right
	MouseButtonMiddle MouseButton = C.ImGuiMouseButton_Middle
)

// HoveredFlags configures [IsItemHovered]. Mirrors the public ImGuiHoveredFlags_.
type HoveredFlags int32

const (
	HoveredFlagsNone                         HoveredFlags = C.ImGuiHoveredFlags_None
	HoveredFlagsChildWindows                 HoveredFlags = C.ImGuiHoveredFlags_ChildWindows
	HoveredFlagsRootWindow                   HoveredFlags = C.ImGuiHoveredFlags_RootWindow
	HoveredFlagsAnyWindow                    HoveredFlags = C.ImGuiHoveredFlags_AnyWindow
	HoveredFlagsNoPopupHierarchy             HoveredFlags = C.ImGuiHoveredFlags_NoPopupHierarchy
	HoveredFlagsAllowWhenBlockedByPopup      HoveredFlags = C.ImGuiHoveredFlags_AllowWhenBlockedByPopup
	HoveredFlagsAllowWhenBlockedByActiveItem HoveredFlags = C.ImGuiHoveredFlags_AllowWhenBlockedByActiveItem
	HoveredFlagsAllowWhenOverlapped          HoveredFlags = C.ImGuiHoveredFlags_AllowWhenOverlapped
	HoveredFlagsAllowWhenDisabled            HoveredFlags = C.ImGuiHoveredFlags_AllowWhenDisabled
	HoveredFlagsRectOnly                     HoveredFlags = C.ImGuiHoveredFlags_RectOnly
	HoveredFlagsRootAndChildWindows          HoveredFlags = C.ImGuiHoveredFlags_RootAndChildWindows
	HoveredFlagsForTooltip                   HoveredFlags = C.ImGuiHoveredFlags_ForTooltip
	HoveredFlagsStationary                   HoveredFlags = C.ImGuiHoveredFlags_Stationary
	HoveredFlagsDelayNone                    HoveredFlags = C.ImGuiHoveredFlags_DelayNone
	HoveredFlagsDelayShort                   HoveredFlags = C.ImGuiHoveredFlags_DelayShort
	HoveredFlagsDelayNormal                  HoveredFlags = C.ImGuiHoveredFlags_DelayNormal
)

// IsItemHovered reports whether the previous item is hovered.
func IsItemHovered(flags HoveredFlags) bool {
	return bool(C.igIsItemHovered(C.ImGuiHoveredFlags(flags)))
}

// IsItemActive reports whether the previous item is active (e.g. held).
func IsItemActive() bool { return bool(C.igIsItemActive()) }

// IsItemClicked reports whether the previous item was clicked with the given
// mouse button.
func IsItemClicked(button MouseButton) bool {
	return bool(C.igIsItemClicked(C.ImGuiMouseButton(button)))
}

// SetItemDefaultFocus makes the previous item the default focused one when the
// window appears.
func SetItemDefaultFocus() { C.igSetItemDefaultFocus() }

// SetKeyboardFocusHere focuses the next item (offset 0) or a later item.
func SetKeyboardFocusHere(offset int32) { C.igSetKeyboardFocusHere(C.int(offset)) }

// PushItemWidth pushes the width of common widgets; pair with [PopItemWidth].
func PushItemWidth(itemWidth float32) { C.igPushItemWidth(C.float(itemWidth)) }

// PopItemWidth restores the width pushed by [PushItemWidth].
func PopItemWidth() { C.igPopItemWidth() }

// SetNextItemWidth sets the width of the next common widget.
func SetNextItemWidth(itemWidth float32) { C.igSetNextItemWidth(C.float(itemWidth)) }

// BeginDisabled begins a disabled block when disabled is true; pair with
// [EndDisabled].
func BeginDisabled(disabled bool) { C.igBeginDisabled(C.bool(disabled)) }

// EndDisabled ends the block opened by [BeginDisabled].
func EndDisabled() { C.igEndDisabled() }
