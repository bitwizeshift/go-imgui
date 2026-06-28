package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// floatsPtr returns a C float pointer to the first element of v, or nil if empty.
func floatsPtr(v []float32) *C.float {
	if len(v) == 0 {
		return nil
	}
	return (*C.float)(&v[0])
}

// PlotLines_FloatPtr draws a line plot of values. A zero graphSize auto-fits;
// scaleMin/scaleMax may be set to the same value to auto-scale. overlayText may
// be empty. valuesOffset rotates the starting index; stride is the byte stride
// between samples (use 4 for a packed []float32).
func PlotLines_FloatPtr(label string, values []float32, valuesOffset int32, overlayText string, scaleMin, scaleMax float32, graphSize Vec2, stride int32) {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	var coverlay *C.char
	if overlayText != "" {
		coverlay = C.CString(overlayText)
		defer C.free(unsafe.Pointer(coverlay))
	}
	C.igPlotLines_FloatPtr(clabel, floatsPtr(values), C.int(len(values)), C.int(valuesOffset),
		coverlay, C.float(scaleMin), C.float(scaleMax), graphSize.c(), C.int(stride))
}

// PlotHistogram_FloatPtr draws a histogram of values. See [PlotLines_FloatPtr]
// for the meaning of the remaining parameters.
func PlotHistogram_FloatPtr(label string, values []float32, valuesOffset int32, overlayText string, scaleMin, scaleMax float32, graphSize Vec2, stride int32) {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	var coverlay *C.char
	if overlayText != "" {
		coverlay = C.CString(overlayText)
		defer C.free(unsafe.Pointer(coverlay))
	}
	C.igPlotHistogram_FloatPtr(clabel, floatsPtr(values), C.int(len(values)), C.int(valuesOffset),
		coverlay, C.float(scaleMin), C.float(scaleMax), graphSize.c(), C.int(stride))
}
