package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// PlotLines draws a line plot of values. A zero graphSize auto-fits; equal
// scaleMin and scaleMax auto-scale; overlayText may be empty. valuesOffset
// rotates the starting index and stride is the byte stride between samples (use 4
// for a packed []float32). It models ImGui::PlotLines (the values overload).
func PlotLines(label string, values []float32, valuesOffset int32, overlayText string, scaleMin, scaleMax float32, graphSize Vec2, stride int32) {
	cimgui.PlotLines_FloatPtr(label, values, valuesOffset, overlayText, scaleMin, scaleMax, graphSize, stride)
}

// PlotLinesFunc draws a line plot of valuesCount samples produced lazily by
// getter. See [PlotLines] for the remaining parameters. It models ImGui::PlotLines
// (the getter overload).
func PlotLinesFunc(label string, getter func(idx int32) float32, valuesCount, valuesOffset int32, overlayText string, scaleMin, scaleMax float32, graphSize Vec2) {
	cimgui.PlotLines_FnFloatPtr(label, getter, valuesCount, valuesOffset, overlayText, scaleMin, scaleMax, graphSize)
}

// PlotHistogram draws a histogram of values. See [PlotLines] for the meaning of
// the remaining parameters. It models ImGui::PlotHistogram (the values overload).
func PlotHistogram(label string, values []float32, valuesOffset int32, overlayText string, scaleMin, scaleMax float32, graphSize Vec2, stride int32) {
	cimgui.PlotHistogram_FloatPtr(label, values, valuesOffset, overlayText, scaleMin, scaleMax, graphSize, stride)
}

// PlotHistogramFunc draws a histogram of valuesCount samples produced lazily by
// getter. See [PlotLines] for the remaining parameters. It models
// ImGui::PlotHistogram (the getter overload).
func PlotHistogramFunc(label string, getter func(idx int32) float32, valuesCount, valuesOffset int32, overlayText string, scaleMin, scaleMax float32, graphSize Vec2) {
	cimgui.PlotHistogram_FnFloatPtr(label, getter, valuesCount, valuesOffset, overlayText, scaleMin, scaleMax, graphSize)
}
