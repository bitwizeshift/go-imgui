package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// ShowDemoWindow displays the Dear ImGui demo window. pOpen may be nil; when
// non-nil it shows a close button and is updated with the open state. It models
// ImGui::ShowDemoWindow.
func ShowDemoWindow(pOpen *bool) {
	cimgui.ShowDemoWindow(pOpen)
}

// ShowMetricsWindow displays the metrics/debug window. pOpen may be nil. It
// models ImGui::ShowMetricsWindow.
func ShowMetricsWindow(pOpen *bool) {
	cimgui.ShowMetricsWindow(pOpen)
}

// ShowAboutWindow displays the about window. pOpen may be nil. It models
// ImGui::ShowAboutWindow.
func ShowAboutWindow(pOpen *bool) {
	cimgui.ShowAboutWindow(pOpen)
}
