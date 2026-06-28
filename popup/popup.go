// Package popup provides popups, modal dialogs and context menus.
package popup

import (
	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// Open marks the popup with the given id to open. Call it from an event handler
// such as a button's OnClick; the matching [Popup] or [Modal] then appears.
func Open(id string) { cimgui.OpenPopup_Str(id, cimgui.PopupFlagsNone) }

// IsOpen reports whether the popup with the given id is currently open.
func IsOpen(id string) bool { return cimgui.IsPopupOpen_Str(id, cimgui.PopupFlagsNone) }

// CloseCurrent closes the popup currently being displayed. Call it from within a
// popup's layout (e.g. an OK button's OnClick).
func CloseCurrent() { cimgui.CloseCurrentPopup() }

// Popup is a (non-modal) popup window keyed by ID. Trigger it with [Open]. Its
// children are drawn only while it is open.
type Popup struct {
	ID      string
	Widgets []imgui.Widget
}

// New returns a popup with the given id.
func New(id string) *Popup { return &Popup{ID: id} }

// AddWidget appends child widgets.
func (p *Popup) AddWidget(ws ...imgui.Widget) { p.Widgets = append(p.Widgets, ws...) }

// SetLayout replaces the child widgets.
func (p *Popup) SetLayout(ws ...imgui.Widget) { p.Widgets = ws }

// Display draws the popup when open.
func (p *Popup) Display() {
	if cimgui.BeginPopup(p.ID, cimgui.WindowFlagsNone) {
		displayAll(p.Widgets)
		cimgui.EndPopup()
	}
}

// Modal is a modal dialog keyed by Name. Trigger it with Open(Name). When Open is
// non-nil it shows a close button and is cleared when dismissed.
type Modal struct {
	Name    string
	Open    *bool
	Widgets []imgui.Widget
}

// NewModal returns a modal dialog with the given name.
func NewModal(name string) *Modal { return &Modal{Name: name} }

// AddWidget appends child widgets.
func (m *Modal) AddWidget(ws ...imgui.Widget) { m.Widgets = append(m.Widgets, ws...) }

// SetLayout replaces the child widgets.
func (m *Modal) SetLayout(ws ...imgui.Widget) { m.Widgets = ws }

// Display draws the modal when open.
func (m *Modal) Display() {
	if cimgui.BeginPopupModal(m.Name, m.Open, cimgui.WindowFlagsNone) {
		displayAll(m.Widgets)
		cimgui.EndPopup()
	}
}

// ContextItem is a context menu opened by right-clicking the previous widget. ID
// must be unique.
type ContextItem struct {
	ID      string
	Widgets []imgui.Widget
}

// NewContextItem returns a context menu.
func NewContextItem(id string) *ContextItem { return &ContextItem{ID: id} }

// AddWidget appends child widgets.
func (c *ContextItem) AddWidget(ws ...imgui.Widget) { c.Widgets = append(c.Widgets, ws...) }

// SetLayout replaces the child widgets.
func (c *ContextItem) SetLayout(ws ...imgui.Widget) { c.Widgets = ws }

// Display draws the context menu when triggered.
func (c *ContextItem) Display() {
	if cimgui.BeginPopupContextItem(c.ID, cimgui.PopupFlagsMouseButtonRight) {
		displayAll(c.Widgets)
		cimgui.EndPopup()
	}
}

func displayAll(ws []imgui.Widget) {
	for _, w := range ws {
		if w != nil {
			w.Display()
		}
	}
}
