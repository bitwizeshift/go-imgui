package input

import (
	"github.com/bitwizeshift/go-imgui"
	"github.com/bitwizeshift/go-imgui/internal/cimgui"
)

// Float2 edits a bound 2-component float vector in a box.
type Float2 struct {
	Label  string
	Value  *[2]float32
	Format string // default "%.3f"
	textFlags
	OnChange func([2]float32)
	changed  bool
	scratch  [2]float32
}

// NewFloat2 returns a 2-component float input bound to value.
func NewFloat2(label string, value *[2]float32) *Float2 {
	return &Float2{Label: label, Value: value}
}

// Display draws the input.
func (f *Float2) Display() {
	v := f.Value
	if v == nil {
		v = &f.scratch
	}
	format := f.Format
	if format == "" {
		format = "%.3f"
	}
	f.changed = cimgui.InputFloat2(f.Label, v, format, f.flags)
	if f.changed && f.OnChange != nil {
		f.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (f *Float2) Changed() bool { return f.changed }

var _ imgui.Widget = (*Float2)(nil)

// Float3 edits a bound 3-component float vector in a box.
type Float3 struct {
	Label  string
	Value  *[3]float32
	Format string // default "%.3f"
	textFlags
	OnChange func([3]float32)
	changed  bool
	scratch  [3]float32
}

// NewFloat3 returns a 3-component float input bound to value.
func NewFloat3(label string, value *[3]float32) *Float3 {
	return &Float3{Label: label, Value: value}
}

// Display draws the input.
func (f *Float3) Display() {
	v := f.Value
	if v == nil {
		v = &f.scratch
	}
	format := f.Format
	if format == "" {
		format = "%.3f"
	}
	f.changed = cimgui.InputFloat3(f.Label, v, format, f.flags)
	if f.changed && f.OnChange != nil {
		f.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (f *Float3) Changed() bool { return f.changed }

var _ imgui.Widget = (*Float3)(nil)

// Float4 edits a bound 4-component float vector in a box.
type Float4 struct {
	Label  string
	Value  *[4]float32
	Format string // default "%.3f"
	textFlags
	OnChange func([4]float32)
	changed  bool
	scratch  [4]float32
}

// NewFloat4 returns a 4-component float input bound to value.
func NewFloat4(label string, value *[4]float32) *Float4 {
	return &Float4{Label: label, Value: value}
}

// Display draws the input.
func (f *Float4) Display() {
	v := f.Value
	if v == nil {
		v = &f.scratch
	}
	format := f.Format
	if format == "" {
		format = "%.3f"
	}
	f.changed = cimgui.InputFloat4(f.Label, v, format, f.flags)
	if f.changed && f.OnChange != nil {
		f.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (f *Float4) Changed() bool { return f.changed }

var _ imgui.Widget = (*Float4)(nil)

// Int2 edits a bound 2-component int vector in a box.
type Int2 struct {
	Label string
	Value *[2]int32
	textFlags
	OnChange func([2]int32)
	changed  bool
	scratch  [2]int32
}

// NewInt2 returns a 2-component int input bound to value.
func NewInt2(label string, value *[2]int32) *Int2 {
	return &Int2{Label: label, Value: value}
}

// Display draws the input.
func (i *Int2) Display() {
	v := i.Value
	if v == nil {
		v = &i.scratch
	}
	i.changed = cimgui.InputInt2(i.Label, v, i.flags)
	if i.changed && i.OnChange != nil {
		i.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (i *Int2) Changed() bool { return i.changed }

var _ imgui.Widget = (*Int2)(nil)

// Int3 edits a bound 3-component int vector in a box.
type Int3 struct {
	Label string
	Value *[3]int32
	textFlags
	OnChange func([3]int32)
	changed  bool
	scratch  [3]int32
}

// NewInt3 returns a 3-component int input bound to value.
func NewInt3(label string, value *[3]int32) *Int3 {
	return &Int3{Label: label, Value: value}
}

// Display draws the input.
func (i *Int3) Display() {
	v := i.Value
	if v == nil {
		v = &i.scratch
	}
	i.changed = cimgui.InputInt3(i.Label, v, i.flags)
	if i.changed && i.OnChange != nil {
		i.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (i *Int3) Changed() bool { return i.changed }

var _ imgui.Widget = (*Int3)(nil)

// Int4 edits a bound 4-component int vector in a box.
type Int4 struct {
	Label string
	Value *[4]int32
	textFlags
	OnChange func([4]int32)
	changed  bool
	scratch  [4]int32
}

// NewInt4 returns a 4-component int input bound to value.
func NewInt4(label string, value *[4]int32) *Int4 {
	return &Int4{Label: label, Value: value}
}

// Display draws the input.
func (i *Int4) Display() {
	v := i.Value
	if v == nil {
		v = &i.scratch
	}
	i.changed = cimgui.InputInt4(i.Label, v, i.flags)
	if i.changed && i.OnChange != nil {
		i.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (i *Int4) Changed() bool { return i.changed }

var _ imgui.Widget = (*Int4)(nil)
