package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// PopupFlags configures popup queries and context-menu helpers. Mirrors the
// public ImGuiPopupFlags_.
type PopupFlags int32

const (
	PopupFlagsNone                  PopupFlags = C.ImGuiPopupFlags_None
	PopupFlagsMouseButtonLeft       PopupFlags = C.ImGuiPopupFlags_MouseButtonLeft
	PopupFlagsMouseButtonRight      PopupFlags = C.ImGuiPopupFlags_MouseButtonRight
	PopupFlagsMouseButtonMiddle     PopupFlags = C.ImGuiPopupFlags_MouseButtonMiddle
	PopupFlagsNoReopen              PopupFlags = C.ImGuiPopupFlags_NoReopen
	PopupFlagsNoOpenOverExistingPopup PopupFlags = C.ImGuiPopupFlags_NoOpenOverExistingPopup
	PopupFlagsNoOpenOverItems       PopupFlags = C.ImGuiPopupFlags_NoOpenOverItems
	PopupFlagsAnyPopupId            PopupFlags = C.ImGuiPopupFlags_AnyPopupId
	PopupFlagsAnyPopupLevel         PopupFlags = C.ImGuiPopupFlags_AnyPopupLevel
	PopupFlagsAnyPopup              PopupFlags = C.ImGuiPopupFlags_AnyPopup
)

// OpenPopup_Str marks the popup with the given string ID to open.
func OpenPopup_Str(strID string, popupFlags PopupFlags) {
	cid := C.CString(strID)
	defer C.free(unsafe.Pointer(cid))
	C.igOpenPopup_Str(cid, C.ImGuiPopupFlags(popupFlags))
}

// OpenPopup_ID marks the popup with the given numeric ID to open.
func OpenPopup_ID(id uint32, popupFlags PopupFlags) {
	C.igOpenPopup_ID(C.ImGuiID(id), C.ImGuiPopupFlags(popupFlags))
}

// BeginPopup opens a popup window if it has been marked open. Call [EndPopup]
// only if it returns true.
func BeginPopup(strID string, flags WindowFlags) bool {
	cid := C.CString(strID)
	defer C.free(unsafe.Pointer(cid))
	return bool(C.igBeginPopup(cid, C.ImGuiWindowFlags(flags)))
}

// BeginPopupModal opens a modal popup. pOpen may be nil; when non-nil it shows a
// close button and is updated with the open state. Call [EndPopup] only if it
// returns true.
func BeginPopupModal(name string, pOpen *bool, flags WindowFlags) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var ret C.bool
	withBoolPtr(pOpen, func(p *C.bool) {
		ret = C.igBeginPopupModal(cname, p, C.ImGuiWindowFlags(flags))
	})
	return bool(ret)
}

// EndPopup closes the popup opened by [BeginPopup] or [BeginPopupModal].
func EndPopup() { C.igEndPopup() }

// BeginPopupContextItem opens a popup on right-click of the previous item. An
// empty strID reuses the previous item's ID. Call [EndPopup] only if it returns
// true.
func BeginPopupContextItem(strID string, popupFlags PopupFlags) bool {
	var cid *C.char
	if strID != "" {
		cid = C.CString(strID)
		defer C.free(unsafe.Pointer(cid))
	}
	return bool(C.igBeginPopupContextItem(cid, C.ImGuiPopupFlags(popupFlags)))
}

// BeginPopupContextWindow opens a popup on right-click of the current window.
func BeginPopupContextWindow(strID string, popupFlags PopupFlags) bool {
	var cid *C.char
	if strID != "" {
		cid = C.CString(strID)
		defer C.free(unsafe.Pointer(cid))
	}
	return bool(C.igBeginPopupContextWindow(cid, C.ImGuiPopupFlags(popupFlags)))
}

// BeginPopupContextVoid opens a popup on right-click of empty space (no window).
func BeginPopupContextVoid(strID string, popupFlags PopupFlags) bool {
	var cid *C.char
	if strID != "" {
		cid = C.CString(strID)
		defer C.free(unsafe.Pointer(cid))
	}
	return bool(C.igBeginPopupContextVoid(cid, C.ImGuiPopupFlags(popupFlags)))
}

// CloseCurrentPopup closes the popup currently being drawn.
func CloseCurrentPopup() { C.igCloseCurrentPopup() }

// IsPopupOpen_Str reports whether the popup with the given string ID is open.
func IsPopupOpen_Str(strID string, flags PopupFlags) bool {
	cid := C.CString(strID)
	defer C.free(unsafe.Pointer(cid))
	return bool(C.igIsPopupOpen_Str(cid, C.ImGuiPopupFlags(flags)))
}
