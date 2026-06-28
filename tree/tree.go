// Package tree provides collapsible tree nodes and headers.
package tree

import (
	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// flagSetters is the shared cached tree bitfield plus its setters, embedded by
// both Node and Header. Setters update the cached field when called so Display
// never recomputes it.
type flagSetters struct {
	flags cimgui.TreeNodeFlags
}

func setTreeFlag(f *cimgui.TreeNodeFlags, bit cimgui.TreeNodeFlags, on bool) {
	if on {
		*f |= bit
	} else {
		*f &^= bit
	}
}

// SetDefaultOpen opens the node by default.
func (f *flagSetters) SetDefaultOpen(on bool) {
	setTreeFlag(&f.flags, cimgui.TreeNodeFlagsDefaultOpen, on)
}

// SetLeaf draws the node as a leaf with no expand arrow and no children.
func (f *flagSetters) SetLeaf(on bool) {
	setTreeFlag(&f.flags, cimgui.TreeNodeFlagsLeaf, on)
}

// SetBullet draws a bullet instead of an expand arrow.
func (f *flagSetters) SetBullet(on bool) {
	setTreeFlag(&f.flags, cimgui.TreeNodeFlagsBullet, on)
}

// SetFramed draws a full-width framed header.
func (f *flagSetters) SetFramed(on bool) {
	setTreeFlag(&f.flags, cimgui.TreeNodeFlagsFramed, on)
}

// SetSpanFullWidth makes the hit box span the full available width.
func (f *flagSetters) SetSpanFullWidth(on bool) {
	setTreeFlag(&f.flags, cimgui.TreeNodeFlagsSpanFullWidth, on)
}

// SetOpenOnArrow opens the node only when the arrow is clicked.
func (f *flagSetters) SetOpenOnArrow(on bool) {
	setTreeFlag(&f.flags, cimgui.TreeNodeFlagsOpenOnArrow, on)
}

// Node is a collapsible tree node. Its children are drawn only while it is open.
type Node struct {
	Label   string
	Widgets []imgui.Widget
	flagSetters
}

// New returns a tree node labelled label.
func New(label string) *Node { return &Node{Label: label} }

// AddWidget appends child widgets.
func (n *Node) AddWidget(ws ...imgui.Widget) { n.Widgets = append(n.Widgets, ws...) }

// SetLayout replaces the child widgets.
func (n *Node) SetLayout(ws ...imgui.Widget) { n.Widgets = ws }

// Display draws the node and, when open, its children.
func (n *Node) Display() {
	if cimgui.TreeNodeEx_Str(n.Label, n.flags) {
		for _, w := range n.Widgets {
			if w != nil {
				w.Display()
			}
		}
		cimgui.TreePop()
	}
}

// Header is a collapsing header. When Open is non-nil it shows a close button and
// is cleared when dismissed.
type Header struct {
	Label   string
	Open    *bool
	Widgets []imgui.Widget
	flagSetters
}

// NewHeader returns a collapsing header labelled label.
func NewHeader(label string) *Header { return &Header{Label: label} }

// AddWidget appends child widgets.
func (h *Header) AddWidget(ws ...imgui.Widget) { h.Widgets = append(h.Widgets, ws...) }

// SetLayout replaces the child widgets.
func (h *Header) SetLayout(ws ...imgui.Widget) { h.Widgets = ws }

// Display draws the header and, when open, its children.
func (h *Header) Display() {
	var open bool
	if h.Open != nil {
		open = cimgui.CollapsingHeader_BoolPtr(h.Label, h.Open, h.flags)
	} else {
		open = cimgui.CollapsingHeader_TreeNodeFlags(h.Label, h.flags)
	}
	if open {
		for _, w := range h.Widgets {
			if w != nil {
				w.Display()
			}
		}
	}
}
