// Package combo provides combo boxes, list boxes and selectable items.
package combo

import (
	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// Combo is a drop-down selecting one of Items by index, bound to Value.
type Combo struct {
	Label    string
	Value    *int32
	Items    []string
	OnChange func(int32)
	changed  bool
	scratch  int32
}

// New returns a combo box bound to value (the selected index).
func New(label string, value *int32, items []string) *Combo {
	return &Combo{Label: label, Value: value, Items: items}
}

// Display draws the combo box.
func (c *Combo) Display() {
	v := c.Value
	if v == nil {
		v = &c.scratch
	}
	c.changed = cimgui.Combo_Str_arr(c.Label, v, c.Items, -1)
	if c.changed && c.OnChange != nil {
		c.OnChange(*v)
	}
}

// Changed reports whether the selection changed during the last Display.
func (c *Combo) Changed() bool { return c.changed }

// ListBox is a scrolling list selecting one of Items by index, bound to Value.
type ListBox struct {
	Label       string
	Value       *int32
	Items       []string
	HeightItems int32 // visible rows; <=0 uses the default
	OnChange    func(int32)
	changed     bool
	scratch     int32
}

// NewListBox returns a list box bound to value (the selected index).
func NewListBox(label string, value *int32, items []string) *ListBox {
	return &ListBox{Label: label, Value: value, Items: items, HeightItems: -1}
}

// Display draws the list box.
func (l *ListBox) Display() {
	v := l.Value
	if v == nil {
		v = &l.scratch
	}
	h := l.HeightItems
	if h == 0 {
		h = -1
	}
	l.changed = cimgui.ListBox_Str_arr(l.Label, v, l.Items, h)
	if l.changed && l.OnChange != nil {
		l.OnChange(*v)
	}
}

// Changed reports whether the selection changed during the last Display.
func (l *ListBox) Changed() bool { return l.changed }

// Selectable is a clickable, selectable line of text. When Selected is non-nil it
// is toggled on click; OnClick (if set) fires on click. Poll with [Selectable.Clicked].
type Selectable struct {
	Label    string
	Selected *bool
	OnClick  func()
	clicked  bool
}

// NewSelectable returns a selectable item.
func NewSelectable(label string) *Selectable { return &Selectable{Label: label} }

// Display draws the selectable.
func (s *Selectable) Display() {
	if s.Selected != nil {
		s.clicked = cimgui.Selectable_BoolPtr(s.Label, s.Selected, cimgui.SelectableFlagsNone, imgui.Vec2{})
	} else {
		s.clicked = cimgui.Selectable_Bool(s.Label, false, cimgui.SelectableFlagsNone, imgui.Vec2{})
	}
	if s.clicked && s.OnClick != nil {
		s.OnClick()
	}
}

// Clicked reports whether the item was clicked during the last Display.
func (s *Selectable) Clicked() bool { return s.clicked }
