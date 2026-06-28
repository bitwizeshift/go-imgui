package input

import (
	"github.com/bitwizeshift/go-imgui"
	"github.com/bitwizeshift/go-imgui/internal/cimgui"
)

// Double edits a bound float64 in a box with optional +/- step buttons.
type Double struct {
	Label    string
	Value    *float64
	Step     float64
	StepFast float64
	Format   string // default "%.6f"
	textFlags
	OnChange func(float64)
	changed  bool
	scratch  float64
}

// NewDouble returns a double input bound to value.
func NewDouble(label string, value *float64) *Double {
	return &Double{Label: label, Value: value}
}

// Display draws the input.
func (d *Double) Display() {
	v := d.Value
	if v == nil {
		v = &d.scratch
	}
	format := d.Format
	if format == "" {
		format = "%.6f"
	}
	d.changed = cimgui.InputDouble(d.Label, v, d.Step, d.StepFast, format, d.flags)
	if d.changed && d.OnChange != nil {
		d.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (d *Double) Changed() bool { return d.changed }

var _ imgui.Widget = (*Double)(nil)
