// Package combo provides combo boxes, list boxes and selectable items.
package combo

import (
	"github.com/bitwizeshift/go-imgui"
	"github.com/bitwizeshift/go-imgui/internal/cimgui"
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

// Func is a drop-down whose Count labels are produced lazily by Getter, selecting
// one by index bound to Value. Use it when the items are computed rather than held
// in a slice.
type Func struct {
	Label    string
	Value    *int32
	Count    int32
	Getter   func(i int32) string
	OnChange func(int32)
	changed  bool
	scratch  int32
}

// NewFunc returns a combo box of count items sourced from getter, bound to value.
func NewFunc(label string, value *int32, count int32, getter func(i int32) string) *Func {
	return &Func{Label: label, Value: value, Count: count, Getter: getter}
}

// Display draws the combo box.
func (c *Func) Display() {
	v := c.Value
	if v == nil {
		v = &c.scratch
	}
	c.changed = cimgui.Combo_FnStrPtr(c.Label, v, c.Getter, c.Count, -1)
	if c.changed && c.OnChange != nil {
		c.OnChange(*v)
	}
}

// Changed reports whether the selection changed during the last Display.
func (c *Func) Changed() bool { return c.changed }

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

// Custom is a combo box whose open drop-down is an arbitrary set of child widgets
// (typically Selectables). Preview is the text shown in the closed box.
type Custom struct {
	Label   string
	Preview string
	Widgets []imgui.Widget
}

// NewCustom returns a combo box with the given preview text and an arbitrary body.
func NewCustom(label, preview string) *Custom {
	return &Custom{Label: label, Preview: preview}
}

// AddWidget appends widgets to the drop-down body.
func (c *Custom) AddWidget(ws ...imgui.Widget) { c.Widgets = append(c.Widgets, ws...) }

// Display draws the combo box and, while open, its body.
func (c *Custom) Display() {
	if cimgui.BeginCombo(c.Label, c.Preview, cimgui.ComboFlagsNone) {
		for _, w := range c.Widgets {
			if w != nil {
				w.Display()
			}
		}
		cimgui.EndCombo()
	}
}

var _ imgui.Widget = (*Custom)(nil)

// CustomList is a scrolling list box whose body is an arbitrary set of child
// widgets (typically Selectables). A zero Size auto-fits.
type CustomList struct {
	Label   string
	Size    imgui.Vec2
	Widgets []imgui.Widget
}

// NewCustomList returns a list box with an arbitrary body.
func NewCustomList(label string) *CustomList { return &CustomList{Label: label} }

// AddWidget appends widgets to the list body.
func (l *CustomList) AddWidget(ws ...imgui.Widget) { l.Widgets = append(l.Widgets, ws...) }

// Display draws the list box and, while open, its body.
func (l *CustomList) Display() {
	if cimgui.BeginListBox(l.Label, l.Size) {
		for _, w := range l.Widgets {
			if w != nil {
				w.Display()
			}
		}
		cimgui.EndListBox()
	}
}

var _ imgui.Widget = (*CustomList)(nil)

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
