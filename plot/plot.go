// Package plot provides simple line and histogram plots.
package plot

import (
	"math"

	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// autoScale is Dear ImGui's sentinel for "compute the scale from the data".
const autoScale = math.MaxFloat32

// Lines plots a line graph. The samples come from Values, or from Getter (over
// Count samples) when it is set. When ScaleMin == ScaleMax the range is computed
// automatically.
type Lines struct {
	Label    string
	Values   []float32
	Getter   func(i int32) float32 // when set, supplies Count samples lazily
	Count    int32
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
	if l.Getter != nil {
		cimgui.PlotLines_FnFloatPtr(l.Label, l.Getter, l.Count, 0, l.Overlay, lo, hi, l.Size)
		return
	}
	cimgui.PlotLines_FloatPtr(l.Label, l.Values, 0, l.Overlay, lo, hi, l.Size, 4)
}

// Histogram plots a histogram. The samples come from Values, or from Getter (over
// Count samples) when it is set. When ScaleMin == ScaleMax the range is computed
// automatically.
type Histogram struct {
	Label    string
	Values   []float32
	Getter   func(i int32) float32 // when set, supplies Count samples lazily
	Count    int32
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
	if h.Getter != nil {
		cimgui.PlotHistogram_FnFloatPtr(h.Label, h.Getter, h.Count, 0, h.Overlay, lo, hi, h.Size)
		return
	}
	cimgui.PlotHistogram_FloatPtr(h.Label, h.Values, 0, h.Overlay, lo, hi, h.Size, 4)
}

func scale(lo, hi float32) (float32, float32) {
	if lo == hi {
		return autoScale, autoScale
	}
	return lo, hi
}
