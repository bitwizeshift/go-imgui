package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// WindowOptions are the optional inputs to [Window]. A nil *WindowOptions uses
// Dear ImGui's defaults; each field maps to an ImGuiWindowFlags_ bit.
type WindowOptions struct {
	NoTitleBar                bool // ImGuiWindowFlags_NoTitleBar
	NoResize                  bool // ImGuiWindowFlags_NoResize
	NoMove                    bool // ImGuiWindowFlags_NoMove
	NoScrollbar               bool // ImGuiWindowFlags_NoScrollbar
	NoScrollWithMouse         bool // ImGuiWindowFlags_NoScrollWithMouse
	NoCollapse                bool // ImGuiWindowFlags_NoCollapse
	AlwaysAutoResize          bool // ImGuiWindowFlags_AlwaysAutoResize
	NoBackground              bool // ImGuiWindowFlags_NoBackground
	NoSavedSettings           bool // ImGuiWindowFlags_NoSavedSettings
	NoMouseInputs             bool // ImGuiWindowFlags_NoMouseInputs
	MenuBar                   bool // ImGuiWindowFlags_MenuBar
	HorizontalScrollbar       bool // ImGuiWindowFlags_HorizontalScrollbar
	NoFocusOnAppearing        bool // ImGuiWindowFlags_NoFocusOnAppearing
	NoBringToFrontOnFocus     bool // ImGuiWindowFlags_NoBringToFrontOnFocus
	AlwaysVerticalScrollbar   bool // ImGuiWindowFlags_AlwaysVerticalScrollbar
	AlwaysHorizontalScrollbar bool // ImGuiWindowFlags_AlwaysHorizontalScrollbar
	NoNavInputs               bool // ImGuiWindowFlags_NoNavInputs
	NoNavFocus                bool // ImGuiWindowFlags_NoNavFocus
	UnsavedDocument           bool // ImGuiWindowFlags_UnsavedDocument
	NoNav                     bool // ImGuiWindowFlags_NoNav (composite)
	NoDecoration              bool // ImGuiWindowFlags_NoDecoration (composite)
	NoInputs                  bool // ImGuiWindowFlags_NoInputs (composite)
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *WindowOptions) flags() cimgui.WindowFlags {
	if o == nil {
		return cimgui.WindowFlagsNone
	}
	var f cimgui.WindowFlags
	if o.NoTitleBar {
		f |= cimgui.WindowFlagsNoTitleBar
	}
	if o.NoResize {
		f |= cimgui.WindowFlagsNoResize
	}
	if o.NoMove {
		f |= cimgui.WindowFlagsNoMove
	}
	if o.NoScrollbar {
		f |= cimgui.WindowFlagsNoScrollbar
	}
	if o.NoScrollWithMouse {
		f |= cimgui.WindowFlagsNoScrollWithMouse
	}
	if o.NoCollapse {
		f |= cimgui.WindowFlagsNoCollapse
	}
	if o.AlwaysAutoResize {
		f |= cimgui.WindowFlagsAlwaysAutoResize
	}
	if o.NoBackground {
		f |= cimgui.WindowFlagsNoBackground
	}
	if o.NoSavedSettings {
		f |= cimgui.WindowFlagsNoSavedSettings
	}
	if o.NoMouseInputs {
		f |= cimgui.WindowFlagsNoMouseInputs
	}
	if o.MenuBar {
		f |= cimgui.WindowFlagsMenuBar
	}
	if o.HorizontalScrollbar {
		f |= cimgui.WindowFlagsHorizontalScrollbar
	}
	if o.NoFocusOnAppearing {
		f |= cimgui.WindowFlagsNoFocusOnAppearing
	}
	if o.NoBringToFrontOnFocus {
		f |= cimgui.WindowFlagsNoBringToFrontOnFocus
	}
	if o.AlwaysVerticalScrollbar {
		f |= cimgui.WindowFlagsAlwaysVerticalScrollbar
	}
	if o.AlwaysHorizontalScrollbar {
		f |= cimgui.WindowFlagsAlwaysHorizontalScrollbar
	}
	if o.NoNavInputs {
		f |= cimgui.WindowFlagsNoNavInputs
	}
	if o.NoNavFocus {
		f |= cimgui.WindowFlagsNoNavFocus
	}
	if o.UnsavedDocument {
		f |= cimgui.WindowFlagsUnsavedDocument
	}
	if o.NoNav {
		f |= cimgui.WindowFlagsNoNav
	}
	if o.NoDecoration {
		f |= cimgui.WindowFlagsNoDecoration
	}
	if o.NoInputs {
		f |= cimgui.WindowFlagsNoInputs
	}
	return f
}

// ChildOptions are the optional inputs to [Child]. A nil *ChildOptions uses Dear
// ImGui's defaults; each field maps to an ImGuiChildFlags_ bit.
type ChildOptions struct {
	Borders                bool // ImGuiChildFlags_Borders
	AlwaysUseWindowPadding bool // ImGuiChildFlags_AlwaysUseWindowPadding
	ResizeX                bool // ImGuiChildFlags_ResizeX
	ResizeY                bool // ImGuiChildFlags_ResizeY
	AutoResizeX            bool // ImGuiChildFlags_AutoResizeX
	AutoResizeY            bool // ImGuiChildFlags_AutoResizeY
	AlwaysAutoResize       bool // ImGuiChildFlags_AlwaysAutoResize
	FrameStyle             bool // ImGuiChildFlags_FrameStyle
	NavFlattened           bool // ImGuiChildFlags_NavFlattened
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *ChildOptions) flags() cimgui.ChildFlags {
	if o == nil {
		return cimgui.ChildFlagsNone
	}
	var f cimgui.ChildFlags
	if o.Borders {
		f |= cimgui.ChildFlagsBorders
	}
	if o.AlwaysUseWindowPadding {
		f |= cimgui.ChildFlagsAlwaysUseWindowPadding
	}
	if o.ResizeX {
		f |= cimgui.ChildFlagsResizeX
	}
	if o.ResizeY {
		f |= cimgui.ChildFlagsResizeY
	}
	if o.AutoResizeX {
		f |= cimgui.ChildFlagsAutoResizeX
	}
	if o.AutoResizeY {
		f |= cimgui.ChildFlagsAutoResizeY
	}
	if o.AlwaysAutoResize {
		f |= cimgui.ChildFlagsAlwaysAutoResize
	}
	if o.FrameStyle {
		f |= cimgui.ChildFlagsFrameStyle
	}
	if o.NavFlattened {
		f |= cimgui.ChildFlagsNavFlattened
	}
	return f
}

// Window begins a window and pushes it onto the window stack. It models
// ImGui::Begin. When open is non-nil a close button is shown and *open is
// updated; visible reports whether the window's contents should be drawn. The
// returned [EndFunc] (ImGui::End) must always be called, regardless of visible.
func Window(name string, open *bool, opts *WindowOptions) (visible bool, end EndFunc) {
	visible = cimgui.Begin(name, open, opts.flags())
	return visible, cimgui.End
}

// Child begins a child region identified by id (a string label or a precomputed
// uint32 id) and models ImGui::BeginChild. A zero size fills the available
// space. open reports whether the region is visible; the returned [EndFunc]
// (ImGui::EndChild) must always be called.
func Child[T ID](id T, size Vec2, child *ChildOptions, window *WindowOptions) (open bool, end EndFunc) {
	switch v := any(id).(type) {
	case string:
		open = cimgui.BeginChild_Str(v, size, child.flags(), window.flags())
	case uint32:
		open = cimgui.BeginChild_ID(v, size, child.flags(), window.flags())
	}
	return open, cimgui.EndChild
}

// SetNextWindowPos sets the position applied to the next window, with an
// optional pivot (0,0 top-left .. 1,1 bottom-right). It models
// ImGui::SetNextWindowPos.
func SetNextWindowPos(pos Vec2, cond Cond, pivot Vec2) {
	cimgui.SetNextWindowPos(pos, cond, pivot)
}

// SetNextWindowSize sets the size applied to the next window. It models
// ImGui::SetNextWindowSize.
func SetNextWindowSize(size Vec2, cond Cond) {
	cimgui.SetNextWindowSize(size, cond)
}

// SetNextWindowContentSize sets the content size applied to the next window. It
// models ImGui::SetNextWindowContentSize.
func SetNextWindowContentSize(size Vec2) {
	cimgui.SetNextWindowContentSize(size)
}

// SetNextWindowCollapsed sets the collapsed state applied to the next window. It
// models ImGui::SetNextWindowCollapsed.
func SetNextWindowCollapsed(collapsed bool, cond Cond) {
	cimgui.SetNextWindowCollapsed(collapsed, cond)
}

// SetNextWindowFocus focuses the next window. It models ImGui::SetNextWindowFocus.
func SetNextWindowFocus() {
	cimgui.SetNextWindowFocus()
}

// SetNextWindowBgAlpha overrides the background alpha of the next window. It
// models ImGui::SetNextWindowBgAlpha.
func SetNextWindowBgAlpha(alpha float32) {
	cimgui.SetNextWindowBgAlpha(alpha)
}
