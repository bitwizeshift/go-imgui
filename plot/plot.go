// Package plot provides simple line and histogram plots.
package plot

import (
	"math"

	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// autoScale is Dear ImGui's sentinel for "compute the scale from the data".
const autoScale = math.MaxFloat32

// Lines plots Values as a line graph. When ScaleMin == ScaleMax the range is
// computed automatically.
type Lines struct {
	Label    string
	Values   []float32
	Overlay  string
	Size     imgui.Vec2
	ScaleMin float32
	ScaleMax float32
}

// NewLines returns a line plot of values.
func NewLines(label string, values []float32) *Lines { return &Lines{Label: label, Values: values} }

// Display draws the line plot.
func (l *Lines) Display() {
	lo, hi := scale(l.ScaleMin, l.ScaleMax)
	cimgui.PlotLines_FloatPtr(l.Label, l.Values, 0, l.Overlay, lo, hi, l.Size, 4)
}

// Histogram plots Values as a histogram. When ScaleMin == ScaleMax the range is
// computed automatically.
type Histogram struct {
	Label    string
	Values   []float32
	Overlay  string
	Size     imgui.Vec2
	ScaleMin float32
	ScaleMax float32
}

// NewHistogram returns a histogram of values.
func NewHistogram(label string, values []float32) *Histogram {
	return &Histogram{Label: label, Values: values}
}

// Display draws the histogram.
func (h *Histogram) Display() {
	lo, hi := scale(h.ScaleMin, h.ScaleMax)
	cimgui.PlotHistogram_FloatPtr(h.Label, h.Values, 0, h.Overlay, lo, hi, h.Size, 4)
}

func scale(lo, hi float32) (float32, float32) {
	if lo == hi {
		return autoScale, autoScale
	}
	return lo, hi
}
