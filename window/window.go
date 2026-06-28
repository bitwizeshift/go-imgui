// Package window provides top-level windows and child regions.
package window

import (
	"github.com/bitwizeshift/go-imgui"
	"github.com/bitwizeshift/go-imgui/internal/cimgui"
)

// flagSetters is the shared cached window bitfield plus its setters, embedded by
// both Window and Child. Setters update the cached field when called so Display
// never recomputes it.
type flagSetters struct {
	flags cimgui.WindowFlags
}

func setWindowFlag(f *cimgui.WindowFlags, bit cimgui.WindowFlags, on bool) {
	if on {
		*f |= bit
	} else {
		*f &^= bit
	}
}

// SetTitleBar shows or hides the title bar (default shown).
func (f *flagSetters) SetTitleBar(on bool) {
	setWindowFlag(&f.flags, cimgui.WindowFlagsNoTitleBar, !on)
}

// SetResizable allows or prevents the user resizing the window (default allowed).
func (f *flagSetters) SetResizable(on bool) {
	setWindowFlag(&f.flags, cimgui.WindowFlagsNoResize, !on)
}

// SetMovable allows or prevents the user moving the window (default allowed).
func (f *flagSetters) SetMovable(on bool) {
	setWindowFlag(&f.flags, cimgui.WindowFlagsNoMove, !on)
}

// SetScrollbar shows or hides the scrollbar (default shown).
func (f *flagSetters) SetScrollbar(on bool) {
	setWindowFlag(&f.flags, cimgui.WindowFlagsNoScrollbar, !on)
}

// SetCollapsible allows or prevents collapsing via the title bar (default allowed).
func (f *flagSetters) SetCollapsible(on bool) {
	setWindowFlag(&f.flags, cimgui.WindowFlagsNoCollapse, !on)
}

// SetAutoResize makes the window always resize to fit its content.
func (f *flagSetters) SetAutoResize(on bool) {
	setWindowFlag(&f.flags, cimgui.WindowFlagsAlwaysAutoResize, on)
}

// SetBackground draws or omits the window background (default drawn).
func (f *flagSetters) SetBackground(on bool) {
	setWindowFlag(&f.flags, cimgui.WindowFlagsNoBackground, !on)
}

// SetMenuBar reserves space for a menu bar.
func (f *flagSetters) SetMenuBar(on bool) {
	setWindowFlag(&f.flags, cimgui.WindowFlagsMenuBar, on)
}

// SetHorizontalScrollbar allows horizontal scrolling.
func (f *flagSetters) SetHorizontalScrollbar(on bool) {
	setWindowFlag(&f.flags, cimgui.WindowFlagsHorizontalScrollbar, on)
}

// SetNav enables keyboard/gamepad navigation within the window (default enabled).
func (f *flagSetters) SetNav(on bool) {
	setWindowFlag(&f.flags, cimgui.WindowFlagsNoNav, !on)
}

// SetDecorated shows or hides the window decorations: title bar, border and
// scrollbar (default shown).
func (f *flagSetters) SetDecorated(on bool) {
	setWindowFlag(&f.flags, cimgui.WindowFlagsNoDecoration, !on)
}

// Window is a top-level window. When Open is non-nil the title bar shows a close
// button and Open is cleared when the user closes it. Pos/Size, when set, are
// applied on first use so the user can still move/resize afterwards.
type Window struct {
	Title   string
	Open    *bool
	Pos     *imgui.Vec2
	Size    *imgui.Vec2
	Widgets []imgui.Widget
	flagSetters
}

// New returns a window with the given title.
func New(title string) *Window { return &Window{Title: title} }

// AddWidget appends widgets to the window.
func (w *Window) AddWidget(ws ...imgui.Widget) { w.Widgets = append(w.Widgets, ws...) }

// SetLayout replaces the window's widgets.
func (w *Window) SetLayout(ws ...imgui.Widget) { w.Widgets = ws }

// Display draws the window and its children.
func (w *Window) Display() {
	if w.Pos != nil {
		cimgui.SetNextWindowPos(*w.Pos, cimgui.CondFirstUseEver, imgui.Vec2{})
	}
	if w.Size != nil {
		cimgui.SetNextWindowSize(*w.Size, cimgui.CondFirstUseEver)
	}
	open := cimgui.Begin(w.Title, w.Open, w.flags)
	if open {
		for _, c := range w.Widgets {
			if c != nil {
				c.Display()
			}
		}
	}
	cimgui.End()
}

// Child is a scrollable, optionally bordered sub-region. ID must be unique.
type Child struct {
	ID      string
	Size    imgui.Vec2 // zero fills available space
	Border  bool
	Widgets []imgui.Widget
	flagSetters
}

// NewChild returns a child region.
func NewChild(id string) *Child { return &Child{ID: id} }

// AddWidget appends widgets to the child.
func (c *Child) AddWidget(ws ...imgui.Widget) { c.Widgets = append(c.Widgets, ws...) }

// SetLayout replaces the child's widgets.
func (c *Child) SetLayout(ws ...imgui.Widget) { c.Widgets = ws }

// Display draws the child region and its children.
func (c *Child) Display() {
	var cf cimgui.ChildFlags
	if c.Border {
		cf |= cimgui.ChildFlagsBorders
	}
	visible := cimgui.BeginChild_Str(c.ID, c.Size, cf, c.flags)
	if visible {
		for _, w := range c.Widgets {
			if w != nil {
				w.Display()
			}
		}
	}
	cimgui.EndChild()
}
