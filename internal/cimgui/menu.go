package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// BeginMenuBar appends to the menu bar of the current window (requires
// WindowFlagsMenuBar). Call [EndMenuBar] only if it returns true.
func BeginMenuBar() bool { return bool(C.igBeginMenuBar()) }

// EndMenuBar closes the menu bar opened by [BeginMenuBar].
func EndMenuBar() { C.igEndMenuBar() }

// BeginMainMenuBar opens a full-screen menu bar at the top of the viewport. Call
// [EndMainMenuBar] only if it returns true.
func BeginMainMenuBar() bool { return bool(C.igBeginMainMenuBar()) }

// EndMainMenuBar closes the menu bar opened by [BeginMainMenuBar].
func EndMainMenuBar() { C.igEndMainMenuBar() }

// BeginMenu opens a sub-menu entry. Call [EndMenu] only if it returns true.
func BeginMenu(label string, enabled bool) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igBeginMenu(clabel, C.bool(enabled)))
}

// EndMenu closes the menu opened by [BeginMenu].
func EndMenu() { C.igEndMenu() }

// MenuItem_Bool draws a menu item and reports whether it was activated. shortcut
// may be empty.
func MenuItem_Bool(label, shortcut string, selected, enabled bool) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	var cshortcut *C.char
	if shortcut != "" {
		cshortcut = C.CString(shortcut)
		defer C.free(unsafe.Pointer(cshortcut))
	}
	return bool(C.igMenuItem_Bool(clabel, cshortcut, C.bool(selected), C.bool(enabled)))
}

// MenuItem_BoolPtr draws a menu item bound to pSelected (toggled on activation)
// and reports whether it was activated. shortcut may be empty.
func MenuItem_BoolPtr(label, shortcut string, pSelected *bool, enabled bool) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	var cshortcut *C.char
	if shortcut != "" {
		cshortcut = C.CString(shortcut)
		defer C.free(unsafe.Pointer(cshortcut))
	}
	var ret C.bool
	withBoolPtr(pSelected, func(p *C.bool) {
		ret = C.igMenuItem_BoolPtr(clabel, cshortcut, p, C.bool(enabled))
	})
	return bool(ret)
}
