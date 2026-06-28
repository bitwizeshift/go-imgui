package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// SliderFlags configures sliders and drags. Mirrors the public ImGuiSliderFlags_.
type SliderFlags int32

const (
	SliderFlagsNone            SliderFlags = C.ImGuiSliderFlags_None
	SliderFlagsLogarithmic     SliderFlags = C.ImGuiSliderFlags_Logarithmic
	SliderFlagsNoRoundToFormat SliderFlags = C.ImGuiSliderFlags_NoRoundToFormat
	SliderFlagsNoInput         SliderFlags = C.ImGuiSliderFlags_NoInput
	SliderFlagsWrapAround      SliderFlags = C.ImGuiSliderFlags_WrapAround
	SliderFlagsClampOnInput    SliderFlags = C.ImGuiSliderFlags_ClampOnInput
	SliderFlagsClampZeroRange  SliderFlags = C.ImGuiSliderFlags_ClampZeroRange
	SliderFlagsNoSpeedTweaks   SliderFlags = C.ImGuiSliderFlags_NoSpeedTweaks
	SliderFlagsAlwaysClamp     SliderFlags = C.ImGuiSliderFlags_AlwaysClamp
)

// ---- Sliders -------------------------------------------------------------

// SliderFloat draws a float slider bound to v and reports whether it changed.
func SliderFloat(label string, v *float32, vMin, vMax float32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderFloat(clabel, (*C.float)(v), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// SliderFloat2 draws a 2-component float slider bound to v.
func SliderFloat2(label string, v *[2]float32, vMin, vMax float32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderFloat2(clabel, (*C.float)(&v[0]), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// SliderFloat3 draws a 3-component float slider bound to v.
func SliderFloat3(label string, v *[3]float32, vMin, vMax float32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderFloat3(clabel, (*C.float)(&v[0]), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// SliderFloat4 draws a 4-component float slider bound to v.
func SliderFloat4(label string, v *[4]float32, vMin, vMax float32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderFloat4(clabel, (*C.float)(&v[0]), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// SliderAngle draws a slider editing v (radians) shown in degrees.
func SliderAngle(label string, vRad *float32, vDegreesMin, vDegreesMax float32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderAngle(clabel, (*C.float)(vRad), C.float(vDegreesMin), C.float(vDegreesMax), cformat, C.ImGuiSliderFlags(flags)))
}

// SliderInt draws an int slider bound to v and reports whether it changed.
func SliderInt(label string, v *int32, vMin, vMax int32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderInt(clabel, (*C.int)(v), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// SliderInt2 draws a 2-component int slider bound to v.
func SliderInt2(label string, v *[2]int32, vMin, vMax int32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderInt2(clabel, (*C.int)(&v[0]), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// SliderInt3 draws a 3-component int slider bound to v.
func SliderInt3(label string, v *[3]int32, vMin, vMax int32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderInt3(clabel, (*C.int)(&v[0]), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// SliderInt4 draws a 4-component int slider bound to v.
func SliderInt4(label string, v *[4]int32, vMin, vMax int32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderInt4(clabel, (*C.int)(&v[0]), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// SliderScalar draws a slider for an arbitrary data type. pData, pMin and pMax
// point to values of the given DataType.
func SliderScalar(label string, dataType DataType, pData, pMin, pMax unsafe.Pointer, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderScalar(clabel, C.ImGuiDataType(dataType), pData, pMin, pMax, cformat, C.ImGuiSliderFlags(flags)))
}

// SliderScalarN draws a slider editing components values of an arbitrary data
// type stored contiguously at pData.
func SliderScalarN(label string, dataType DataType, pData unsafe.Pointer, components int32, pMin, pMax unsafe.Pointer, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igSliderScalarN(clabel, C.ImGuiDataType(dataType), pData, C.int(components), pMin, pMax, cformat, C.ImGuiSliderFlags(flags)))
}

// VSliderFloat draws a vertical float slider of the given size.
func VSliderFloat(label string, size Vec2, v *float32, vMin, vMax float32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igVSliderFloat(clabel, size.c(), (*C.float)(v), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// VSliderInt draws a vertical int slider of the given size.
func VSliderInt(label string, size Vec2, v *int32, vMin, vMax int32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igVSliderInt(clabel, size.c(), (*C.int)(v), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// VSliderScalar draws a vertical slider for an arbitrary data type.
func VSliderScalar(label string, size Vec2, dataType DataType, pData, pMin, pMax unsafe.Pointer, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igVSliderScalar(clabel, size.c(), C.ImGuiDataType(dataType), pData, pMin, pMax, cformat, C.ImGuiSliderFlags(flags)))
}

// ---- Drags ---------------------------------------------------------------

// DragFloat draws a draggable float bound to v and reports whether it changed.
func DragFloat(label string, v *float32, speed, vMin, vMax float32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragFloat(clabel, (*C.float)(v), C.float(speed), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// DragFloat2 draws a draggable 2-component float bound to v.
func DragFloat2(label string, v *[2]float32, speed, vMin, vMax float32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragFloat2(clabel, (*C.float)(&v[0]), C.float(speed), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// DragFloat3 draws a draggable 3-component float bound to v.
func DragFloat3(label string, v *[3]float32, speed, vMin, vMax float32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragFloat3(clabel, (*C.float)(&v[0]), C.float(speed), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// DragFloat4 draws a draggable 4-component float bound to v.
func DragFloat4(label string, v *[4]float32, speed, vMin, vMax float32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragFloat4(clabel, (*C.float)(&v[0]), C.float(speed), C.float(vMin), C.float(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// DragInt draws a draggable int bound to v and reports whether it changed.
func DragInt(label string, v *int32, speed float32, vMin, vMax int32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragInt(clabel, (*C.int)(v), C.float(speed), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// DragInt2 draws a draggable 2-component int bound to v.
func DragInt2(label string, v *[2]int32, speed float32, vMin, vMax int32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragInt2(clabel, (*C.int)(&v[0]), C.float(speed), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// DragInt3 draws a draggable 3-component int bound to v.
func DragInt3(label string, v *[3]int32, speed float32, vMin, vMax int32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragInt3(clabel, (*C.int)(&v[0]), C.float(speed), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// DragInt4 draws a draggable 4-component int bound to v.
func DragInt4(label string, v *[4]int32, speed float32, vMin, vMax int32, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragInt4(clabel, (*C.int)(&v[0]), C.float(speed), C.int(vMin), C.int(vMax), cformat, C.ImGuiSliderFlags(flags)))
}

// DragFloatRange2 draws two draggable floats editing a [min,max] range.
func DragFloatRange2(label string, vCurrentMin, vCurrentMax *float32, speed, vMin, vMax float32, format, formatMax string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	cformatMax := C.CString(formatMax)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	defer C.free(unsafe.Pointer(cformatMax))
	return bool(C.igDragFloatRange2(clabel, (*C.float)(vCurrentMin), (*C.float)(vCurrentMax), C.float(speed), C.float(vMin), C.float(vMax), cformat, cformatMax, C.ImGuiSliderFlags(flags)))
}

// DragIntRange2 draws two draggable ints editing a [min,max] range.
func DragIntRange2(label string, vCurrentMin, vCurrentMax *int32, speed float32, vMin, vMax int32, format, formatMax string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	cformatMax := C.CString(formatMax)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	defer C.free(unsafe.Pointer(cformatMax))
	return bool(C.igDragIntRange2(clabel, (*C.int)(vCurrentMin), (*C.int)(vCurrentMax), C.float(speed), C.int(vMin), C.int(vMax), cformat, cformatMax, C.ImGuiSliderFlags(flags)))
}

// DragScalar draws a draggable widget for an arbitrary data type.
func DragScalar(label string, dataType DataType, pData unsafe.Pointer, speed float32, pMin, pMax unsafe.Pointer, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragScalar(clabel, C.ImGuiDataType(dataType), pData, C.float(speed), pMin, pMax, cformat, C.ImGuiSliderFlags(flags)))
}

// DragScalarN draws draggable widgets for components values of an arbitrary data
// type stored contiguously at pData.
func DragScalarN(label string, dataType DataType, pData unsafe.Pointer, components int32, speed float32, pMin, pMax unsafe.Pointer, format string, flags SliderFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igDragScalarN(clabel, C.ImGuiDataType(dataType), pData, C.int(components), C.float(speed), pMin, pMax, cformat, C.ImGuiSliderFlags(flags)))
}
