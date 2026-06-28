package input

import (
	"unsafe"

	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// Scalarable is the set of numeric types the generic scalar widgets accept. Each
// maps onto a cimgui.DataType.
type Scalarable interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

// dataType returns the cimgui.DataType describing T.
func dataType[T Scalarable]() cimgui.DataType {
	switch any(*new(T)).(type) {
	case int8:
		return cimgui.DataTypeS8
	case uint8:
		return cimgui.DataTypeU8
	case int16:
		return cimgui.DataTypeS16
	case uint16:
		return cimgui.DataTypeU16
	case int32:
		return cimgui.DataTypeS32
	case uint32:
		return cimgui.DataTypeU32
	case int64:
		return cimgui.DataTypeS64
	case uint64:
		return cimgui.DataTypeU64
	case float32:
		return cimgui.DataTypeFloat
	default:
		return cimgui.DataTypeDouble
	}
}

// scalarFormat returns format, or a type-appropriate default when format is empty.
func scalarFormat[T Scalarable](format string) string {
	if format != "" {
		return format
	}
	switch any(*new(T)).(type) {
	case float32, float64:
		return "%.3f"
	default:
		return "%d"
	}
}

// optScalarPtr returns a pointer to v, or nil when v is the zero value (so the
// optional step/fast parameters default in ImGui).
func optScalarPtr[T Scalarable](v *T) unsafe.Pointer {
	if *v == 0 {
		return nil
	}
	return unsafe.Pointer(v)
}

// Scalar edits a bound value of any [Scalarable] type in a box with optional
// +/- step buttons.
type Scalar[T Scalarable] struct {
	Label    string
	Value    *T
	Step     T
	StepFast T
	Format   string // default per type ("%d" or "%.3f")
	textFlags
	OnChange func(T)
	changed  bool
	scratch  T
}

// NewScalar returns a scalar input bound to value.
func NewScalar[T Scalarable](label string, value *T) *Scalar[T] {
	return &Scalar[T]{Label: label, Value: value}
}

// Display draws the input.
func (s *Scalar[T]) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	// Pass a pointer to a standalone copy: cgo's pointer check deep-scans the
	// object an unsafe.Pointer points into, so it must not point into a struct
	// that holds Go pointers.
	val := *v
	step, stepFast := s.Step, s.StepFast
	s.changed = cimgui.InputScalar(s.Label, dataType[T](), unsafe.Pointer(&val),
		optScalarPtr(&step), optScalarPtr(&stepFast), scalarFormat[T](s.Format), s.flags)
	*v = val
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *Scalar[T]) Changed() bool { return s.changed }

var _ imgui.Widget = (*Scalar[int32])(nil)

// ScalarN edits a bound slice of any [Scalarable] type as one widget per element.
type ScalarN[T Scalarable] struct {
	Label  string
	Value  []T
	Format string // default per type
	textFlags
	OnChange func([]T)
	changed  bool
}

// NewScalarN returns a multi-component scalar input bound to value.
func NewScalarN[T Scalarable](label string, value []T) *ScalarN[T] {
	return &ScalarN[T]{Label: label, Value: value}
}

// Display draws the inputs.
func (s *ScalarN[T]) Display() {
	if len(s.Value) == 0 {
		return
	}
	s.changed = cimgui.InputScalarN(s.Label, dataType[T](), unsafe.Pointer(&s.Value[0]),
		int32(len(s.Value)), nil, nil, scalarFormat[T](s.Format), s.flags)
	if s.changed && s.OnChange != nil {
		s.OnChange(s.Value)
	}
}

// Changed reports whether any component changed during the last Display.
func (s *ScalarN[T]) Changed() bool { return s.changed }

var _ imgui.Widget = (*ScalarN[int32])(nil)

// SliderScalar edits a bound [Scalarable] value by dragging within [Min, Max].
type SliderScalar[T Scalarable] struct {
	Label    string
	Value    *T
	Min, Max T
	Format   string // default per type
	sliderFlags
	OnChange func(T)
	changed  bool
	scratch  T
}

// NewSliderScalar returns a scalar slider bound to value.
func NewSliderScalar[T Scalarable](label string, value *T, min, max T) *SliderScalar[T] {
	return &SliderScalar[T]{Label: label, Value: value, Min: min, Max: max}
}

// Display draws the slider.
func (s *SliderScalar[T]) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	val, min, max := *v, s.Min, s.Max
	s.changed = cimgui.SliderScalar(s.Label, dataType[T](), unsafe.Pointer(&val),
		unsafe.Pointer(&min), unsafe.Pointer(&max), scalarFormat[T](s.Format), s.flags)
	*v = val
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *SliderScalar[T]) Changed() bool { return s.changed }

var _ imgui.Widget = (*SliderScalar[int32])(nil)

// SliderScalarN edits a bound slice of [Scalarable] values within [Min, Max].
type SliderScalarN[T Scalarable] struct {
	Label    string
	Value    []T
	Min, Max T
	Format   string // default per type
	sliderFlags
	OnChange func([]T)
	changed  bool
}

// NewSliderScalarN returns a multi-component scalar slider bound to value.
func NewSliderScalarN[T Scalarable](label string, value []T, min, max T) *SliderScalarN[T] {
	return &SliderScalarN[T]{Label: label, Value: value, Min: min, Max: max}
}

// Display draws the slider.
func (s *SliderScalarN[T]) Display() {
	if len(s.Value) == 0 {
		return
	}
	min, max := s.Min, s.Max
	s.changed = cimgui.SliderScalarN(s.Label, dataType[T](), unsafe.Pointer(&s.Value[0]),
		int32(len(s.Value)), unsafe.Pointer(&min), unsafe.Pointer(&max), scalarFormat[T](s.Format), s.flags)
	if s.changed && s.OnChange != nil {
		s.OnChange(s.Value)
	}
}

// Changed reports whether any component changed during the last Display.
func (s *SliderScalarN[T]) Changed() bool { return s.changed }

var _ imgui.Widget = (*SliderScalarN[int32])(nil)

// DragScalar edits a bound [Scalarable] value by dragging. Min==Max leaves it
// unbounded.
type DragScalar[T Scalarable] struct {
	Label    string
	Value    *T
	Speed    float32 // default 1
	Min, Max T
	Format   string // default per type
	sliderFlags
	OnChange func(T)
	changed  bool
	scratch  T
}

// NewDragScalar returns a draggable scalar bound to value.
func NewDragScalar[T Scalarable](label string, value *T) *DragScalar[T] {
	return &DragScalar[T]{Label: label, Value: value}
}

// Display draws the drag.
func (d *DragScalar[T]) Display() {
	v := d.Value
	if v == nil {
		v = &d.scratch
	}
	val, min, max := *v, d.Min, d.Max
	d.changed = cimgui.DragScalar(d.Label, dataType[T](), unsafe.Pointer(&val), dragSpeed(d.Speed),
		unsafe.Pointer(&min), unsafe.Pointer(&max), scalarFormat[T](d.Format), d.flags)
	*v = val
	if d.changed && d.OnChange != nil {
		d.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (d *DragScalar[T]) Changed() bool { return d.changed }

var _ imgui.Widget = (*DragScalar[int32])(nil)

// DragScalarN edits a bound slice of [Scalarable] values by dragging.
type DragScalarN[T Scalarable] struct {
	Label    string
	Value    []T
	Speed    float32 // default 1
	Min, Max T
	Format   string // default per type
	sliderFlags
	OnChange func([]T)
	changed  bool
}

// NewDragScalarN returns a multi-component draggable scalar bound to value.
func NewDragScalarN[T Scalarable](label string, value []T) *DragScalarN[T] {
	return &DragScalarN[T]{Label: label, Value: value}
}

// Display draws the drag.
func (d *DragScalarN[T]) Display() {
	if len(d.Value) == 0 {
		return
	}
	min, max := d.Min, d.Max
	d.changed = cimgui.DragScalarN(d.Label, dataType[T](), unsafe.Pointer(&d.Value[0]),
		int32(len(d.Value)), dragSpeed(d.Speed), unsafe.Pointer(&min), unsafe.Pointer(&max),
		scalarFormat[T](d.Format), d.flags)
	if d.changed && d.OnChange != nil {
		d.OnChange(d.Value)
	}
}

// Changed reports whether any component changed during the last Display.
func (d *DragScalarN[T]) Changed() bool { return d.changed }

var _ imgui.Widget = (*DragScalarN[int32])(nil)

// VSliderScalar edits a bound [Scalarable] value with a vertical slider of the
// given Size.
type VSliderScalar[T Scalarable] struct {
	Label    string
	Size     imgui.Vec2
	Value    *T
	Min, Max T
	Format   string // default per type
	sliderFlags
	OnChange func(T)
	changed  bool
	scratch  T
}

// NewVSliderScalar returns a vertical scalar slider bound to value.
func NewVSliderScalar[T Scalarable](label string, size imgui.Vec2, value *T, min, max T) *VSliderScalar[T] {
	return &VSliderScalar[T]{Label: label, Size: size, Value: value, Min: min, Max: max}
}

// Display draws the slider.
func (s *VSliderScalar[T]) Display() {
	v := s.Value
	if v == nil {
		v = &s.scratch
	}
	val, min, max := *v, s.Min, s.Max
	s.changed = cimgui.VSliderScalar(s.Label, s.Size, dataType[T](), unsafe.Pointer(&val),
		unsafe.Pointer(&min), unsafe.Pointer(&max), scalarFormat[T](s.Format), s.flags)
	*v = val
	if s.changed && s.OnChange != nil {
		s.OnChange(*v)
	}
}

// Changed reports whether the value changed during the last Display.
func (s *VSliderScalar[T]) Changed() bool { return s.changed }

var _ imgui.Widget = (*VSliderScalar[int32])(nil)
