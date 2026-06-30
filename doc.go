// Package imgui is an idiomatic Go wrapper around Dear ImGui.
//
// The intended use-case for this package is to leverage the "Widget" hierarchy,
// and to use optional Display types for everything. However, this package
// _also_ exposes an "immediate"-mode API that models the Begin/End API from
// ImGui near 1:1 -- except instead of "End" functions, the Begin function
// returns an [EndFunc] which can be called in a defer statement.
//
// ## Widgets
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
//
// ## Immediate-mode
//
// Alongside the retained [Widget] tree, this package also exposes a near 1:1
// immediate-mode mapping of the Dear ImGui C API for cases the widget tree does
// not yet model. These functions are typically driven from a [Custom] widget (or
// any code running between [NewFrame] and [Render]) and each documents the
// ImGui:: function it models. Three conventions adapt the C API to Go:
//
//   - Enum bitflags become an "Options" struct of bool fields, always passed by
//     pointer so a nil argument selects Dear ImGui's defaults (see for example
//     [WindowOptions] and [Window]). Optional scalar parameters such as a printf
//     format string are carried on the same struct.
//   - Single-choice enums (such as [Dir], [Cond] and [Col]) remain typed
//     constants.
//   - The begin/end scopes return an [EndFunc] to call (typically deferred)
//     instead of a separate End function; for example
//     "open, end := imgui.Window(...); defer end()". The Push/Pop primitives are
//     instead exposed as standalone pairs (such as [PushStyleColor] and
//     [PopStyleColor], or [TreePush] and [TreePop]).
//
// C++ overloads that differ only by a value type are unified with generics (for
// example [CheckboxFlags], the value/pointer overloads of [Selectable] and
// [MenuItem], and the id-overloaded [Child], whose id may be a string label or a
// precomputed uint32). Overloads that differ in arity or behaviour keep distinct
// names (for example [RadioButton] versus [RadioButtonGroup]).
package imgui

//go:generate go run ./scripts/gencsources.go
