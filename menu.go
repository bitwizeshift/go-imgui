package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// MenuBar appends to the menu bar of the current window, which must have been
// opened with WindowOptions.MenuBar set. It models ImGui::BeginMenuBar. open
// reports whether the bar is visible; the returned [EndFunc] (ImGui::EndMenuBar)
// ends it only when open.
func MenuBar() (open bool, end EndFunc) {
	open = cimgui.BeginMenuBar()
	if !open {
		return open, func() {}
	}
	return open, cimgui.EndMenuBar
}

// MainMenuBar opens a full-screen menu bar at the top of the viewport. It models
// ImGui::BeginMainMenuBar. open reports whether the bar is visible; the returned
// [EndFunc] (ImGui::EndMainMenuBar) ends it only when open.
func MainMenuBar() (open bool, end EndFunc) {
	open = cimgui.BeginMainMenuBar()
	if !open {
		return open, func() {}
	}
	return open, cimgui.EndMainMenuBar
}

// Menu opens a sub-menu entry labelled label. It models ImGui::BeginMenu. open
// reports whether the menu is expanded; the returned [EndFunc] (ImGui::EndMenu)
// ends it only when open.
func Menu(label string, enabled bool) (open bool, end EndFunc) {
	open = cimgui.BeginMenu(label, enabled)
	if !open {
		return open, func() {}
	}
	return open, cimgui.EndMenu
}

// MenuItem draws a menu item and reports whether it was activated. shortcut may
// be empty. selected is either a bool (rendered checked) or a *bool (toggled on
// activation); the form is inferred at the call site. It models ImGui::MenuItem.
func MenuItem[T BoolOrPtr](label, shortcut string, selected T, enabled bool) bool {
	switch v := any(selected).(type) {
	case bool:
		return cimgui.MenuItem_Bool(label, shortcut, v, enabled)
	case *bool:
		return cimgui.MenuItem_BoolPtr(label, shortcut, v, enabled)
	}
	return false
}
