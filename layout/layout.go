// Package layout provides spacing and grouping widgets.
package layout

import (
	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// widgetFunc adapts a plain function to the imgui.Widget interface.
type widgetFunc func()

func (f widgetFunc) Display() { f() }

// Separator draws a horizontal line.
func Separator() imgui.Widget { return widgetFunc(cimgui.Separator) }

// Spacing adds vertical spacing.
func Spacing() imgui.Widget { return widgetFunc(cimgui.Spacing) }

// NewLine moves to the next line.
func NewLine() imgui.Widget { return widgetFunc(cimgui.NewLine) }

// SameLine keeps the next widget on the current line.
func SameLine() imgui.Widget {
	return widgetFunc(func() { cimgui.SameLine(0, -1) })
}

// SameLineEx keeps the next widget on the current line at the given offset and
// spacing (negative spacing uses the default).
func SameLineEx(offsetFromStart, spacing float32) imgui.Widget {
	return widgetFunc(func() { cimgui.SameLine(offsetFromStart, spacing) })
}

// Dummy reserves an empty item of the given size.
func Dummy(size imgui.Vec2) imgui.Widget {
	return widgetFunc(func() { cimgui.Dummy(size) })
}

// Indent increases the indent for following widgets by width (0 = default). It is
// not scoped; pair with [Unindent] or use a [Group].
func Indent(width float32) imgui.Widget {
	return widgetFunc(func() { cimgui.Indent(width) })
}

// Unindent decreases the indent by width (0 = default).
func Unindent(width float32) imgui.Widget {
	return widgetFunc(func() { cimgui.Unindent(width) })
}

// Group lays out its children as a single item (for alignment and item queries).
type Group struct {
	Widgets []imgui.Widget
}

// NewGroup returns a group of the given widgets.
func NewGroup(ws ...imgui.Widget) *Group { return &Group{Widgets: ws} }

// AddWidget appends widgets to the group.
func (g *Group) AddWidget(ws ...imgui.Widget) { g.Widgets = append(g.Widgets, ws...) }

// SetLayout replaces the group's widgets.
func (g *Group) SetLayout(ws ...imgui.Widget) { g.Widgets = ws }

// Display draws the group.
func (g *Group) Display() {
	cimgui.BeginGroup()
	for _, w := range g.Widgets {
		w.Display()
	}
	cimgui.EndGroup()
}

// Disabled greys out and disables interaction with its children when Disabled is
// true.
type Disabled struct {
	Disabled bool
	Widgets  []imgui.Widget
}

// NewDisabled returns a disabled-scope container.
func NewDisabled(disabled bool, ws ...imgui.Widget) *Disabled {
	return &Disabled{Disabled: disabled, Widgets: ws}
}

// AddWidget appends widgets.
func (d *Disabled) AddWidget(ws ...imgui.Widget) { d.Widgets = append(d.Widgets, ws...) }

// SetLayout replaces the widgets.
func (d *Disabled) SetLayout(ws ...imgui.Widget) { d.Widgets = ws }

// Display draws the children inside a disabled scope.
func (d *Disabled) Display() {
	cimgui.BeginDisabled(d.Disabled)
	for _, w := range d.Widgets {
		w.Display()
	}
	cimgui.EndDisabled()
}
