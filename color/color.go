// Package color provides color-editing widgets bound to an [imgui.Color].
package color

import (
	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// flagSetters is the shared cached color-edit bitfield plus its setters,
// embedded by both Edit and Button. Setters update the cached field when called
// so Display never recomputes it.
type flagSetters struct {
	flags cimgui.ColorEditFlags
}

func setColorFlag(f *cimgui.ColorEditFlags, bit cimgui.ColorEditFlags, on bool) {
	if on {
		*f |= bit
	} else {
		*f &^= bit
	}
}

// SetAlpha includes the alpha channel in the editor/inputs (default included).
func (f *flagSetters) SetAlpha(on bool) {
	setColorFlag(&f.flags, cimgui.ColorEditFlagsNoAlpha, !on)
}

// SetAlphaBar shows a vertical alpha bar alongside the picker.
func (f *flagSetters) SetAlphaBar(on bool) {
	setColorFlag(&f.flags, cimgui.ColorEditFlagsAlphaBar, on)
}

// SetFloat displays component values as 0..1 floats instead of 0..255 integers.
func (f *flagSetters) SetFloat(on bool) {
	setColorFlag(&f.flags, cimgui.ColorEditFlagsFloat, on)
}

// SetDisplayHSV displays the value as HSV instead of RGB.
func (f *flagSetters) SetDisplayHSV(on bool) {
	setColorFlag(&f.flags, cimgui.ColorEditFlagsDisplayHSV, on)
}

// SetInputs shows or hides the numeric input fields (default shown).
func (f *flagSetters) SetInputs(on bool) {
	setColorFlag(&f.flags, cimgui.ColorEditFlagsNoInputs, !on)
}

// Edit edits a bound color. By default it is a compact RGB editor; set Alpha to
// include the alpha channel and Picker to show the full picker.
type Edit struct {
	Label    string
	Value    *imgui.Color
	Alpha    bool
	Picker   bool
	OnChange func(imgui.Color)
	changed  bool
	scratch  imgui.Color
	flagSetters
}

// NewEdit returns a compact RGB color editor bound to value.
func NewEdit(label string, value *imgui.Color) *Edit { return &Edit{Label: label, Value: value} }

// NewEditAlpha returns a compact RGBA color editor bound to value.
func NewEditAlpha(label string, value *imgui.Color) *Edit {
	return &Edit{Label: label, Value: value, Alpha: true}
}

// NewPicker returns an RGBA color picker bound to value.
func NewPicker(label string, value *imgui.Color) *Edit {
	return &Edit{Label: label, Value: value, Alpha: true, Picker: true}
}

// Display draws the editor or picker.
func (e *Edit) Display() {
	c := e.Value
	if c == nil {
		c = &e.scratch
	}
	if e.Alpha {
		arr := [4]float32{c.R, c.G, c.B, c.A}
		if e.Picker {
			e.changed = cimgui.ColorPicker4(e.Label, &arr, e.flags, nil)
		} else {
			e.changed = cimgui.ColorEdit4(e.Label, &arr, e.flags)
		}
		c.R, c.G, c.B, c.A = arr[0], arr[1], arr[2], arr[3]
	} else {
		arr := [3]float32{c.R, c.G, c.B}
		if e.Picker {
			e.changed = cimgui.ColorPicker3(e.Label, &arr, e.flags)
		} else {
			e.changed = cimgui.ColorEdit3(e.Label, &arr, e.flags)
		}
		c.R, c.G, c.B = arr[0], arr[1], arr[2]
	}
	if e.changed && e.OnChange != nil {
		e.OnChange(*c)
	}
}

// Changed reports whether the color changed during the last Display.
func (e *Edit) Changed() bool { return e.changed }

// Button is a clickable color swatch. ID must be unique.
type Button struct {
	ID      string
	Color   imgui.Color
	Size    imgui.Vec2
	OnClick func()
	pressed bool
	flagSetters
}

// NewButton returns a color-swatch button.
func NewButton(id string, c imgui.Color) *Button { return &Button{ID: id, Color: c} }

// Display draws the swatch.
func (b *Button) Display() {
	b.pressed = cimgui.ColorButton(b.ID, b.Color.Vec4(), b.flags, b.Size)
	if b.pressed && b.OnClick != nil {
		b.OnClick()
	}
}

// Pressed reports whether the swatch was clicked during the last Display.
func (b *Button) Pressed() bool { return b.pressed }
