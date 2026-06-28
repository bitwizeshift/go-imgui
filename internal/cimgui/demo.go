package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

// ShowDemoWindow displays the Dear ImGui demo window. pOpen may be nil.
func ShowDemoWindow(pOpen *bool) {
	withBoolPtr(pOpen, func(p *C.bool) { C.igShowDemoWindow(p) })
}

// ShowMetricsWindow displays the metrics/debug window. pOpen may be nil.
func ShowMetricsWindow(pOpen *bool) {
	withBoolPtr(pOpen, func(p *C.bool) { C.igShowMetricsWindow(p) })
}

// ShowAboutWindow displays the about window. pOpen may be nil.
func ShowAboutWindow(pOpen *bool) {
	withBoolPtr(pOpen, func(p *C.bool) { C.igShowAboutWindow(p) })
}
