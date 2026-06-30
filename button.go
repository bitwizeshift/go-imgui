package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// ButtonOptions are the optional inputs to [InvisibleButton]. A nil
// *ButtonOptions uses Dear ImGui's defaults; each field maps to an
// ImGuiButtonFlags_ bit.
type ButtonOptions struct {
	MouseButtonLeft   bool // ImGuiButtonFlags_MouseButtonLeft
	MouseButtonRight  bool // ImGuiButtonFlags_MouseButtonRight
	MouseButtonMiddle bool // ImGuiButtonFlags_MouseButtonMiddle
	EnableNav         bool // ImGuiButtonFlags_EnableNav
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *ButtonOptions) flags() cimgui.ButtonFlags {
	if o == nil {
		return cimgui.ButtonFlagsNone
	}
	var f cimgui.ButtonFlags
	if o.MouseButtonLeft {
		f |= cimgui.ButtonFlagsMouseButtonLeft
	}
	if o.MouseButtonRight {
		f |= cimgui.ButtonFlagsMouseButtonRight
	}
	if o.MouseButtonMiddle {
		f |= cimgui.ButtonFlagsMouseButtonMiddle
	}
	if o.EnableNav {
		f |= cimgui.ButtonFlagsEnableNav
	}
	return f
}

// Button draws a button and reports whether it was clicked this frame. A zero
// size auto-fits the label. It models ImGui::Button.
func Button(label string, size Vec2) bool {
	return cimgui.Button(label, size)
}

// SmallButton draws a button with no frame padding and reports whether it was
// clicked. It models ImGui::SmallButton.
func SmallButton(label string) bool {
	return cimgui.SmallButton(label)
}

// InvisibleButton draws a sizeable behaviour-only button with no visuals and
// reports whether it was clicked. It models ImGui::InvisibleButton.
func InvisibleButton(strID string, size Vec2, opts *ButtonOptions) bool {
	return cimgui.InvisibleButton(strID, size, opts.flags())
}

// ArrowButton draws a square button containing an arrow in dir and reports
// whether it was clicked. It models ImGui::ArrowButton.
func ArrowButton(strID string, dir Dir) bool {
	return cimgui.ArrowButton(strID, dir)
}

// Bullet draws a small bullet and keeps the cursor on the same line. It models
// ImGui::Bullet.
func Bullet() {
	cimgui.Bullet()
}

// Checkbox draws a checkbox bound to v and reports whether it changed this
// frame. It models ImGui::Checkbox.
func Checkbox(label string, v *bool) bool {
	return cimgui.Checkbox(label, v)
}

// CheckboxFlags draws a checkbox that toggles flagsValue within the bitset
// flags, reporting whether it changed. flags may point to an int32 or uint32,
// modelling ImGui::CheckboxFlags.
func CheckboxFlags[T SignedOrUnsigned32](label string, flags *T, flagsValue T) bool {
	switch p := any(flags).(type) {
	case *int32:
		return cimgui.CheckboxFlags_IntPtr(label, p, any(flagsValue).(int32))
	case *uint32:
		return cimgui.CheckboxFlags_UintPtr(label, p, any(flagsValue).(uint32))
	}
	return false
}

// RadioButton draws a radio button rendered active and reports whether it was
// clicked. It models ImGui::RadioButton (the bool overload).
func RadioButton(label string, active bool) bool {
	return cimgui.RadioButton_Bool(label, active)
}

// RadioButtonGroup draws a radio button that sets *v to vButton when clicked and
// renders active while *v equals vButton, reporting whether it changed. It
// models ImGui::RadioButton (the int* overload).
func RadioButtonGroup(label string, v *int32, vButton int32) bool {
	return cimgui.RadioButton_IntPtr(label, v, vButton)
}

// ProgressBar draws a progress bar filled to fraction (0..1). A zero size
// auto-fits; a non-empty overlay is drawn centered over the bar. It models
// ImGui::ProgressBar.
func ProgressBar(fraction float32, size Vec2, overlay string) {
	cimgui.ProgressBar(fraction, size, overlay)
}
