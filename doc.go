// Package imgui is an idiomatic Go wrapper around Dear ImGui.
//
// The UI is a retained tree of [Widget] values rebuilt each frame inside the
// render callback (see app.Run). Containers hold child widgets and own their
// begin/end pairing, so ordering can never be wrong; build widgets as plain
// structs (composite literals) or with the per-package constructors, set their
// fields, and add them to a container:
//
//	w := window.New("Hello")
//	b := button.New("Click")
//	b.OnClick = func() { count++ }
//	w.AddWidget(text.Labelf("count = %d", count), b)
//	w.Display()
//
// Widgets live in concept subpackages (text, button, input, color, layout,
// window, tree, tab, table, menu, popup, combo, tooltip, plot, texture, debug).
// This root package holds the [Widget] interface, the shared value types
// ([Vec2], [Vec4], [Color]), the context/frame lifecycle, and the [Custom]
// escape hatch. The thin C bindings live in the internal cimgui package and are
// not part of the public API.
package imgui

//go:generate go run ./scripts/gencsources.go
