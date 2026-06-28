package input

import (
	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// CheckboxFlags is a checkbox that toggles FlagsValue within the bound int
// bitfield. It appears checked when all of FlagsValue's bits are set, and
// indeterminate when only some are.
type CheckboxFlags struct {
	Label      string
	Value      *int32
	FlagsValue int32
	OnChange   func(int32)
	changed    bool
	scratch    int32
}

// NewCheckboxFlags returns a flag checkbox toggling flagsValue within value.
func NewCheckboxFlags(label string, value *int32, flagsValue int32) *CheckboxFlags {
	return &CheckboxFlags{Label: label, Value: value, FlagsValue: flagsValue}
}

// Display draws the checkbox.
func (c *CheckboxFlags) Display() {
	v := c.Value
	if v == nil {
		v = &c.scratch
	}
	c.changed = cimgui.CheckboxFlags_IntPtr(c.Label, v, c.FlagsValue)
	if c.changed && c.OnChange != nil {
		c.OnChange(*v)
	}
}

// Changed reports whether the bitfield changed during the last Display.
func (c *CheckboxFlags) Changed() bool { return c.changed }

var _ imgui.Widget = (*CheckboxFlags)(nil)

// CheckboxFlagsUint is [CheckboxFlags] over an unsigned bitfield.
type CheckboxFlagsUint struct {
	Label      string
	Value      *uint32
	FlagsValue uint32
	OnChange   func(uint32)
	changed    bool
	scratch    uint32
}

// NewCheckboxFlagsUint returns a flag checkbox toggling flagsValue within value.
func NewCheckboxFlagsUint(label string, value *uint32, flagsValue uint32) *CheckboxFlagsUint {
	return &CheckboxFlagsUint{Label: label, Value: value, FlagsValue: flagsValue}
}

// Display draws the checkbox.
func (c *CheckboxFlagsUint) Display() {
	v := c.Value
	if v == nil {
		v = &c.scratch
	}
	c.changed = cimgui.CheckboxFlags_UintPtr(c.Label, v, c.FlagsValue)
	if c.changed && c.OnChange != nil {
		c.OnChange(*v)
	}
}

// Changed reports whether the bitfield changed during the last Display.
func (c *CheckboxFlagsUint) Changed() bool { return c.changed }

var _ imgui.Widget = (*CheckboxFlagsUint)(nil)
