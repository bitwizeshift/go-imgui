package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// WindowFlags configures a window opened with [Begin]. Mirrors ImGuiWindowFlags_.
type WindowFlags int32

const (
	WindowFlagsNone                      WindowFlags = C.ImGuiWindowFlags_None
	WindowFlagsNoTitleBar                WindowFlags = C.ImGuiWindowFlags_NoTitleBar
	WindowFlagsNoResize                  WindowFlags = C.ImGuiWindowFlags_NoResize
	WindowFlagsNoMove                    WindowFlags = C.ImGuiWindowFlags_NoMove
	WindowFlagsNoScrollbar               WindowFlags = C.ImGuiWindowFlags_NoScrollbar
	WindowFlagsNoScrollWithMouse         WindowFlags = C.ImGuiWindowFlags_NoScrollWithMouse
	WindowFlagsNoCollapse                WindowFlags = C.ImGuiWindowFlags_NoCollapse
	WindowFlagsAlwaysAutoResize          WindowFlags = C.ImGuiWindowFlags_AlwaysAutoResize
	WindowFlagsNoBackground              WindowFlags = C.ImGuiWindowFlags_NoBackground
	WindowFlagsNoSavedSettings           WindowFlags = C.ImGuiWindowFlags_NoSavedSettings
	WindowFlagsNoMouseInputs             WindowFlags = C.ImGuiWindowFlags_NoMouseInputs
	WindowFlagsMenuBar                   WindowFlags = C.ImGuiWindowFlags_MenuBar
	WindowFlagsHorizontalScrollbar       WindowFlags = C.ImGuiWindowFlags_HorizontalScrollbar
	WindowFlagsNoFocusOnAppearing        WindowFlags = C.ImGuiWindowFlags_NoFocusOnAppearing
	WindowFlagsNoBringToFrontOnFocus     WindowFlags = C.ImGuiWindowFlags_NoBringToFrontOnFocus
	WindowFlagsAlwaysVerticalScrollbar   WindowFlags = C.ImGuiWindowFlags_AlwaysVerticalScrollbar
	WindowFlagsAlwaysHorizontalScrollbar WindowFlags = C.ImGuiWindowFlags_AlwaysHorizontalScrollbar
	WindowFlagsNoNavInputs               WindowFlags = C.ImGuiWindowFlags_NoNavInputs
	WindowFlagsNoNavFocus                WindowFlags = C.ImGuiWindowFlags_NoNavFocus
	WindowFlagsUnsavedDocument           WindowFlags = C.ImGuiWindowFlags_UnsavedDocument
	WindowFlagsNoNav                     WindowFlags = C.ImGuiWindowFlags_NoNav
	WindowFlagsNoDecoration              WindowFlags = C.ImGuiWindowFlags_NoDecoration
	WindowFlagsNoInputs                  WindowFlags = C.ImGuiWindowFlags_NoInputs
)

// ChildFlags configures a child region opened with [BeginChild_Str]. Mirrors ImGuiChildFlags_.
type ChildFlags int32

const (
	ChildFlagsNone                   ChildFlags = C.ImGuiChildFlags_None
	ChildFlagsBorders                ChildFlags = C.ImGuiChildFlags_Borders
	ChildFlagsAlwaysUseWindowPadding ChildFlags = C.ImGuiChildFlags_AlwaysUseWindowPadding
	ChildFlagsResizeX                ChildFlags = C.ImGuiChildFlags_ResizeX
	ChildFlagsResizeY                ChildFlags = C.ImGuiChildFlags_ResizeY
	ChildFlagsAutoResizeX            ChildFlags = C.ImGuiChildFlags_AutoResizeX
	ChildFlagsAutoResizeY            ChildFlags = C.ImGuiChildFlags_AutoResizeY
	ChildFlagsAlwaysAutoResize       ChildFlags = C.ImGuiChildFlags_AlwaysAutoResize
	ChildFlagsFrameStyle             ChildFlags = C.ImGuiChildFlags_FrameStyle
	ChildFlagsNavFlattened           ChildFlags = C.ImGuiChildFlags_NavFlattened
)

// Cond selects when a state-setting call applies. Mirrors ImGuiCond_.
type Cond int32

const (
	CondNone         Cond = C.ImGuiCond_None
	CondAlways       Cond = C.ImGuiCond_Always
	CondOnce         Cond = C.ImGuiCond_Once
	CondFirstUseEver Cond = C.ImGuiCond_FirstUseEver
	CondAppearing    Cond = C.ImGuiCond_Appearing
)

// Begin opens a window. pOpen may be nil; when non-nil it is updated with the
// window's open state. Returns whether the window's contents are visible.
func Begin(name string, pOpen *bool, flags WindowFlags) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var ret C.bool
	withBoolPtr(pOpen, func(p *C.bool) {
		ret = C.igBegin(cname, p, C.ImGuiWindowFlags(flags))
	})
	return bool(ret)
}

// End closes the window opened by the matching [Begin].
func End() { C.igEnd() }

// BeginChild_Str opens a child region identified by a string. A zero size fills
// the available space. Call [EndChild] unconditionally afterwards.
func BeginChild_Str(strID string, size Vec2, childFlags ChildFlags, windowFlags WindowFlags) bool {
	cid := C.CString(strID)
	defer C.free(unsafe.Pointer(cid))
	return bool(C.igBeginChild_Str(cid, size.c(),
		C.ImGuiChildFlags(childFlags), C.ImGuiWindowFlags(windowFlags)))
}

// BeginChild_ID opens a child region identified by a numeric ID.
func BeginChild_ID(id uint32, size Vec2, childFlags ChildFlags, windowFlags WindowFlags) bool {
	return bool(C.igBeginChild_ID(C.ImGuiID(id), size.c(),
		C.ImGuiChildFlags(childFlags), C.ImGuiWindowFlags(windowFlags)))
}

// EndChild closes the child region opened by [BeginChild_Str] or [BeginChild_ID].
func EndChild() { C.igEndChild() }

// SetNextWindowPos sets the position of the next window, with an optional pivot.
func SetNextWindowPos(pos Vec2, cond Cond, pivot Vec2) {
	C.igSetNextWindowPos(pos.c(), C.ImGuiCond(cond), pivot.c())
}

// SetNextWindowSize sets the size of the next window.
func SetNextWindowSize(size Vec2, cond Cond) {
	C.igSetNextWindowSize(size.c(), C.ImGuiCond(cond))
}

// SetNextWindowContentSize sets the content size of the next window.
func SetNextWindowContentSize(size Vec2) { C.igSetNextWindowContentSize(size.c()) }

// SetNextWindowCollapsed sets the collapsed state of the next window.
func SetNextWindowCollapsed(collapsed bool, cond Cond) {
	C.igSetNextWindowCollapsed(C.bool(collapsed), C.ImGuiCond(cond))
}

// SetNextWindowFocus focuses the next window.
func SetNextWindowFocus() { C.igSetNextWindowFocus() }

// SetNextWindowBgAlpha overrides the background alpha of the next window.
func SetNextWindowBgAlpha(alpha float32) { C.igSetNextWindowBgAlpha(C.float(alpha)) }
