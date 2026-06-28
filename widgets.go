package imgui

import "rodusek.dev/pkg/imgui/internal/cimgui"

// Button draws a button and reports whether it was clicked this frame.
func Button(label string) bool { return cimgui.Button(label, 0, 0) }

// ButtonV draws a button with an explicit size. A zero size component auto-fits
// the label on that axis.
func ButtonV(label string, size Vec2) bool { return cimgui.Button(label, size.X, size.Y) }

// SmallButton draws a button with no frame padding.
func SmallButton(label string) bool { return cimgui.SmallButton(label) }

// Checkbox draws a checkbox bound to v and reports whether it changed.
func Checkbox(label string, v *bool) bool { return cimgui.Checkbox(label, v) }

// RadioButton draws a radio button and reports whether it was clicked.
func RadioButton(label string, active bool) bool { return cimgui.RadioButton(label, active) }

// SliderFloat draws a float slider bound to v and reports whether it changed.
func SliderFloat(label string, v *float32, min, max float32) bool {
	return cimgui.SliderFloat(label, v, min, max, "%.3f", 0)
}

// SliderFloatV draws a float slider with a custom display format and flags.
func SliderFloatV(label string, v *float32, min, max float32, format string, flags SliderFlags) bool {
	return cimgui.SliderFloat(label, v, min, max, format, int(flags))
}

// SliderInt draws an int slider bound to v and reports whether it changed.
func SliderInt(label string, v *int32, min, max int32) bool {
	return cimgui.SliderInt(label, v, min, max, "%d", 0)
}

// SliderIntV draws an int slider with a custom display format and flags.
func SliderIntV(label string, v *int32, min, max int32, format string, flags SliderFlags) bool {
	return cimgui.SliderInt(label, v, min, max, format, int(flags))
}

// InputFloat draws a float input bound to v and reports whether it changed.
func InputFloat(label string, v *float32) bool {
	return cimgui.InputFloat(label, v, 0, 0, "%.3f", 0)
}

// InputFloatV draws a float input with step buttons, format, and flags.
func InputFloatV(label string, v *float32, step, stepFast float32, format string, flags InputTextFlags) bool {
	return cimgui.InputFloat(label, v, step, stepFast, format, int(flags))
}

// InputInt draws an int input bound to v and reports whether it changed.
func InputInt(label string, v *int32) bool {
	return cimgui.InputInt(label, v, 1, 100, 0)
}

// InputIntV draws an int input with step buttons and flags.
func InputIntV(label string, v *int32, step, stepFast int32, flags InputTextFlags) bool {
	return cimgui.InputInt(label, v, step, stepFast, int(flags))
}

// DragFloat draws a draggable float bound to v and reports whether it changed.
func DragFloat(label string, v *float32) bool {
	return cimgui.DragFloat(label, v, 1, 0, 0, "%.3f", 0)
}

// DragFloatV draws a draggable float with a speed, range, format, and flags. A
// zero min and max leaves the value unclamped.
func DragFloatV(label string, v *float32, speed, min, max float32, format string, flags SliderFlags) bool {
	return cimgui.DragFloat(label, v, speed, min, max, format, int(flags))
}

// DragInt draws a draggable int bound to v and reports whether it changed.
func DragInt(label string, v *int32) bool {
	return cimgui.DragInt(label, v, 1, 0, 0, "%d", 0)
}

// DragIntV draws a draggable int with a speed, range, format, and flags. A zero
// min and max leaves the value unclamped.
func DragIntV(label string, v *int32, speed float32, min, max int32, format string, flags SliderFlags) bool {
	return cimgui.DragInt(label, v, speed, min, max, format, int(flags))
}

// ColorEdit3 edits an RGB color in place and reports whether it changed.
func ColorEdit3(label string, col *[3]float32) bool {
	return cimgui.ColorEdit3(label, col, 0)
}

// ColorEdit3V edits an RGB color with flags.
func ColorEdit3V(label string, col *[3]float32, flags ColorEditFlags) bool {
	return cimgui.ColorEdit3(label, col, int(flags))
}

// ColorEdit4 edits an RGBA color in place and reports whether it changed.
func ColorEdit4(label string, col *[4]float32) bool {
	return cimgui.ColorEdit4(label, col, 0)
}

// ColorEdit4V edits an RGBA color with flags.
func ColorEdit4V(label string, col *[4]float32, flags ColorEditFlags) bool {
	return cimgui.ColorEdit4(label, col, int(flags))
}

// TreeNode opens a tree node. Call [TreePop] only if it returns true.
func TreeNode(label string) bool { return cimgui.TreeNode(label) }

// TreePop closes the tree node opened by [TreeNode].
func TreePop() { cimgui.TreePop() }

// CollapsingHeader draws a collapsing header and reports whether it is open.
func CollapsingHeader(label string) bool { return cimgui.CollapsingHeader(label, 0) }

// CollapsingHeaderV draws a collapsing header with flags.
func CollapsingHeaderV(label string, flags TreeNodeFlags) bool {
	return cimgui.CollapsingHeader(label, int(flags))
}

// BeginCombo opens a combo box. Call [EndCombo] only if it returns true.
func BeginCombo(label, preview string) bool { return cimgui.BeginCombo(label, preview, 0) }

// BeginComboV opens a combo box with flags.
func BeginComboV(label, preview string, flags ComboFlags) bool {
	return cimgui.BeginCombo(label, preview, int(flags))
}

// EndCombo closes the combo box opened by [BeginCombo].
func EndCombo() { cimgui.EndCombo() }

// Selectable draws a selectable item and reports whether it was clicked.
func Selectable(label string, selected bool) bool {
	return cimgui.Selectable(label, selected, 0, 0, 0)
}

// SelectableV draws a selectable item with flags and an explicit size.
func SelectableV(label string, selected bool, flags SelectableFlags, size Vec2) bool {
	return cimgui.Selectable(label, selected, int(flags), size.X, size.Y)
}

// BeginTooltip opens a tooltip. Call [EndTooltip] only if it returns true.
func BeginTooltip() bool { return cimgui.BeginTooltip() }

// EndTooltip closes the tooltip opened by [BeginTooltip].
func EndTooltip() { cimgui.EndTooltip() }
