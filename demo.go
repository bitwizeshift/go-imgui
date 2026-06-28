package imgui

import "rodusek.dev/pkg/imgui/internal/cimgui"

// ShowDemoWindow displays the Dear ImGui demo window, a live reference for the
// whole API. open may be nil; when non-nil, closing the window sets *open false.
func ShowDemoWindow(open *bool) { cimgui.ShowDemoWindow(open) }

// ShowMetricsWindow displays the metrics/debugger window. open may be nil.
func ShowMetricsWindow(open *bool) { cimgui.ShowMetricsWindow(open) }

// ShowAboutWindow displays the about window. open may be nil.
func ShowAboutWindow(open *bool) { cimgui.ShowAboutWindow(open) }
