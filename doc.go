// package imgui is an idiomatic Go wrapper around Dear ImGui.
//
// It is a thin, curated layer over the vendored cimgui C API (compiled from
// source via cgo in the internal cimgui package). Functions follow Dear ImGui's
// immediate-mode model: call them between [NewFrame] and [Render] each frame.
// Common calls have a short form plus a "V" (verbose) variant that exposes the
// full set of options, e.g. [Button] and [ButtonV], [Begin] and [BeginV].
//
// Most programs do not call [NewFrame]/[Render] or manage a window directly;
// the app subpackage drives the window and frame loop. See app.Run.
package imgui

//go:generate go run ./scripts/gencsources.go
