package imgui

import "rodusek.dev/pkg/imgui/internal/cimgui"

// Separator draws a horizontal line.
func Separator() { cimgui.Separator() }

// SameLine keeps the next widget on the current line using default spacing.
func SameLine() { cimgui.SameLine(0, 0) }

// SameLineV keeps the next widget on the current line at the given offset from
// the line start, with the given spacing. Pass zeros for defaults.
func SameLineV(offsetFromStartX, spacing float32) {
	cimgui.SameLine(offsetFromStartX, spacing)
}

// NewLine moves to the start of the next line.
func NewLine() { cimgui.NewLine() }

// Spacing adds vertical space between widgets.
func Spacing() { cimgui.Spacing() }

// Dummy reserves an empty rectangle of the given size.
func Dummy(size Vec2) { cimgui.Dummy(size.X, size.Y) }

// Indent increases the horizontal indent by the default amount.
func Indent() { cimgui.Indent(0) }

// IndentV increases the horizontal indent by width.
func IndentV(width float32) { cimgui.Indent(width) }

// Unindent decreases the horizontal indent by the default amount.
func Unindent() { cimgui.Unindent(0) }

// UnindentV decreases the horizontal indent by width.
func UnindentV(width float32) { cimgui.Unindent(width) }
