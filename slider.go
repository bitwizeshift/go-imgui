package imgui

import (
	"unsafe"

	"github.com/bitwizeshift/go-imgui/internal/cimgui"
)

// SliderOptions are the optional inputs to the slider and drag widgets. A nil
// *SliderOptions uses Dear ImGui's defaults. Format overrides the printf-style
// display format (empty selects each widget's default); FormatMax does the same
// for the upper bound of the range widgets ([DragFloatRange2], [DragIntRange2]).
// The remaining fields map to ImGuiSliderFlags_ bits.
type SliderOptions struct {
	Format          string // display format; empty uses the widget default
	FormatMax       string // upper-bound format for range widgets
	Logarithmic     bool   // ImGuiSliderFlags_Logarithmic
	NoRoundToFormat bool   // ImGuiSliderFlags_NoRoundToFormat
	NoInput         bool   // ImGuiSliderFlags_NoInput
	WrapAround      bool   // ImGuiSliderFlags_WrapAround
	ClampOnInput    bool   // ImGuiSliderFlags_ClampOnInput
	ClampZeroRange  bool   // ImGuiSliderFlags_ClampZeroRange
	NoSpeedTweaks   bool   // ImGuiSliderFlags_NoSpeedTweaks
	AlwaysClamp     bool   // ImGuiSliderFlags_AlwaysClamp
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *SliderOptions) flags() cimgui.SliderFlags {
	if o == nil {
		return cimgui.SliderFlagsNone
	}
	var f cimgui.SliderFlags
	if o.Logarithmic {
		f |= cimgui.SliderFlagsLogarithmic
	}
	if o.NoRoundToFormat {
		f |= cimgui.SliderFlagsNoRoundToFormat
	}
	if o.NoInput {
		f |= cimgui.SliderFlagsNoInput
	}
	if o.WrapAround {
		f |= cimgui.SliderFlagsWrapAround
	}
	if o.ClampOnInput {
		f |= cimgui.SliderFlagsClampOnInput
	}
	if o.ClampZeroRange {
		f |= cimgui.SliderFlagsClampZeroRange
	}
	if o.NoSpeedTweaks {
		f |= cimgui.SliderFlagsNoSpeedTweaks
	}
	if o.AlwaysClamp {
		f |= cimgui.SliderFlagsAlwaysClamp
	}
	return f
}

// format returns o.Format, or def when o is nil or Format is empty.
func (o *SliderOptions) format(def string) string {
	if o == nil || o.Format == "" {
		return def
	}
	return o.Format
}

// formatMax returns o.FormatMax, or def when o is nil or FormatMax is empty.
func (o *SliderOptions) formatMax(def string) string {
	if o == nil || o.FormatMax == "" {
		return def
	}
	return o.FormatMax
}

// ---- Sliders -------------------------------------------------------------

// SliderFloat draws a float slider bound to v, constrained to [vMin,vMax], and
// reports whether it changed. It models ImGui::SliderFloat.
func SliderFloat(label string, v *float32, vMin, vMax float32, opts *SliderOptions) bool {
	return cimgui.SliderFloat(label, v, vMin, vMax, opts.format("%.3f"), opts.flags())
}

// SliderFloat2 draws a 2-component float slider bound to v. It models
// ImGui::SliderFloat2.
func SliderFloat2(label string, v *[2]float32, vMin, vMax float32, opts *SliderOptions) bool {
	return cimgui.SliderFloat2(label, v, vMin, vMax, opts.format("%.3f"), opts.flags())
}

// SliderFloat3 draws a 3-component float slider bound to v. It models
// ImGui::SliderFloat3.
func SliderFloat3(label string, v *[3]float32, vMin, vMax float32, opts *SliderOptions) bool {
	return cimgui.SliderFloat3(label, v, vMin, vMax, opts.format("%.3f"), opts.flags())
}

// SliderFloat4 draws a 4-component float slider bound to v. It models
// ImGui::SliderFloat4.
func SliderFloat4(label string, v *[4]float32, vMin, vMax float32, opts *SliderOptions) bool {
	return cimgui.SliderFloat4(label, v, vMin, vMax, opts.format("%.3f"), opts.flags())
}

// SliderAngle draws a slider editing vRad (radians) shown in degrees, clamped to
// [vDegreesMin,vDegreesMax]. It models ImGui::SliderAngle.
func SliderAngle(label string, vRad *float32, vDegreesMin, vDegreesMax float32, opts *SliderOptions) bool {
	return cimgui.SliderAngle(label, vRad, vDegreesMin, vDegreesMax, opts.format("%.0f deg"), opts.flags())
}

// SliderInt draws an int slider bound to v, constrained to [vMin,vMax], and
// reports whether it changed. It models ImGui::SliderInt.
func SliderInt(label string, v *int32, vMin, vMax int32, opts *SliderOptions) bool {
	return cimgui.SliderInt(label, v, vMin, vMax, opts.format("%d"), opts.flags())
}

// SliderInt2 draws a 2-component int slider bound to v. It models
// ImGui::SliderInt2.
func SliderInt2(label string, v *[2]int32, vMin, vMax int32, opts *SliderOptions) bool {
	return cimgui.SliderInt2(label, v, vMin, vMax, opts.format("%d"), opts.flags())
}

// SliderInt3 draws a 3-component int slider bound to v. It models
// ImGui::SliderInt3.
func SliderInt3(label string, v *[3]int32, vMin, vMax int32, opts *SliderOptions) bool {
	return cimgui.SliderInt3(label, v, vMin, vMax, opts.format("%d"), opts.flags())
}

// SliderInt4 draws a 4-component int slider bound to v. It models
// ImGui::SliderInt4.
func SliderInt4(label string, v *[4]int32, vMin, vMax int32, opts *SliderOptions) bool {
	return cimgui.SliderInt4(label, v, vMin, vMax, opts.format("%d"), opts.flags())
}

// SliderScalar draws a slider for an arbitrary data type. pData, pMin and pMax
// point to values of dataType. It models ImGui::SliderScalar.
func SliderScalar(label string, dataType DataType, pData, pMin, pMax unsafe.Pointer, opts *SliderOptions) bool {
	return cimgui.SliderScalar(label, dataType, pData, pMin, pMax, opts.format(""), opts.flags())
}

// SliderScalarN draws a slider editing components values of dataType stored
// contiguously at pData. It models ImGui::SliderScalarN.
func SliderScalarN(label string, dataType DataType, pData unsafe.Pointer, components int32, pMin, pMax unsafe.Pointer, opts *SliderOptions) bool {
	return cimgui.SliderScalarN(label, dataType, pData, components, pMin, pMax, opts.format(""), opts.flags())
}

// VSliderFloat draws a vertical float slider of the given size bound to v. It
// models ImGui::VSliderFloat.
func VSliderFloat(label string, size Vec2, v *float32, vMin, vMax float32, opts *SliderOptions) bool {
	return cimgui.VSliderFloat(label, size, v, vMin, vMax, opts.format("%.3f"), opts.flags())
}

// VSliderInt draws a vertical int slider of the given size bound to v. It models
// ImGui::VSliderInt.
func VSliderInt(label string, size Vec2, v *int32, vMin, vMax int32, opts *SliderOptions) bool {
	return cimgui.VSliderInt(label, size, v, vMin, vMax, opts.format("%d"), opts.flags())
}

// VSliderScalar draws a vertical slider of the given size for an arbitrary data
// type. It models ImGui::VSliderScalar.
func VSliderScalar(label string, size Vec2, dataType DataType, pData, pMin, pMax unsafe.Pointer, opts *SliderOptions) bool {
	return cimgui.VSliderScalar(label, size, dataType, pData, pMin, pMax, opts.format(""), opts.flags())
}

// ---- Drags ---------------------------------------------------------------

// DragFloat draws a draggable float bound to v with the given drag speed,
// soft-clamped to [vMin,vMax], and reports whether it changed. It models
// ImGui::DragFloat.
func DragFloat(label string, v *float32, speed, vMin, vMax float32, opts *SliderOptions) bool {
	return cimgui.DragFloat(label, v, speed, vMin, vMax, opts.format("%.3f"), opts.flags())
}

// DragFloat2 draws a draggable 2-component float bound to v. It models
// ImGui::DragFloat2.
func DragFloat2(label string, v *[2]float32, speed, vMin, vMax float32, opts *SliderOptions) bool {
	return cimgui.DragFloat2(label, v, speed, vMin, vMax, opts.format("%.3f"), opts.flags())
}

// DragFloat3 draws a draggable 3-component float bound to v. It models
// ImGui::DragFloat3.
func DragFloat3(label string, v *[3]float32, speed, vMin, vMax float32, opts *SliderOptions) bool {
	return cimgui.DragFloat3(label, v, speed, vMin, vMax, opts.format("%.3f"), opts.flags())
}

// DragFloat4 draws a draggable 4-component float bound to v. It models
// ImGui::DragFloat4.
func DragFloat4(label string, v *[4]float32, speed, vMin, vMax float32, opts *SliderOptions) bool {
	return cimgui.DragFloat4(label, v, speed, vMin, vMax, opts.format("%.3f"), opts.flags())
}

// DragInt draws a draggable int bound to v with the given drag speed,
// soft-clamped to [vMin,vMax], and reports whether it changed. It models
// ImGui::DragInt.
func DragInt(label string, v *int32, speed float32, vMin, vMax int32, opts *SliderOptions) bool {
	return cimgui.DragInt(label, v, speed, vMin, vMax, opts.format("%d"), opts.flags())
}

// DragInt2 draws a draggable 2-component int bound to v. It models
// ImGui::DragInt2.
func DragInt2(label string, v *[2]int32, speed float32, vMin, vMax int32, opts *SliderOptions) bool {
	return cimgui.DragInt2(label, v, speed, vMin, vMax, opts.format("%d"), opts.flags())
}

// DragInt3 draws a draggable 3-component int bound to v. It models
// ImGui::DragInt3.
func DragInt3(label string, v *[3]int32, speed float32, vMin, vMax int32, opts *SliderOptions) bool {
	return cimgui.DragInt3(label, v, speed, vMin, vMax, opts.format("%d"), opts.flags())
}

// DragInt4 draws a draggable 4-component int bound to v. It models
// ImGui::DragInt4.
func DragInt4(label string, v *[4]int32, speed float32, vMin, vMax int32, opts *SliderOptions) bool {
	return cimgui.DragInt4(label, v, speed, vMin, vMax, opts.format("%d"), opts.flags())
}

// DragFloatRange2 draws two draggable floats editing the range
// [vCurrentMin,vCurrentMax]. It models ImGui::DragFloatRange2.
func DragFloatRange2(label string, vCurrentMin, vCurrentMax *float32, speed, vMin, vMax float32, opts *SliderOptions) bool {
	format := opts.format("%.3f")
	return cimgui.DragFloatRange2(label, vCurrentMin, vCurrentMax, speed, vMin, vMax, format, opts.formatMax(format), opts.flags())
}

// DragIntRange2 draws two draggable ints editing the range
// [vCurrentMin,vCurrentMax]. It models ImGui::DragIntRange2.
func DragIntRange2(label string, vCurrentMin, vCurrentMax *int32, speed float32, vMin, vMax int32, opts *SliderOptions) bool {
	format := opts.format("%d")
	return cimgui.DragIntRange2(label, vCurrentMin, vCurrentMax, speed, vMin, vMax, format, opts.formatMax(format), opts.flags())
}

// DragScalar draws a draggable widget for an arbitrary data type. It models
// ImGui::DragScalar.
func DragScalar(label string, dataType DataType, pData unsafe.Pointer, speed float32, pMin, pMax unsafe.Pointer, opts *SliderOptions) bool {
	return cimgui.DragScalar(label, dataType, pData, speed, pMin, pMax, opts.format(""), opts.flags())
}

// DragScalarN draws draggable widgets for components values of dataType stored
// contiguously at pData. It models ImGui::DragScalarN.
func DragScalarN(label string, dataType DataType, pData unsafe.Pointer, components int32, speed float32, pMin, pMax unsafe.Pointer, opts *SliderOptions) bool {
	return cimgui.DragScalarN(label, dataType, pData, components, speed, pMin, pMax, opts.format(""), opts.flags())
}
