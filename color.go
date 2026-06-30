package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// ColorEditOptions are the optional inputs to the color editors, pickers and
// swatch button. A nil *ColorEditOptions uses Dear ImGui's defaults; each field
// maps to an ImGuiColorEditFlags_ bit.
type ColorEditOptions struct {
	NoAlpha          bool // ImGuiColorEditFlags_NoAlpha
	NoPicker         bool // ImGuiColorEditFlags_NoPicker
	NoOptions        bool // ImGuiColorEditFlags_NoOptions
	NoSmallPreview   bool // ImGuiColorEditFlags_NoSmallPreview
	NoInputs         bool // ImGuiColorEditFlags_NoInputs
	NoTooltip        bool // ImGuiColorEditFlags_NoTooltip
	NoLabel          bool // ImGuiColorEditFlags_NoLabel
	NoSidePreview    bool // ImGuiColorEditFlags_NoSidePreview
	NoDragDrop       bool // ImGuiColorEditFlags_NoDragDrop
	NoBorder         bool // ImGuiColorEditFlags_NoBorder
	AlphaOpaque      bool // ImGuiColorEditFlags_AlphaOpaque
	AlphaNoBg        bool // ImGuiColorEditFlags_AlphaNoBg
	AlphaPreviewHalf bool // ImGuiColorEditFlags_AlphaPreviewHalf
	AlphaBar         bool // ImGuiColorEditFlags_AlphaBar
	HDR              bool // ImGuiColorEditFlags_HDR
	DisplayRGB       bool // ImGuiColorEditFlags_DisplayRGB
	DisplayHSV       bool // ImGuiColorEditFlags_DisplayHSV
	DisplayHex       bool // ImGuiColorEditFlags_DisplayHex
	Uint8            bool // ImGuiColorEditFlags_Uint8
	Float            bool // ImGuiColorEditFlags_Float
	PickerHueBar     bool // ImGuiColorEditFlags_PickerHueBar
	PickerHueWheel   bool // ImGuiColorEditFlags_PickerHueWheel
	InputRGB         bool // ImGuiColorEditFlags_InputRGB
	InputHSV         bool // ImGuiColorEditFlags_InputHSV
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *ColorEditOptions) flags() cimgui.ColorEditFlags {
	if o == nil {
		return cimgui.ColorEditFlagsNone
	}
	var f cimgui.ColorEditFlags
	if o.NoAlpha {
		f |= cimgui.ColorEditFlagsNoAlpha
	}
	if o.NoPicker {
		f |= cimgui.ColorEditFlagsNoPicker
	}
	if o.NoOptions {
		f |= cimgui.ColorEditFlagsNoOptions
	}
	if o.NoSmallPreview {
		f |= cimgui.ColorEditFlagsNoSmallPreview
	}
	if o.NoInputs {
		f |= cimgui.ColorEditFlagsNoInputs
	}
	if o.NoTooltip {
		f |= cimgui.ColorEditFlagsNoTooltip
	}
	if o.NoLabel {
		f |= cimgui.ColorEditFlagsNoLabel
	}
	if o.NoSidePreview {
		f |= cimgui.ColorEditFlagsNoSidePreview
	}
	if o.NoDragDrop {
		f |= cimgui.ColorEditFlagsNoDragDrop
	}
	if o.NoBorder {
		f |= cimgui.ColorEditFlagsNoBorder
	}
	if o.AlphaOpaque {
		f |= cimgui.ColorEditFlagsAlphaOpaque
	}
	if o.AlphaNoBg {
		f |= cimgui.ColorEditFlagsAlphaNoBg
	}
	if o.AlphaPreviewHalf {
		f |= cimgui.ColorEditFlagsAlphaPreviewHalf
	}
	if o.AlphaBar {
		f |= cimgui.ColorEditFlagsAlphaBar
	}
	if o.HDR {
		f |= cimgui.ColorEditFlagsHDR
	}
	if o.DisplayRGB {
		f |= cimgui.ColorEditFlagsDisplayRGB
	}
	if o.DisplayHSV {
		f |= cimgui.ColorEditFlagsDisplayHSV
	}
	if o.DisplayHex {
		f |= cimgui.ColorEditFlagsDisplayHex
	}
	if o.Uint8 {
		f |= cimgui.ColorEditFlagsUint8
	}
	if o.Float {
		f |= cimgui.ColorEditFlagsFloat
	}
	if o.PickerHueBar {
		f |= cimgui.ColorEditFlagsPickerHueBar
	}
	if o.PickerHueWheel {
		f |= cimgui.ColorEditFlagsPickerHueWheel
	}
	if o.InputRGB {
		f |= cimgui.ColorEditFlagsInputRGB
	}
	if o.InputHSV {
		f |= cimgui.ColorEditFlagsInputHSV
	}
	return f
}

// ColorEdit3 edits an RGB color in place and reports whether it changed. It
// models ImGui::ColorEdit3.
func ColorEdit3(label string, col *[3]float32, opts *ColorEditOptions) bool {
	return cimgui.ColorEdit3(label, col, opts.flags())
}

// ColorEdit4 edits an RGBA color in place and reports whether it changed. It
// models ImGui::ColorEdit4.
func ColorEdit4(label string, col *[4]float32, opts *ColorEditOptions) bool {
	return cimgui.ColorEdit4(label, col, opts.flags())
}

// ColorPicker3 shows an RGB color picker editing col in place. It models
// ImGui::ColorPicker3.
func ColorPicker3(label string, col *[3]float32, opts *ColorEditOptions) bool {
	return cimgui.ColorPicker3(label, col, opts.flags())
}

// ColorPicker4 shows an RGBA color picker editing col in place. refCol, when
// non-nil, supplies the reference color swatch. It models ImGui::ColorPicker4.
func ColorPicker4(label string, col *[4]float32, refCol *[4]float32, opts *ColorEditOptions) bool {
	return cimgui.ColorPicker4(label, col, opts.flags(), refCol)
}

// ColorButton draws a color swatch button and reports whether it was clicked. A
// zero size auto-fits. It models ImGui::ColorButton.
func ColorButton(descID string, col Vec4, size Vec2, opts *ColorEditOptions) bool {
	return cimgui.ColorButton(descID, col, opts.flags(), size)
}

// ColorConvertFloat4ToU32 packs an RGBA [Vec4] (components in 0..1) into a [U32].
// It models ImGui::ColorConvertFloat4ToU32.
func ColorConvertFloat4ToU32(col Vec4) U32 {
	return cimgui.ColorConvertFloat4ToU32(col)
}

// GetColorU32 returns the current style color for idx, scaled by alphaMul, as a
// [U32]. It models ImGui::GetColorU32 (the ImGuiCol overload).
func GetColorU32(idx Col, alphaMul float32) U32 {
	return cimgui.GetColorU32_Col(idx, alphaMul)
}
