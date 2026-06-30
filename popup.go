package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// PopupOptions are the optional inputs to [OpenPopup], the context-menu helpers
// and [IsPopupOpen]. A nil *PopupOptions uses Dear ImGui's defaults; each field
// maps to an ImGuiPopupFlags_ bit.
type PopupOptions struct {
	MouseButtonLeft         bool // ImGuiPopupFlags_MouseButtonLeft
	MouseButtonRight        bool // ImGuiPopupFlags_MouseButtonRight
	MouseButtonMiddle       bool // ImGuiPopupFlags_MouseButtonMiddle
	NoReopen                bool // ImGuiPopupFlags_NoReopen
	NoOpenOverExistingPopup bool // ImGuiPopupFlags_NoOpenOverExistingPopup
	NoOpenOverItems         bool // ImGuiPopupFlags_NoOpenOverItems
	AnyPopupID              bool // ImGuiPopupFlags_AnyPopupId
	AnyPopupLevel           bool // ImGuiPopupFlags_AnyPopupLevel
	AnyPopup                bool // ImGuiPopupFlags_AnyPopup (composite)
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *PopupOptions) flags() cimgui.PopupFlags {
	if o == nil {
		return cimgui.PopupFlagsNone
	}
	var f cimgui.PopupFlags
	if o.MouseButtonLeft {
		f |= cimgui.PopupFlagsMouseButtonLeft
	}
	if o.MouseButtonRight {
		f |= cimgui.PopupFlagsMouseButtonRight
	}
	if o.MouseButtonMiddle {
		f |= cimgui.PopupFlagsMouseButtonMiddle
	}
	if o.NoReopen {
		f |= cimgui.PopupFlagsNoReopen
	}
	if o.NoOpenOverExistingPopup {
		f |= cimgui.PopupFlagsNoOpenOverExistingPopup
	}
	if o.NoOpenOverItems {
		f |= cimgui.PopupFlagsNoOpenOverItems
	}
	if o.AnyPopupID {
		f |= cimgui.PopupFlagsAnyPopupId
	}
	if o.AnyPopupLevel {
		f |= cimgui.PopupFlagsAnyPopupLevel
	}
	if o.AnyPopup {
		f |= cimgui.PopupFlagsAnyPopup
	}
	return f
}

// OpenPopup marks the popup identified by id (a string label or precomputed
// uint32 id) to open on the next frame. It models ImGui::OpenPopup.
func OpenPopup[T ID](id T, opts *PopupOptions) {
	switch v := any(id).(type) {
	case string:
		cimgui.OpenPopup_Str(v, opts.flags())
	case uint32:
		cimgui.OpenPopup_ID(v, opts.flags())
	}
}

// Popup begins the popup identified by strID if it has been marked open. It
// models ImGui::BeginPopup. open reports whether the popup is open; the returned
// [EndFunc] (ImGui::EndPopup) ends it only when open.
func Popup(strID string, opts *WindowOptions) (open bool, end EndFunc) {
	open = cimgui.BeginPopup(strID, opts.flags())
	if !open {
		return open, func() {}
	}
	return open, cimgui.EndPopup
}

// PopupModal begins a modal popup named name. It models ImGui::BeginPopupModal.
// When open is non-nil a close button is shown and *open is updated; visible
// reports whether the popup is open. The returned [EndFunc] (ImGui::EndPopup)
// ends it only when visible.
func PopupModal(name string, open *bool, opts *WindowOptions) (visible bool, end EndFunc) {
	visible = cimgui.BeginPopupModal(name, open, opts.flags())
	if !visible {
		return visible, func() {}
	}
	return visible, cimgui.EndPopup
}

// PopupContextItem begins a popup on right-click of the previous item. An empty
// strID reuses the previous item's ID. It models ImGui::BeginPopupContextItem.
// open reports whether the popup is open; the returned [EndFunc] ends it only
// when open.
func PopupContextItem(strID string, opts *PopupOptions) (open bool, end EndFunc) {
	open = cimgui.BeginPopupContextItem(strID, opts.flags())
	if !open {
		return open, func() {}
	}
	return open, cimgui.EndPopup
}

// PopupContextWindow begins a popup on right-click of the current window. It
// models ImGui::BeginPopupContextWindow. open reports whether the popup is open;
// the returned [EndFunc] ends it only when open.
func PopupContextWindow(strID string, opts *PopupOptions) (open bool, end EndFunc) {
	open = cimgui.BeginPopupContextWindow(strID, opts.flags())
	if !open {
		return open, func() {}
	}
	return open, cimgui.EndPopup
}

// PopupContextVoid begins a popup on right-click of empty space (no window). It
// models ImGui::BeginPopupContextVoid. open reports whether the popup is open;
// the returned [EndFunc] ends it only when open.
func PopupContextVoid(strID string, opts *PopupOptions) (open bool, end EndFunc) {
	open = cimgui.BeginPopupContextVoid(strID, opts.flags())
	if !open {
		return open, func() {}
	}
	return open, cimgui.EndPopup
}

// CloseCurrentPopup closes the popup currently being drawn. It models
// ImGui::CloseCurrentPopup.
func CloseCurrentPopup() {
	cimgui.CloseCurrentPopup()
}

// IsPopupOpen reports whether the popup with the given string ID is open. It
// models ImGui::IsPopupOpen.
func IsPopupOpen(strID string, opts *PopupOptions) bool {
	return cimgui.IsPopupOpen_Str(strID, opts.flags())
}
