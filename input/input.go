// Package input provides value-editing widgets: checkboxes (including flag
// checkboxes), radio buttons, sliders and drags (scalar, multi-component and
// vertical), numeric and vector inputs, text inputs, and generic widgets over any
// numeric type (see [Scalar]). Each binds a pointer to the edited value and
// reports edits via OnChange and a Changed poll.
package input

import (
	"github.com/bitwizeshift/go-imgui/internal/cimgui"
)

// sliderFlags is the shared cached slider/drag bitfield plus its setters,
// embedded by SliderFloat, SliderInt, DragFloat and DragInt. Setters update the
// cached field when called so Display never recomputes it.
type sliderFlags struct {
	flags cimgui.SliderFlags
}

func setSliderFlag(f *cimgui.SliderFlags, bit cimgui.SliderFlags, on bool) {
	if on {
		*f |= bit
	} else {
		*f &^= bit
	}
}

// SetLogarithmic makes the slider/drag operate on a logarithmic scale.
func (f *sliderFlags) SetLogarithmic(on bool) {
	setSliderFlag(&f.flags, cimgui.SliderFlagsLogarithmic, on)
}

// SetInputAllowed allows or prevents Ctrl+Click text entry of the value
// (default allowed).
func (f *sliderFlags) SetInputAllowed(on bool) {
	setSliderFlag(&f.flags, cimgui.SliderFlagsNoInput, !on)
}

// SetAlwaysClamp clamps the value to [Min, Max] even when typed in.
func (f *sliderFlags) SetAlwaysClamp(on bool) {
	setSliderFlag(&f.flags, cimgui.SliderFlagsAlwaysClamp, on)
}

// textFlags is the shared cached input-text bitfield plus its setters, embedded
// by Int, Float and Text. Setters update the cached field when called so Display
// never recomputes it.
type textFlags struct {
	flags cimgui.InputTextFlags
}

func setTextFlag(f *cimgui.InputTextFlags, bit cimgui.InputTextFlags, on bool) {
	if on {
		*f |= bit
	} else {
		*f &^= bit
	}
}

// SetPassword obscures the typed characters.
func (f *textFlags) SetPassword(on bool) {
	setTextFlag(&f.flags, cimgui.InputTextFlagsPassword, on)
}

// SetReadOnly makes the field read-only.
func (f *textFlags) SetReadOnly(on bool) {
	setTextFlag(&f.flags, cimgui.InputTextFlagsReadOnly, on)
}

// SetDecimalOnly restricts input to decimal characters (0-9, +, -, ., *, /).
func (f *textFlags) SetDecimalOnly(on bool) {
	setTextFlag(&f.flags, cimgui.InputTextFlagsCharsDecimal, on)
}

// SetUppercase converts typed letters to uppercase.
func (f *textFlags) SetUppercase(on bool) {
	setTextFlag(&f.flags, cimgui.InputTextFlagsCharsUppercase, on)
}

// SetEnterReturns makes Changed report true only when Enter is pressed.
func (f *textFlags) SetEnterReturns(on bool) {
	setTextFlag(&f.flags, cimgui.InputTextFlagsEnterReturnsTrue, on)
}

// Checkbox toggles a bound bool.
type Checkbox struct {
	Label    string
	Value    *bool
	OnChange func(bool)
	changed  bool
	scratch  bool
}

// NewCheckbox returns a checkbox bound to value.
func NewCheckbox(label string, value *bool) *Checkbox {
	return &Checkbox{Label: label, Value: value}
}

// Display draws the checkbox.
func (c *Checkbox) Display() {
	v := c.Value
	if v == nil {
		v = &c.scratch
	}
	c.changed = cimgui.Checkbox(c.Label, v)
	if c.changed && c.OnChange != nil {
		c.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (c *Checkbox) Changed() bool { return c.changed }

// Radio is one option of a radio group: it sets the bound int to Choice when
// selected. Give every option in a group the same Value pointer.
type Radio struct {
	Label    string
	Value    *int32
	Choice   int32
	OnChange func(int32)
	changed  bool
	scratch  int32
}

// NewRadio returns a radio button that sets value to choice when selected.
func NewRadio(label string, value *int32, choice int32) *Radio {
	return &Radio{Label: label, Value: value, Choice: choice}
}

// Display draws the radio button.
func (r *Radio) Display() {
	v := r.Value
	if v == nil {
		v = &r.scratch
	}
	r.changed = cimgui.RadioButton_IntPtr(r.Label, v, r.Choice)
	if r.changed && r.OnChange != nil {
		r.OnChange(*v)
	}
}

// Changed reports whether the selection changed during the last Display.
func (r *Radio) Changed() bool { return r.changed }

// SliderFloat edits a bound float by dragging within [Min, Max].
type SliderFloat struct {
	Label    string
	Value    *float32
	Min, Max float32
	Format   string // default "%.3f"
	sliderFlags
	OnChange func(float32)
	changed  bool
	scratch  float32
}

// NewSliderFloat returns a float slider bound to value.
func NewSliderFloat(label string, value *float32, min, max float32) *SliderFloat {
	return &SliderFloat{Label: label, Value: value, Min: min, Max: max}
}

// Display draws the slider.
func (s *SliderFloat) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	format := s.Format
	if format == "" {
		format = "%.3f"
	}
	s.changed = cimgui.SliderFloat(s.Label, v, s.Min, s.Max, format, s.flags)
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *SliderFloat) Changed() bool { return s.changed }

// SliderInt edits a bound int by dragging within [Min, Max].
type SliderInt struct {
	Label    string
	Value    *int32
	Min, Max int32
	Format   string // default "%d"
	sliderFlags
	OnChange func(int32)
	changed  bool
	scratch  int32
}

// NewSliderInt returns an int slider bound to value.
func NewSliderInt(label string, value *int32, min, max int32) *SliderInt {
	return &SliderInt{Label: label, Value: value, Min: min, Max: max}
}

// Display draws the slider.
func (s *SliderInt) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	format := s.Format
	if format == "" {
		format = "%d"
	}
	s.changed = cimgui.SliderInt(s.Label, v, s.Min, s.Max, format, s.flags)
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *SliderInt) Changed() bool { return s.changed }

// DragFloat edits a bound float by dragging. Min==Max==0 leaves it unbounded.
type DragFloat struct {
	Label    string
	Value    *float32
	Speed    float32 // default 1
	Min, Max float32
	Format   string // default "%.3f"
	sliderFlags
	OnChange func(float32)
	changed  bool
	scratch  float32
}

// NewDragFloat returns a draggable float bound to value.
func NewDragFloat(label string, value *float32) *DragFloat {
	return &DragFloat{Label: label, Value: value}
}

// Display draws the drag.
func (d *DragFloat) Display() {
	v := d.Value
	if v == nil {
		v = &d.scratch
	}
	speed := d.Speed
	if speed == 0 {
		speed = 1
	}
	format := d.Format
	if format == "" {
		format = "%.3f"
	}
	d.changed = cimgui.DragFloat(d.Label, v, speed, d.Min, d.Max, format, d.flags)
	if d.changed && d.OnChange != nil {
		d.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (d *DragFloat) Changed() bool { return d.changed }

// DragInt edits a bound int by dragging. Min==Max==0 leaves it unbounded.
type DragInt struct {
	Label    string
	Value    *int32
	Speed    float32 // default 1
	Min, Max int32
	Format   string // default "%d"
	sliderFlags
	OnChange func(int32)
	changed  bool
	scratch  int32
}

// NewDragInt returns a draggable int bound to value.
func NewDragInt(label string, value *int32) *DragInt { return &DragInt{Label: label, Value: value} }

// Display draws the drag.
func (d *DragInt) Display() {
	v := d.Value
	if v == nil {
		v = &d.scratch
	}
	speed := d.Speed
	if speed == 0 {
		speed = 1
	}
	format := d.Format
	if format == "" {
		format = "%d"
	}
	d.changed = cimgui.DragInt(d.Label, v, speed, d.Min, d.Max, format, d.flags)
	if d.changed && d.OnChange != nil {
		d.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (d *DragInt) Changed() bool { return d.changed }

// Float edits a bound float in a box with optional +/- step buttons.
type Float struct {
	Label    string
	Value    *float32
	Step     float32
	StepFast float32
	Format   string // default "%.3f"
	textFlags
	OnChange func(float32)
	changed  bool
	scratch  float32
}

// NewFloat returns a float input bound to value.
func NewFloat(label string, value *float32) *Float { return &Float{Label: label, Value: value} }

// Display draws the input.
func (f *Float) Display() {
	v := f.Value
	if v == nil {
		v = &f.scratch
	}
	format := f.Format
	if format == "" {
		format = "%.3f"
	}
	f.changed = cimgui.InputFloat(f.Label, v, f.Step, f.StepFast, format, f.flags)
	if f.changed && f.OnChange != nil {
		f.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (f *Float) Changed() bool { return f.changed }
