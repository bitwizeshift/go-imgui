// Package debug provides Dear ImGui's built-in demo and diagnostic windows as
// widgets.
package debug

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// DemoWindow shows the Dear ImGui demo window, a live reference for the whole
// library. When Open is non-nil it is cleared when the window is closed.
type DemoWindow struct {
	Open *bool
}

// NewDemoWindow returns a demo-window widget.
func NewDemoWindow(open *bool) *DemoWindow { return &DemoWindow{Open: open} }

// Display draws the demo window.
func (d *DemoWindow) Display() { cimgui.ShowDemoWindow(d.Open) }

// MetricsWindow shows the metrics/debugger window.
type MetricsWindow struct {
	Open *bool
}

// NewMetricsWindow returns a metrics-window widget.
func NewMetricsWindow(open *bool) *MetricsWindow { return &MetricsWindow{Open: open} }

// Display draws the metrics window.
func (m *MetricsWindow) Display() { cimgui.ShowMetricsWindow(m.Open) }

// AboutWindow shows the about window.
type AboutWindow struct {
	Open *bool
}

// NewAboutWindow returns an about-window widget.
func NewAboutWindow(open *bool) *AboutWindow { return &AboutWindow{Open: open} }

// Display draws the about window.
func (a *AboutWindow) Display() { cimgui.ShowAboutWindow(a.Open) }
