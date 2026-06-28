package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// DrawList is an opaque handle to an ImDrawList: the per-window or per-viewport
// list of draw commands the primitives below append to. Obtain one with
// [GetWindowDrawList], [GetForegroundDrawList] or [GetBackgroundDrawList].
type DrawList unsafe.Pointer

// DrawFlags configures draw-list primitives (corner rounding, polyline closing).
// Mirrors the public ImDrawFlags_.
type DrawFlags int32

const (
	DrawFlagsNone             DrawFlags = C.ImDrawFlags_None
	DrawFlagsClosed           DrawFlags = C.ImDrawFlags_Closed
	DrawFlagsRoundCornersNone DrawFlags = C.ImDrawFlags_RoundCornersNone
	DrawFlagsRoundCornersAll  DrawFlags = C.ImDrawFlags_RoundCornersAll
)

// GetWindowDrawList returns the draw list of the current window.
func GetWindowDrawList() DrawList {
	return DrawList(C.igGetWindowDrawList())
}

// GetForegroundDrawList returns the draw list rendered in front of every window.
func GetForegroundDrawList() DrawList {
	return DrawList(C.igGetForegroundDrawList_Nil())
}

// GetBackgroundDrawList returns the draw list rendered behind every window.
func GetBackgroundDrawList() DrawList {
	return DrawList(C.igGetBackgroundDrawList_Nil())
}

// clist converts a [DrawList] handle to the underlying C pointer.
func clist(d DrawList) *C.ImDrawList {
	return (*C.ImDrawList)(unsafe.Pointer(d))
}

// DrawListAddLine draws a line from p1 to p2.
func DrawListAddLine(d DrawList, p1, p2 Vec2, col U32, thickness float32) {
	C.ImDrawList_AddLine(clist(d), p1.c(), p2.c(), C.ImU32(col), C.float(thickness))
}

// DrawListAddRect draws a rectangle outline between pMin and pMax.
func DrawListAddRect(d DrawList, pMin, pMax Vec2, col U32, rounding, thickness float32, flags DrawFlags) {
	C.ImDrawList_AddRect(clist(d), pMin.c(), pMax.c(), C.ImU32(col), C.float(rounding), C.float(thickness), C.ImDrawFlags(flags))
}

// DrawListAddRectFilled draws a filled rectangle between pMin and pMax.
func DrawListAddRectFilled(d DrawList, pMin, pMax Vec2, col U32, rounding float32, flags DrawFlags) {
	C.ImDrawList_AddRectFilled(clist(d), pMin.c(), pMax.c(), C.ImU32(col), C.float(rounding), C.ImDrawFlags(flags))
}

// DrawListAddCircle draws a circle outline.
func DrawListAddCircle(d DrawList, center Vec2, radius float32, col U32, numSegments int32, thickness float32) {
	C.ImDrawList_AddCircle(clist(d), center.c(), C.float(radius), C.ImU32(col), C.int(numSegments), C.float(thickness))
}

// DrawListAddCircleFilled draws a filled circle.
func DrawListAddCircleFilled(d DrawList, center Vec2, radius float32, col U32, numSegments int32) {
	C.ImDrawList_AddCircleFilled(clist(d), center.c(), C.float(radius), C.ImU32(col), C.int(numSegments))
}

// DrawListAddTriangle draws a triangle outline.
func DrawListAddTriangle(d DrawList, p1, p2, p3 Vec2, col U32, thickness float32) {
	C.ImDrawList_AddTriangle(clist(d), p1.c(), p2.c(), p3.c(), C.ImU32(col), C.float(thickness))
}

// DrawListAddTriangleFilled draws a filled triangle.
func DrawListAddTriangleFilled(d DrawList, p1, p2, p3 Vec2, col U32) {
	C.ImDrawList_AddTriangleFilled(clist(d), p1.c(), p2.c(), p3.c(), C.ImU32(col))
}

// DrawListAddQuad draws a quad outline.
func DrawListAddQuad(d DrawList, p1, p2, p3, p4 Vec2, col U32, thickness float32) {
	C.ImDrawList_AddQuad(clist(d), p1.c(), p2.c(), p3.c(), p4.c(), C.ImU32(col), C.float(thickness))
}

// DrawListAddQuadFilled draws a filled quad.
func DrawListAddQuadFilled(d DrawList, p1, p2, p3, p4 Vec2, col U32) {
	C.ImDrawList_AddQuadFilled(clist(d), p1.c(), p2.c(), p3.c(), p4.c(), C.ImU32(col))
}

// DrawListAddText draws text at pos.
func DrawListAddText(d DrawList, pos Vec2, col U32, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	end := (*C.char)(unsafe.Add(unsafe.Pointer(ctext), len(text)))
	C.ImDrawList_AddText_Vec2(clist(d), pos.c(), C.ImU32(col), ctext, end)
}

// DrawListAddBezierCubic draws a cubic Bézier curve through the four control
// points. A zero numSegments auto-tessellates.
func DrawListAddBezierCubic(d DrawList, p1, p2, p3, p4 Vec2, col U32, thickness float32, numSegments int32) {
	C.ImDrawList_AddBezierCubic(clist(d), p1.c(), p2.c(), p3.c(), p4.c(), C.ImU32(col), C.float(thickness), C.int(numSegments))
}

// pointsArray builds a C array of ImVec2_c from points. The returned free
// releases it; call it after the C call returns.
func pointsArray(points []Vec2) (arr *C.ImVec2_c, free func()) {
	if len(points) == 0 {
		return nil, func() {}
	}
	n := len(points)
	block := C.malloc(C.size_t(uintptr(n) * unsafe.Sizeof(C.ImVec2_c{})))
	view := unsafe.Slice((*C.ImVec2_c)(block), n)
	for i, p := range points {
		view[i] = p.c()
	}
	return (*C.ImVec2_c)(block), func() { C.free(block) }
}

// DrawListAddPolyline draws a connected sequence of line segments.
func DrawListAddPolyline(d DrawList, points []Vec2, col U32, flags DrawFlags, thickness float32) {
	arr, free := pointsArray(points)
	defer free()
	C.ImDrawList_AddPolyline(clist(d), arr, C.int(len(points)), C.ImU32(col), C.float(thickness), C.ImDrawFlags(flags))
}

// DrawListAddConvexPolyFilled fills the convex polygon described by points.
func DrawListAddConvexPolyFilled(d DrawList, points []Vec2, col U32) {
	arr, free := pointsArray(points)
	defer free()
	C.ImDrawList_AddConvexPolyFilled(clist(d), arr, C.int(len(points)), C.ImU32(col))
}

// DrawListPushClipRect restricts subsequent drawing to the given rectangle.
func DrawListPushClipRect(d DrawList, min, max Vec2, intersectWithCurrent bool) {
	C.ImDrawList_PushClipRect(clist(d), min.c(), max.c(), C.bool(intersectWithCurrent))
}

// DrawListPopClipRect undoes the most recent [DrawListPushClipRect].
func DrawListPopClipRect(d DrawList) {
	C.ImDrawList_PopClipRect(clist(d))
}
