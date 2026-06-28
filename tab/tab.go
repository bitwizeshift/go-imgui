// Package tab provides tab bars and tab items.
package tab

import (
	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// Side places a tab at the leading or trailing edge of the bar. The zero value
// leaves the tab in normal flow.
type Side int

const (
	// SideNone leaves the tab in normal left-to-right flow.
	SideNone Side = iota
	// SideLeading pins the tab to the leading (left) edge.
	SideLeading
	// SideTrailing pins the tab to the trailing (right) edge.
	SideTrailing
)

func setBarFlag(f *cimgui.TabBarFlags, bit cimgui.TabBarFlags, on bool) {
	if on {
		*f |= bit
	} else {
		*f &^= bit
	}
}

func setItemFlag(f *cimgui.TabItemFlags, bit cimgui.TabItemFlags, on bool) {
	if on {
		*f |= bit
	} else {
		*f &^= bit
	}
}

// Bar is a tab bar containing tab [Item]s. ID must be unique.
type Bar struct {
	ID    string
	Items []imgui.Widget
	flags cimgui.TabBarFlags
}

// NewBar returns a tab bar.
func NewBar(id string) *Bar { return &Bar{ID: id} }

// SetReorderable allows the user to drag tabs to reorder them.
func (b *Bar) SetReorderable(on bool) {
	setBarFlag(&b.flags, cimgui.TabBarFlagsReorderable, on)
}

// SetAutoSelectNewTabs selects a tab the frame it first appears.
func (b *Bar) SetAutoSelectNewTabs(on bool) {
	setBarFlag(&b.flags, cimgui.TabBarFlagsAutoSelectNewTabs, on)
}

// SetFittingScroll keeps tabs full width and adds scroll buttons when they
// overflow instead of shrinking them.
func (b *Bar) SetFittingScroll(on bool) {
	setBarFlag(&b.flags, cimgui.TabBarFlagsFittingPolicyScroll, on)
}

// AddWidget appends items (usually *Item) to the bar.
func (b *Bar) AddWidget(ws ...imgui.Widget) { b.Items = append(b.Items, ws...) }

// SetLayout replaces the bar's items.
func (b *Bar) SetLayout(ws ...imgui.Widget) { b.Items = ws }

// Display draws the tab bar.
func (b *Bar) Display() {
	if cimgui.BeginTabBar(b.ID, b.flags) {
		for _, it := range b.Items {
			if it != nil {
				it.Display()
			}
		}
		cimgui.EndTabBar()
	}
}

// ItemButton is a button that sits in a tab bar's tab row (alongside the tabs)
// rather than being a selectable tab. Add it to a [Bar].
type ItemButton struct {
	Label   string
	OnClick func()
	flags   cimgui.TabItemFlags
	clicked bool
}

// NewItemButton returns a tab-bar button.
func NewItemButton(label string) *ItemButton { return &ItemButton{Label: label} }

// SetSide pins the button to the leading or trailing edge of the bar.
func (b *ItemButton) SetSide(s Side) {
	b.flags &^= cimgui.TabItemFlagsLeading | cimgui.TabItemFlagsTrailing
	switch s {
	case SideLeading:
		b.flags |= cimgui.TabItemFlagsLeading
	case SideTrailing:
		b.flags |= cimgui.TabItemFlagsTrailing
	}
}

// Display draws the button. Place it inside a [Bar].
func (b *ItemButton) Display() {
	b.clicked = cimgui.TabItemButton(b.Label, b.flags)
	if b.clicked && b.OnClick != nil {
		b.OnClick()
	}
}

// Clicked reports whether the button was clicked during the last Display.
func (b *ItemButton) Clicked() bool { return b.clicked }

var _ imgui.Widget = (*ItemButton)(nil)

// Item is a single tab. When Open is non-nil the tab shows a close button and is
// cleared when closed. Children are drawn only while the tab is selected.
type Item struct {
	Label   string
	Open    *bool
	Widgets []imgui.Widget
	flags   cimgui.TabItemFlags
}

// NewItem returns a tab item.
func NewItem(label string) *Item { return &Item{Label: label} }

// SetSelected forces this tab to be selected on the next frame.
func (i *Item) SetSelected(on bool) {
	setItemFlag(&i.flags, cimgui.TabItemFlagsSetSelected, on)
}

// SetUnsavedDocument shows the unsaved-document marker on the tab.
func (i *Item) SetUnsavedDocument(on bool) {
	setItemFlag(&i.flags, cimgui.TabItemFlagsUnsavedDocument, on)
}

// SetSide pins the tab to the leading or trailing edge of the bar.
func (i *Item) SetSide(s Side) {
	i.flags &^= cimgui.TabItemFlagsLeading | cimgui.TabItemFlagsTrailing
	switch s {
	case SideLeading:
		i.flags |= cimgui.TabItemFlagsLeading
	case SideTrailing:
		i.flags |= cimgui.TabItemFlagsTrailing
	}
}

// AddWidget appends child widgets.
func (i *Item) AddWidget(ws ...imgui.Widget) { i.Widgets = append(i.Widgets, ws...) }

// SetLayout replaces the child widgets.
func (i *Item) SetLayout(ws ...imgui.Widget) { i.Widgets = ws }

// Display draws the tab and, when selected, its children.
func (i *Item) Display() {
	if cimgui.BeginTabItem(i.Label, i.Open, i.flags) {
		for _, w := range i.Widgets {
			if w != nil {
				w.Display()
			}
		}
		cimgui.EndTabItem()
	}
}
