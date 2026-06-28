package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

// Separator draws a horizontal line.
func Separator() { C.igSeparator() }

// SameLine continues the current line. Pass zeros for default spacing.
func SameLine(offsetFromStartX, spacing float32) {
	C.igSameLine(C.float(offsetFromStartX), C.float(spacing))
}

// NewLine moves to the next line.
func NewLine() { C.igNewLine() }

// Spacing adds vertical spacing.
func Spacing() { C.igSpacing() }

// Dummy adds an empty item of the given size.
func Dummy(size Vec2) { C.igDummy(size.c()) }

// Indent increases the indent. A zero width uses the default.
func Indent(indentW float32) { C.igIndent(C.float(indentW)) }

// Unindent decreases the indent. A zero width uses the default.
func Unindent(indentW float32) { C.igUnindent(C.float(indentW)) }

// BeginGroup starts a group; layout queries treat the group as one item until
// the matching [EndGroup].
func BeginGroup() { C.igBeginGroup() }

// EndGroup ends the group opened by [BeginGroup].
func EndGroup() { C.igEndGroup() }

// GetCursorScreenPos returns the cursor position in absolute screen coordinates,
// the origin used by the [DrawList] primitives.
func GetCursorScreenPos() Vec2 {
	return vec2(C.igGetCursorScreenPos())
}

// SetCursorScreenPos moves the cursor to pos in absolute screen coordinates.
func SetCursorScreenPos(pos Vec2) {
	C.igSetCursorScreenPos(pos.c())
}

// GetCursorPos returns the cursor position in window-local coordinates.
func GetCursorPos() Vec2 {
	return vec2(C.igGetCursorPos())
}

// GetContentRegionAvail returns the space remaining from the cursor to the edge
// of the current content region.
func GetContentRegionAvail() Vec2 {
	return vec2(C.igGetContentRegionAvail())
}
