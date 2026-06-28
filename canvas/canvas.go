package canvas

import (
	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// Drawer appends primitives to a draw list. For a [Canvas] its coordinates are
// local to the canvas region (origin at the top-left); for [Foreground] and
// [Background] they are absolute screen coordinates. A Drawer is only valid for
// the duration of the draw callback it is passed to.
type Drawer struct {
	list   cimgui.DrawList
	origin imgui.Vec2
	size   imgui.Vec2
}

// Origin returns the screen position the drawer's coordinates are relative to.
func (d *Drawer) Origin() imgui.Vec2 {
	return d.origin
}

// Avail returns the size of the region the drawer covers. It is zero for the
// overlay drawers, whose coordinates are absolute.
func (d *Drawer) Avail() imgui.Vec2 {
	return d.size
}

// at translates a drawer-local point into absolute screen coordinates.
func (d *Drawer) at(p imgui.Vec2) imgui.Vec2 {
	return imgui.Vec2{X: d.origin.X + p.X, Y: d.origin.Y + p.Y}
}

// AddLine draws a line from p1 to p2.
func (d *Drawer) AddLine(p1, p2 imgui.Vec2, col imgui.Color, thickness float32) {
	cimgui.DrawListAddLine(d.list, d.at(p1), d.at(p2), col.U32(), thickness)
}

// AddRect draws a rectangle outline between min and max.
func (d *Drawer) AddRect(min, max imgui.Vec2, col imgui.Color, rounding, thickness float32) {
	cimgui.DrawListAddRect(d.list, d.at(min), d.at(max), col.U32(), rounding, thickness, cimgui.DrawFlagsRoundCornersAll)
}

// AddRectFilled draws a filled rectangle between min and max.
func (d *Drawer) AddRectFilled(min, max imgui.Vec2, col imgui.Color, rounding float32) {
	cimgui.DrawListAddRectFilled(d.list, d.at(min), d.at(max), col.U32(), rounding, cimgui.DrawFlagsRoundCornersAll)
}

// AddCircle draws a circle outline. A zero segments auto-tessellates.
func (d *Drawer) AddCircle(center imgui.Vec2, radius float32, col imgui.Color, segments int32, thickness float32) {
	cimgui.DrawListAddCircle(d.list, d.at(center), radius, col.U32(), segments, thickness)
}

// AddCircleFilled draws a filled circle. A zero segments auto-tessellates.
func (d *Drawer) AddCircleFilled(center imgui.Vec2, radius float32, col imgui.Color, segments int32) {
	cimgui.DrawListAddCircleFilled(d.list, d.at(center), radius, col.U32(), segments)
}

// AddTriangle draws a triangle outline.
func (d *Drawer) AddTriangle(p1, p2, p3 imgui.Vec2, col imgui.Color, thickness float32) {
	cimgui.DrawListAddTriangle(d.list, d.at(p1), d.at(p2), d.at(p3), col.U32(), thickness)
}

// AddTriangleFilled draws a filled triangle.
func (d *Drawer) AddTriangleFilled(p1, p2, p3 imgui.Vec2, col imgui.Color) {
	cimgui.DrawListAddTriangleFilled(d.list, d.at(p1), d.at(p2), d.at(p3), col.U32())
}

// AddText draws text at pos.
func (d *Drawer) AddText(pos imgui.Vec2, col imgui.Color, text string) {
	cimgui.DrawListAddText(d.list, d.at(pos), col.U32(), text)
}

// AddPolyline draws a connected sequence of line segments, closing the loop when
// closed is set.
func (d *Drawer) AddPolyline(points []imgui.Vec2, col imgui.Color, closed bool, thickness float32) {
	flags := cimgui.DrawFlagsNone
	if closed {
		flags = cimgui.DrawFlagsClosed
	}
	cimgui.DrawListAddPolyline(d.list, d.translate(points), col.U32(), flags, thickness)
}

// AddConvexPolyFilled fills the convex polygon described by points.
func (d *Drawer) AddConvexPolyFilled(points []imgui.Vec2, col imgui.Color) {
	cimgui.DrawListAddConvexPolyFilled(d.list, d.translate(points), col.U32())
}

// translate offsets every point by the drawer origin.
func (d *Drawer) translate(points []imgui.Vec2) []imgui.Vec2 {
	out := make([]imgui.Vec2, len(points))
	for i, p := range points {
		out[i] = d.at(p)
	}
	return out
}

// Canvas reserves a region of the current window and draws into it. Draw receives
// a [Drawer] whose coordinates are local to the region. The region also behaves
// as a button, so hover and click state can be polled or handled via OnClick.
type Canvas struct {
	ID      string     // unique identifier within the window
	Size    imgui.Vec2 // a zero component fills the available space
	Draw    func(*Drawer)
	OnClick func()
	hovered bool
	clicked bool
}

var _ imgui.Widget = (*Canvas)(nil)

// New returns a canvas of the given size that draws with draw. id must be unique
// within the enclosing window.
func New(id string, size imgui.Vec2, draw func(*Drawer)) *Canvas {
	return &Canvas{ID: id, Size: size, Draw: draw}
}

// Display reserves the region and runs the draw callback.
func (c *Canvas) Display() {
	origin := cimgui.GetCursorScreenPos()
	size := c.resolveSize()
	c.clicked = cimgui.InvisibleButton(c.ID, size, cimgui.ButtonFlagsMouseButtonLeft)
	c.hovered = cimgui.IsItemHovered(cimgui.HoveredFlagsNone)
	if c.Draw != nil {
		c.Draw(&Drawer{list: cimgui.GetWindowDrawList(), origin: origin, size: size})
	}
	if c.clicked && c.OnClick != nil {
		c.OnClick()
	}
}

// resolveSize fills any zero component of Size with the available space.
func (c *Canvas) resolveSize() imgui.Vec2 {
	size := c.Size
	if size.X > 0 && size.Y > 0 {
		return size
	}
	avail := cimgui.GetContentRegionAvail()
	if size.X <= 0 {
		size.X = avail.X
	}
	if size.Y <= 0 {
		size.Y = avail.Y
	}
	return size
}

// Hovered reports whether the canvas was hovered during the last [Canvas.Display].
func (c *Canvas) Hovered() bool {
	return c.hovered
}

// Clicked reports whether the canvas was clicked during the last [Canvas.Display].
func (c *Canvas) Clicked() bool {
	return c.clicked
}

// Overlay draws to a per-viewport draw list that is rendered in front of (or
// behind) every window. Its [Drawer] uses absolute screen coordinates.
type Overlay struct {
	Draw       func(*Drawer)
	background bool
}

var _ imgui.Widget = (*Overlay)(nil)

// Foreground returns an overlay drawn in front of every window.
func Foreground(draw func(*Drawer)) *Overlay {
	return &Overlay{Draw: draw}
}

// Background returns an overlay drawn behind every window.
func Background(draw func(*Drawer)) *Overlay {
	return &Overlay{Draw: draw, background: true}
}

// Display runs the overlay draw callback.
func (o *Overlay) Display() {
	if o.Draw == nil {
		return
	}
	list := cimgui.GetForegroundDrawList()
	if o.background {
		list = cimgui.GetBackgroundDrawList()
	}
	o.Draw(&Drawer{list: list})
}
