// Package menu provides menu bars, menus and menu items.
package menu

import (
	"github.com/bitwizeshift/go-imgui"
	"github.com/bitwizeshift/go-imgui/internal/cimgui"
)

// MainBar is the full-width menu bar across the top of the viewport. Add [Menu]s
// to it.
type MainBar struct {
	Menus []imgui.Widget
}

// NewMainBar returns a main menu bar.
func NewMainBar() *MainBar { return &MainBar{} }

// AddWidget appends menus.
func (m *MainBar) AddWidget(ws ...imgui.Widget) { m.Menus = append(m.Menus, ws...) }

// SetLayout replaces the menus.
func (m *MainBar) SetLayout(ws ...imgui.Widget) { m.Menus = ws }

// Display draws the main menu bar.
func (m *MainBar) Display() {
	if cimgui.BeginMainMenuBar() {
		displayAll(m.Menus)
		cimgui.EndMainMenuBar()
	}
}

// Bar is the menu bar of the current window. The window must enable its menu bar
// via (*window.Window).SetMenuBar(true).
type Bar struct {
	Menus []imgui.Widget
}

// NewBar returns a window menu bar.
func NewBar() *Bar { return &Bar{} }

// AddWidget appends menus.
func (b *Bar) AddWidget(ws ...imgui.Widget) { b.Menus = append(b.Menus, ws...) }

// SetLayout replaces the menus.
func (b *Bar) SetLayout(ws ...imgui.Widget) { b.Menus = ws }

// Display draws the window menu bar.
func (b *Bar) Display() {
	if cimgui.BeginMenuBar() {
		displayAll(b.Menus)
		cimgui.EndMenuBar()
	}
}

// Menu is a drop-down menu containing [Item]s and nested menus.
type Menu struct {
	Label    string
	Disabled bool
	Items    []imgui.Widget
}

// New returns a menu labelled label.
func New(label string) *Menu { return &Menu{Label: label} }

// AddWidget appends items (and nested menus).
func (m *Menu) AddWidget(ws ...imgui.Widget) { m.Items = append(m.Items, ws...) }

// SetLayout replaces the items.
func (m *Menu) SetLayout(ws ...imgui.Widget) { m.Items = ws }

// Display draws the menu and, when open, its items.
func (m *Menu) Display() {
	if cimgui.BeginMenu(m.Label, !m.Disabled) {
		displayAll(m.Items)
		cimgui.EndMenu()
	}
}

// Item is a menu entry. When Selected is non-nil it shows a check mark bound to
// that bool. OnClick (if set) fires on activation; poll with [Item.Clicked].
type Item struct {
	Label    string
	Shortcut string
	Selected *bool
	Disabled bool
	OnClick  func()
	clicked  bool
}

// NewItem returns a menu item.
func NewItem(label string) *Item { return &Item{Label: label} }

// Display draws the menu item.
func (i *Item) Display() {
	if i.Selected != nil {
		i.clicked = cimgui.MenuItem_BoolPtr(i.Label, i.Shortcut, i.Selected, !i.Disabled)
	} else {
		i.clicked = cimgui.MenuItem_Bool(i.Label, i.Shortcut, false, !i.Disabled)
	}
	if i.clicked && i.OnClick != nil {
		i.OnClick()
	}
}

// Clicked reports whether the item was activated during the last Display.
func (i *Item) Clicked() bool { return i.clicked }

func displayAll(ws []imgui.Widget) {
	for _, w := range ws {
		if w != nil {
			w.Display()
		}
	}
}
