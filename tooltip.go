package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// Tooltip begins a tooltip window. It models ImGui::BeginTooltip. open reports
// whether the tooltip is being drawn; the returned [EndFunc] (ImGui::EndTooltip)
// ends it only when open.
func Tooltip() (open bool, end EndFunc) {
	open = cimgui.BeginTooltip()
	if !open {
		return open, func() {}
	}
	return open, cimgui.EndTooltip
}

// ItemTooltip begins a tooltip only when the previous item is hovered. It models
// ImGui::BeginItemTooltip. open reports whether the tooltip is being drawn; the
// returned [EndFunc] (ImGui::EndTooltip) ends it only when open.
func ItemTooltip() (open bool, end EndFunc) {
	open = cimgui.BeginItemTooltip()
	if !open {
		return open, func() {}
	}
	return open, cimgui.EndTooltip
}

// SetTooltip sets the contents of a tooltip shown while the previous item is
// hovered. It models ImGui::SetTooltip.
func SetTooltip(text string) {
	cimgui.SetTooltip(text)
}
