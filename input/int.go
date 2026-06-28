package input

import (
	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// Int edits a bound int in a box with optional +/- step buttons.
type Int struct {
	Label    string
	Value    *int32
	Step     int32 // default 1
	StepFast int32 // default 100
	textFlags
	OnChange func(int32)
	changed  bool
	scratch  int32
}

// NewInt returns an integer input bound to value.
func NewInt(label string, value *int32) *Int {
	if value == nil {
		value = new(int32)
	}
	return &Int{
		Label:    label,
		Value:    value,
		Step:     1,
		StepFast: 100,
	}
}

// Display draws the input.
func (i *Int) Display() {
	v := i.Value
	if v == nil {
		v = &i.scratch
	}
	step, stepFast := i.Step, i.StepFast
	if step == 0 {
		step = 1
	}
	if stepFast == 0 {
		stepFast = 100
	}
	i.changed = cimgui.InputInt(i.Label, v, step, stepFast, i.flags)
	if i.changed && i.OnChange != nil {
		i.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (i *Int) Changed() bool {
	return i.changed
}

var _ imgui.Widget = (*Int)(nil)

//------------------------------------------------------------------------------
