package imgui

import "rodusek.dev/pkg/imgui/internal/cimgui"

// Begin opens a window. Always pair it with [End], regardless of the return
// value. Returns false when the window is collapsed or clipped.
func Begin(name string) bool {
	return cimgui.Begin(name, nil, 0)
}

// BeginV opens a window with a close button (when open is non-nil; the user
// closing the window sets *open to false) and the given flags.
func BeginV(name string, open *bool, flags WindowFlags) bool {
	return cimgui.Begin(name, open, int(flags))
}

// End closes the window opened by the matching [Begin] or [BeginV].
func End() { cimgui.End() }

// BeginChild opens a child region filling the available space. Pair it with
// [EndChild].
func BeginChild(id string) bool {
	return cimgui.BeginChild(id, 0, 0, 0, 0)
}

// BeginChildV opens a child region with an explicit size and flags. A zero size
// component fills the available space on that axis.
func BeginChildV(id string, size Vec2, childFlags ChildFlags, windowFlags WindowFlags) bool {
	return cimgui.BeginChild(id, size.X, size.Y, int(childFlags), int(windowFlags))
}

// EndChild closes the child region opened by [BeginChild] or [BeginChildV].
func EndChild() { cimgui.EndChild() }

// SetNextWindowSize sets the size of the next window unconditionally.
func SetNextWindowSize(size Vec2) {
	cimgui.SetNextWindowSize(size.X, size.Y, 0)
}

// SetNextWindowSizeV sets the size of the next window subject to cond.
func SetNextWindowSizeV(size Vec2, cond Cond) {
	cimgui.SetNextWindowSize(size.X, size.Y, int(cond))
}

// SetNextWindowPos sets the position of the next window unconditionally.
func SetNextWindowPos(pos Vec2) {
	cimgui.SetNextWindowPos(pos.X, pos.Y, 0, 0, 0)
}

// SetNextWindowPosV sets the position of the next window subject to cond, with
// pivot selecting the alignment point ((0,0) top-left, (0.5,0.5) centered).
func SetNextWindowPosV(pos Vec2, cond Cond, pivot Vec2) {
	cimgui.SetNextWindowPos(pos.X, pos.Y, int(cond), pivot.X, pivot.Y)
}
