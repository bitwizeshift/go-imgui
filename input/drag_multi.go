package input

import (
	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// dragSpeed returns speed, or 1 when speed is zero.
func dragSpeed(speed float32) float32 {
	if speed == 0 {
		return 1
	}
	return speed
}

// DragFloat2 edits a bound 2-component float vector by dragging.
type DragFloat2 struct {
	Label    string
	Value    *[2]float32
	Speed    float32 // default 1
	Min, Max float32
	Format   string // default "%.3f"
	sliderFlags
	OnChange func([2]float32)
	changed  bool
	scratch  [2]float32
}

// NewDragFloat2 returns a draggable 2-component float bound to value.
func NewDragFloat2(label string, value *[2]float32) *DragFloat2 {
	return &DragFloat2{Label: label, Value: value}
}

// Display draws the drag.
func (d *DragFloat2) Display() {
	v := d.Value
	if v == nil {
		v = &d.scratch
	}
	d.changed = cimgui.DragFloat2(d.Label, v, dragSpeed(d.Speed), d.Min, d.Max, sliderFormat(d.Format, "%.3f"), d.flags)
	if d.changed && d.OnChange != nil {
		d.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (d *DragFloat2) Changed() bool { return d.changed }

var _ imgui.Widget = (*DragFloat2)(nil)

// DragFloat3 edits a bound 3-component float vector by dragging.
type DragFloat3 struct {
	Label    string
	Value    *[3]float32
	Speed    float32 // default 1
	Min, Max float32
	Format   string // default "%.3f"
	sliderFlags
	OnChange func([3]float32)
	changed  bool
	scratch  [3]float32
}

// NewDragFloat3 returns a draggable 3-component float bound to value.
func NewDragFloat3(label string, value *[3]float32) *DragFloat3 {
	return &DragFloat3{Label: label, Value: value}
}

// Display draws the drag.
func (d *DragFloat3) Display() {
	v := d.Value
	if v == nil {
		v = &d.scratch
	}
	d.changed = cimgui.DragFloat3(d.Label, v, dragSpeed(d.Speed), d.Min, d.Max, sliderFormat(d.Format, "%.3f"), d.flags)
	if d.changed && d.OnChange != nil {
		d.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (d *DragFloat3) Changed() bool { return d.changed }

var _ imgui.Widget = (*DragFloat3)(nil)

// DragFloat4 edits a bound 4-component float vector by dragging.
type DragFloat4 struct {
	Label    string
	Value    *[4]float32
	Speed    float32 // default 1
	Min, Max float32
	Format   string // default "%.3f"
	sliderFlags
	OnChange func([4]float32)
	changed  bool
	scratch  [4]float32
}

// NewDragFloat4 returns a draggable 4-component float bound to value.
func NewDragFloat4(label string, value *[4]float32) *DragFloat4 {
	return &DragFloat4{Label: label, Value: value}
}

// Display draws the drag.
func (d *DragFloat4) Display() {
	v := d.Value
	if v == nil {
		v = &d.scratch
	}
	d.changed = cimgui.DragFloat4(d.Label, v, dragSpeed(d.Speed), d.Min, d.Max, sliderFormat(d.Format, "%.3f"), d.flags)
	if d.changed && d.OnChange != nil {
		d.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (d *DragFloat4) Changed() bool { return d.changed }

var _ imgui.Widget = (*DragFloat4)(nil)

// DragInt2 edits a bound 2-component int vector by dragging.
type DragInt2 struct {
	Label    string
	Value    *[2]int32
	Speed    float32 // default 1
	Min, Max int32
	Format   string // default "%d"
	sliderFlags
	OnChange func([2]int32)
	changed  bool
	scratch  [2]int32
}

// NewDragInt2 returns a draggable 2-component int bound to value.
func NewDragInt2(label string, value *[2]int32) *DragInt2 {
	return &DragInt2{Label: label, Value: value}
}

// Display draws the drag.
func (d *DragInt2) Display() {
	v := d.Value
	if v == nil {
		v = &d.scratch
	}
	d.changed = cimgui.DragInt2(d.Label, v, dragSpeed(d.Speed), d.Min, d.Max, sliderFormat(d.Format, "%d"), d.flags)
	if d.changed && d.OnChange != nil {
		d.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (d *DragInt2) Changed() bool { return d.changed }

var _ imgui.Widget = (*DragInt2)(nil)

// DragInt3 edits a bound 3-component int vector by dragging.
type DragInt3 struct {
	Label    string
	Value    *[3]int32
	Speed    float32 // default 1
	Min, Max int32
	Format   string // default "%d"
	sliderFlags
	OnChange func([3]int32)
	changed  bool
	scratch  [3]int32
}

// NewDragInt3 returns a draggable 3-component int bound to value.
func NewDragInt3(label string, value *[3]int32) *DragInt3 {
	return &DragInt3{Label: label, Value: value}
}

// Display draws the drag.
func (d *DragInt3) Display() {
	v := d.Value
	if v == nil {
		v = &d.scratch
	}
	d.changed = cimgui.DragInt3(d.Label, v, dragSpeed(d.Speed), d.Min, d.Max, sliderFormat(d.Format, "%d"), d.flags)
	if d.changed && d.OnChange != nil {
		d.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (d *DragInt3) Changed() bool { return d.changed }

var _ imgui.Widget = (*DragInt3)(nil)

// DragInt4 edits a bound 4-component int vector by dragging.
type DragInt4 struct {
	Label    string
	Value    *[4]int32
	Speed    float32 // default 1
	Min, Max int32
	Format   string // default "%d"
	sliderFlags
	OnChange func([4]int32)
	changed  bool
	scratch  [4]int32
}

// NewDragInt4 returns a draggable 4-component int bound to value.
func NewDragInt4(label string, value *[4]int32) *DragInt4 {
	return &DragInt4{Label: label, Value: value}
}

// Display draws the drag.
func (d *DragInt4) Display() {
	v := d.Value
	if v == nil {
		v = &d.scratch
	}
	d.changed = cimgui.DragInt4(d.Label, v, dragSpeed(d.Speed), d.Min, d.Max, sliderFormat(d.Format, "%d"), d.flags)
	if d.changed && d.OnChange != nil {
		d.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (d *DragInt4) Changed() bool { return d.changed }

var _ imgui.Widget = (*DragInt4)(nil)

// DragFloatRange edits a [Min,Max] float range by dragging two handles, keeping
// CurrentMin <= CurrentMax.
type DragFloatRange struct {
	Label             string
	CurrentMin        *float32
	CurrentMax        *float32
	Speed             float32 // default 1
	Min, Max          float32
	Format, FormatMax string // default "%.3f"; FormatMax falls back to Format
	sliderFlags
	OnChange               func(min, max float32)
	changed                bool
	scratchMin, scratchMax float32
}

// NewDragFloatRange returns a draggable float range bound to currentMin/currentMax.
func NewDragFloatRange(label string, currentMin, currentMax *float32) *DragFloatRange {
	return &DragFloatRange{Label: label, CurrentMin: currentMin, CurrentMax: currentMax}
}

// Display draws the range drag.
func (d *DragFloatRange) Display() {
	lo, hi := d.CurrentMin, d.CurrentMax
	if lo == nil {
		lo = &d.scratchMin
	}
	if hi == nil {
		hi = &d.scratchMax
	}
	format := sliderFormat(d.Format, "%.3f")
	formatMax := sliderFormat(d.FormatMax, format)
	d.changed = cimgui.DragFloatRange2(d.Label, lo, hi, dragSpeed(d.Speed), d.Min, d.Max, format, formatMax, d.flags)
	if d.changed && d.OnChange != nil {
		d.OnChange(*lo, *hi)
	}
}

// Changed reports whether the range changed during the last Display.
func (d *DragFloatRange) Changed() bool { return d.changed }

var _ imgui.Widget = (*DragFloatRange)(nil)

// DragIntRange edits a [Min,Max] int range by dragging two handles, keeping
// CurrentMin <= CurrentMax.
type DragIntRange struct {
	Label             string
	CurrentMin        *int32
	CurrentMax        *int32
	Speed             float32 // default 1
	Min, Max          int32
	Format, FormatMax string // default "%d"; FormatMax falls back to Format
	sliderFlags
	OnChange               func(min, max int32)
	changed                bool
	scratchMin, scratchMax int32
}

// NewDragIntRange returns a draggable int range bound to currentMin/currentMax.
func NewDragIntRange(label string, currentMin, currentMax *int32) *DragIntRange {
	return &DragIntRange{Label: label, CurrentMin: currentMin, CurrentMax: currentMax}
}

// Display draws the range drag.
func (d *DragIntRange) Display() {
	lo, hi := d.CurrentMin, d.CurrentMax
	if lo == nil {
		lo = &d.scratchMin
	}
	if hi == nil {
		hi = &d.scratchMax
	}
	format := sliderFormat(d.Format, "%d")
	formatMax := sliderFormat(d.FormatMax, format)
	d.changed = cimgui.DragIntRange2(d.Label, lo, hi, dragSpeed(d.Speed), d.Min, d.Max, format, formatMax, d.flags)
	if d.changed && d.OnChange != nil {
		d.OnChange(*lo, *hi)
	}
}

// Changed reports whether the range changed during the last Display.
func (d *DragIntRange) Changed() bool { return d.changed }

var _ imgui.Widget = (*DragIntRange)(nil)
