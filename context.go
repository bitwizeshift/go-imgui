package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// Context is a Dear ImGui context. Most programs let app.Run own the context and
// never create one directly.
type Context struct {
	ptr cimgui.Context
}

// CreateContext creates and activates a new context.
func CreateContext() Context {
	return Context{ptr: cimgui.CreateContext()}
}

// Destroy destroys the context.
func (c Context) Destroy() {
	cimgui.DestroyContext(c.ptr)
}

// NewFrame begins a new frame. Call once per frame before any widget calls.
func NewFrame() { cimgui.NewFrame() }

// EndFrame ends the frame; normally [Render] calls this for you.
func EndFrame() { cimgui.EndFrame() }

// Render finalizes the frame's draw data for the renderer backend.
func Render() { cimgui.Render() }

// StyleColorsDark applies the built-in dark style.
func StyleColorsDark() { cimgui.StyleColorsDark() }

// StyleColorsLight applies the built-in light style.
func StyleColorsLight() { cimgui.StyleColorsLight() }

// StyleColorsClassic applies the built-in classic style.
func StyleColorsClassic() { cimgui.StyleColorsClassic() }
