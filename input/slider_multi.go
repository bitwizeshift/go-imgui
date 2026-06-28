package input

import (
	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// SliderFloat2 edits a bound 2-component float vector by dragging within [Min, Max].
type SliderFloat2 struct {
	Label    string
	Value    *[2]float32
	Min, Max float32
	Format   string // default "%.3f"
	sliderFlags
	OnChange func([2]float32)
	changed  bool
	scratch  [2]float32
}

// NewSliderFloat2 returns a 2-component float slider bound to value.
func NewSliderFloat2(label string, value *[2]float32, min, max float32) *SliderFloat2 {
	return &SliderFloat2{Label: label, Value: value, Min: min, Max: max}
}

// Display draws the slider.
func (s *SliderFloat2) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	s.changed = cimgui.SliderFloat2(s.Label, v, s.Min, s.Max, sliderFormat(s.Format, "%.3f"), s.flags)
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *SliderFloat2) Changed() bool { return s.changed }

var _ imgui.Widget = (*SliderFloat2)(nil)

// SliderFloat3 edits a bound 3-component float vector by dragging within [Min, Max].
type SliderFloat3 struct {
	Label    string
	Value    *[3]float32
	Min, Max float32
	Format   string // default "%.3f"
	sliderFlags
	OnChange func([3]float32)
	changed  bool
	scratch  [3]float32
}

// NewSliderFloat3 returns a 3-component float slider bound to value.
func NewSliderFloat3(label string, value *[3]float32, min, max float32) *SliderFloat3 {
	return &SliderFloat3{Label: label, Value: value, Min: min, Max: max}
}

// Display draws the slider.
func (s *SliderFloat3) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	s.changed = cimgui.SliderFloat3(s.Label, v, s.Min, s.Max, sliderFormat(s.Format, "%.3f"), s.flags)
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *SliderFloat3) Changed() bool { return s.changed }

var _ imgui.Widget = (*SliderFloat3)(nil)

// SliderFloat4 edits a bound 4-component float vector by dragging within [Min, Max].
type SliderFloat4 struct {
	Label    string
	Value    *[4]float32
	Min, Max float32
	Format   string // default "%.3f"
	sliderFlags
	OnChange func([4]float32)
	changed  bool
	scratch  [4]float32
}

// NewSliderFloat4 returns a 4-component float slider bound to value.
func NewSliderFloat4(label string, value *[4]float32, min, max float32) *SliderFloat4 {
	return &SliderFloat4{Label: label, Value: value, Min: min, Max: max}
}

// Display draws the slider.
func (s *SliderFloat4) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	s.changed = cimgui.SliderFloat4(s.Label, v, s.Min, s.Max, sliderFormat(s.Format, "%.3f"), s.flags)
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *SliderFloat4) Changed() bool { return s.changed }

var _ imgui.Widget = (*SliderFloat4)(nil)

// SliderInt2 edits a bound 2-component int vector by dragging within [Min, Max].
type SliderInt2 struct {
	Label    string
	Value    *[2]int32
	Min, Max int32
	Format   string // default "%d"
	sliderFlags
	OnChange func([2]int32)
	changed  bool
	scratch  [2]int32
}

// NewSliderInt2 returns a 2-component int slider bound to value.
func NewSliderInt2(label string, value *[2]int32, min, max int32) *SliderInt2 {
	return &SliderInt2{Label: label, Value: value, Min: min, Max: max}
}

// Display draws the slider.
func (s *SliderInt2) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	s.changed = cimgui.SliderInt2(s.Label, v, s.Min, s.Max, sliderFormat(s.Format, "%d"), s.flags)
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *SliderInt2) Changed() bool { return s.changed }

var _ imgui.Widget = (*SliderInt2)(nil)

// SliderInt3 edits a bound 3-component int vector by dragging within [Min, Max].
type SliderInt3 struct {
	Label    string
	Value    *[3]int32
	Min, Max int32
	Format   string // default "%d"
	sliderFlags
	OnChange func([3]int32)
	changed  bool
	scratch  [3]int32
}

// NewSliderInt3 returns a 3-component int slider bound to value.
func NewSliderInt3(label string, value *[3]int32, min, max int32) *SliderInt3 {
	return &SliderInt3{Label: label, Value: value, Min: min, Max: max}
}

// Display draws the slider.
func (s *SliderInt3) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	s.changed = cimgui.SliderInt3(s.Label, v, s.Min, s.Max, sliderFormat(s.Format, "%d"), s.flags)
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *SliderInt3) Changed() bool { return s.changed }

var _ imgui.Widget = (*SliderInt3)(nil)

// SliderInt4 edits a bound 4-component int vector by dragging within [Min, Max].
type SliderInt4 struct {
	Label    string
	Value    *[4]int32
	Min, Max int32
	Format   string // default "%d"
	sliderFlags
	OnChange func([4]int32)
	changed  bool
	scratch  [4]int32
}

// NewSliderInt4 returns a 4-component int slider bound to value.
func NewSliderInt4(label string, value *[4]int32, min, max int32) *SliderInt4 {
	return &SliderInt4{Label: label, Value: value, Min: min, Max: max}
}

// Display draws the slider.
func (s *SliderInt4) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	s.changed = cimgui.SliderInt4(s.Label, v, s.Min, s.Max, sliderFormat(s.Format, "%d"), s.flags)
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *SliderInt4) Changed() bool { return s.changed }

var _ imgui.Widget = (*SliderInt4)(nil)

// SliderAngle edits a bound angle stored in radians, shown and dragged in degrees
// within [MinDegrees, MaxDegrees].
type SliderAngle struct {
	Label                  string
	Value                  *float32 // radians
	MinDegrees, MaxDegrees float32  // default -360, +360
	Format                 string   // default "%.0f deg"
	sliderFlags
	OnChange func(float32)
	changed  bool
	scratch  float32
}

// NewSliderAngle returns an angle slider bound to value (radians).
func NewSliderAngle(label string, value *float32) *SliderAngle {
	return &SliderAngle{Label: label, Value: value, MinDegrees: -360, MaxDegrees: 360}
}

// Display draws the slider.
func (s *SliderAngle) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	s.changed = cimgui.SliderAngle(s.Label, v, s.MinDegrees, s.MaxDegrees, sliderFormat(s.Format, "%.0f deg"), s.flags)
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *SliderAngle) Changed() bool { return s.changed }

var _ imgui.Widget = (*SliderAngle)(nil)

// VSliderFloat edits a bound float with a vertical slider of the given Size.
type VSliderFloat struct {
	Label    string
	Size     imgui.Vec2
	Value    *float32
	Min, Max float32
	Format   string // default "%.3f"
	sliderFlags
	OnChange func(float32)
	changed  bool
	scratch  float32
}

// NewVSliderFloat returns a vertical float slider bound to value.
func NewVSliderFloat(label string, size imgui.Vec2, value *float32, min, max float32) *VSliderFloat {
	return &VSliderFloat{Label: label, Size: size, Value: value, Min: min, Max: max}
}

// Display draws the slider.
func (s *VSliderFloat) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	s.changed = cimgui.VSliderFloat(s.Label, s.Size, v, s.Min, s.Max, sliderFormat(s.Format, "%.3f"), s.flags)
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *VSliderFloat) Changed() bool { return s.changed }

var _ imgui.Widget = (*VSliderFloat)(nil)

// VSliderInt edits a bound int with a vertical slider of the given Size.
type VSliderInt struct {
	Label    string
	Size     imgui.Vec2
	Value    *int32
	Min, Max int32
	Format   string // default "%d"
	sliderFlags
	OnChange func(int32)
	changed  bool
	scratch  int32
}

// NewVSliderInt returns a vertical int slider bound to value.
func NewVSliderInt(label string, size imgui.Vec2, value *int32, min, max int32) *VSliderInt {
	return &VSliderInt{Label: label, Size: size, Value: value, Min: min, Max: max}
}

// Display draws the slider.
func (s *VSliderInt) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	s.changed = cimgui.VSliderInt(s.Label, s.Size, v, s.Min, s.Max, sliderFormat(s.Format, "%d"), s.flags)
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *VSliderInt) Changed() bool { return s.changed }

var _ imgui.Widget = (*VSliderInt)(nil)

// sliderFormat returns format, or fallback when format is empty.
func sliderFormat(format, fallback string) string {
	if format == "" {
		return fallback
	}
	return format
}
