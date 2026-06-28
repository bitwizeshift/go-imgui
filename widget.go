package imgui

// Widget is anything that can draw itself for the current frame. The high-level
// API is a retained tree of Widgets rebuilt each frame: containers hold child
// Widgets and own their begin/end pairing, so widgets can never be issued in the
// wrong order.
type Widget interface {
	// Display draws the widget for the current frame. It must be called between
	// [NewFrame] and [Render] (app.Run handles that).
	Display()
}

// Display draws each widget in order. It is a convenience for the top level of a
// frame, e.g. imgui.Display(window1, window2).
func Display(widgets ...Widget) {
	for _, w := range widgets {
		if w != nil {
			w.Display()
		}
	}
}

// CustomWidget runs an arbitrary function as a widget. It is the escape hatch for
// behaviour not yet modelled by a dedicated widget; Func may call the lower-level
// API directly.
type CustomWidget struct {
	Func func()
}

// Custom returns a [CustomWidget] that runs fn when displayed.
func Custom(fn func()) *CustomWidget { return &CustomWidget{Func: fn} }

// Display runs the wrapped function.
func (c *CustomWidget) Display() {
	if c.Func != nil {
		c.Func()
	}
}
