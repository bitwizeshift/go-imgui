package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// DrawList is an opaque handle to an ImDrawList: the per-window or per-viewport
// list of draw commands the primitives below append to. Obtain one with
// [GetWindowDrawList], [GetForegroundDrawList] or [GetBackgroundDrawList].
type DrawList = cimgui.DrawList

// DrawOptions are the optional inputs to the draw-list primitives that take draw
// flags ([DrawListAddRect], [DrawListAddRectFilled], [DrawListAddPolyline]). A
// nil *DrawOptions uses Dear ImGui's defaults; each field maps to an ImDrawFlags_
// bit.
type DrawOptions struct {
	Closed           bool // ImDrawFlags_Closed
	RoundCornersNone bool // ImDrawFlags_RoundCornersNone
	RoundCornersAll  bool // ImDrawFlags_RoundCornersAll
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *DrawOptions) flags() cimgui.DrawFlags {
	if o == nil {
		return cimgui.DrawFlagsNone
	}
	var f cimgui.DrawFlags
	if o.Closed {
		f |= cimgui.DrawFlagsClosed
	}
	if o.RoundCornersNone {
		f |= cimgui.DrawFlagsRoundCornersNone
	}
	if o.RoundCornersAll {
		f |= cimgui.DrawFlagsRoundCornersAll
	}
	return f
}

// GetWindowDrawList returns the draw list of the current window. It models
// ImGui::GetWindowDrawList.
func GetWindowDrawList() DrawList {
	return cimgui.GetWindowDrawList()
}

// GetForegroundDrawList returns the draw list rendered in front of every window.
// It models ImGui::GetForegroundDrawList.
func GetForegroundDrawList() DrawList {
	return cimgui.GetForegroundDrawList()
}

// GetBackgroundDrawList returns the draw list rendered behind every window. It
// models ImGui::GetBackgroundDrawList.
func GetBackgroundDrawList() DrawList {
	return cimgui.GetBackgroundDrawList()
}

// DrawListAddLine draws a line from p1 to p2. It models ImDrawList::AddLine.
func DrawListAddLine(d DrawList, p1, p2 Vec2, col U32, thickness float32) {
	cimgui.DrawListAddLine(d, p1, p2, col, thickness)
}

// DrawListAddRect draws a rectangle outline between pMin and pMax. It models
// ImDrawList::AddRect.
func DrawListAddRect(d DrawList, pMin, pMax Vec2, col U32, rounding, thickness float32, opts *DrawOptions) {
	cimgui.DrawListAddRect(d, pMin, pMax, col, rounding, thickness, opts.flags())
}

// DrawListAddRectFilled draws a filled rectangle between pMin and pMax. It models
// ImDrawList::AddRectFilled.
func DrawListAddRectFilled(d DrawList, pMin, pMax Vec2, col U32, rounding float32, opts *DrawOptions) {
	cimgui.DrawListAddRectFilled(d, pMin, pMax, col, rounding, opts.flags())
}

// DrawListAddCircle draws a circle outline. A zero numSegments auto-tessellates.
// It models ImDrawList::AddCircle.
func DrawListAddCircle(d DrawList, center Vec2, radius float32, col U32, numSegments int32, thickness float32) {
	cimgui.DrawListAddCircle(d, center, radius, col, numSegments, thickness)
}

// DrawListAddCircleFilled draws a filled circle. It models
// ImDrawList::AddCircleFilled.
func DrawListAddCircleFilled(d DrawList, center Vec2, radius float32, col U32, numSegments int32) {
	cimgui.DrawListAddCircleFilled(d, center, radius, col, numSegments)
}

// DrawListAddTriangle draws a triangle outline. It models
// ImDrawList::AddTriangle.
func DrawListAddTriangle(d DrawList, p1, p2, p3 Vec2, col U32, thickness float32) {
	cimgui.DrawListAddTriangle(d, p1, p2, p3, col, thickness)
}

// DrawListAddTriangleFilled draws a filled triangle. It models
// ImDrawList::AddTriangleFilled.
func DrawListAddTriangleFilled(d DrawList, p1, p2, p3 Vec2, col U32) {
	cimgui.DrawListAddTriangleFilled(d, p1, p2, p3, col)
}

// DrawListAddQuad draws a quad outline. It models ImDrawList::AddQuad.
func DrawListAddQuad(d DrawList, p1, p2, p3, p4 Vec2, col U32, thickness float32) {
	cimgui.DrawListAddQuad(d, p1, p2, p3, p4, col, thickness)
}

// DrawListAddQuadFilled draws a filled quad. It models ImDrawList::AddQuadFilled.
func DrawListAddQuadFilled(d DrawList, p1, p2, p3, p4 Vec2, col U32) {
	cimgui.DrawListAddQuadFilled(d, p1, p2, p3, p4, col)
}

// DrawListAddText draws text at pos. It models ImDrawList::AddText.
func DrawListAddText(d DrawList, pos Vec2, col U32, text string) {
	cimgui.DrawListAddText(d, pos, col, text)
}

// DrawListAddBezierCubic draws a cubic Bézier curve through the four control
// points. A zero numSegments auto-tessellates. It models
// ImDrawList::AddBezierCubic.
func DrawListAddBezierCubic(d DrawList, p1, p2, p3, p4 Vec2, col U32, thickness float32, numSegments int32) {
	cimgui.DrawListAddBezierCubic(d, p1, p2, p3, p4, col, thickness, numSegments)
}

// DrawListAddPolyline draws a connected sequence of line segments. It models
// ImDrawList::AddPolyline.
func DrawListAddPolyline(d DrawList, points []Vec2, col U32, thickness float32, opts *DrawOptions) {
	cimgui.DrawListAddPolyline(d, points, col, opts.flags(), thickness)
}

// DrawListAddConvexPolyFilled fills the convex polygon described by points. It
// models ImDrawList::AddConvexPolyFilled.
func DrawListAddConvexPolyFilled(d DrawList, points []Vec2, col U32) {
	cimgui.DrawListAddConvexPolyFilled(d, points, col)
}

// DrawListPushClipRect restricts subsequent drawing to the given rectangle. It
// models ImDrawList::PushClipRect.
func DrawListPushClipRect(d DrawList, min, max Vec2, intersectWithCurrent bool) {
	cimgui.DrawListPushClipRect(d, min, max, intersectWithCurrent)
}

// DrawListPopClipRect undoes the most recent [DrawListPushClipRect]. It models
// ImDrawList::PopClipRect.
func DrawListPopClipRect(d DrawList) {
	cimgui.DrawListPopClipRect(d)
}
