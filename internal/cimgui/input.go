package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// DataType identifies the element type for scalar widgets. Mirrors ImGuiDataType.
type DataType int32

const (
	DataTypeS8     DataType = C.ImGuiDataType_S8
	DataTypeU8     DataType = C.ImGuiDataType_U8
	DataTypeS16    DataType = C.ImGuiDataType_S16
	DataTypeU16    DataType = C.ImGuiDataType_U16
	DataTypeS32    DataType = C.ImGuiDataType_S32
	DataTypeU32    DataType = C.ImGuiDataType_U32
	DataTypeS64    DataType = C.ImGuiDataType_S64
	DataTypeU64    DataType = C.ImGuiDataType_U64
	DataTypeFloat  DataType = C.ImGuiDataType_Float
	DataTypeDouble DataType = C.ImGuiDataType_Double
	DataTypeBool   DataType = C.ImGuiDataType_Bool
)

// InputTextFlags configures input widgets. Mirrors the public ImGuiInputTextFlags_.
type InputTextFlags int32

const (
	InputTextFlagsNone                InputTextFlags = C.ImGuiInputTextFlags_None
	InputTextFlagsCharsDecimal        InputTextFlags = C.ImGuiInputTextFlags_CharsDecimal
	InputTextFlagsCharsHexadecimal    InputTextFlags = C.ImGuiInputTextFlags_CharsHexadecimal
	InputTextFlagsCharsScientific     InputTextFlags = C.ImGuiInputTextFlags_CharsScientific
	InputTextFlagsCharsUppercase      InputTextFlags = C.ImGuiInputTextFlags_CharsUppercase
	InputTextFlagsCharsNoBlank        InputTextFlags = C.ImGuiInputTextFlags_CharsNoBlank
	InputTextFlagsAllowTabInput       InputTextFlags = C.ImGuiInputTextFlags_AllowTabInput
	InputTextFlagsEnterReturnsTrue    InputTextFlags = C.ImGuiInputTextFlags_EnterReturnsTrue
	InputTextFlagsEscapeClearsAll     InputTextFlags = C.ImGuiInputTextFlags_EscapeClearsAll
	InputTextFlagsCtrlEnterForNewLine InputTextFlags = C.ImGuiInputTextFlags_CtrlEnterForNewLine
	InputTextFlagsReadOnly            InputTextFlags = C.ImGuiInputTextFlags_ReadOnly
	InputTextFlagsPassword            InputTextFlags = C.ImGuiInputTextFlags_Password
	InputTextFlagsAlwaysOverwrite     InputTextFlags = C.ImGuiInputTextFlags_AlwaysOverwrite
	InputTextFlagsAutoSelectAll       InputTextFlags = C.ImGuiInputTextFlags_AutoSelectAll
	InputTextFlagsParseEmptyRefVal    InputTextFlags = C.ImGuiInputTextFlags_ParseEmptyRefVal
	InputTextFlagsDisplayEmptyRefVal  InputTextFlags = C.ImGuiInputTextFlags_DisplayEmptyRefVal
	InputTextFlagsNoHorizontalScroll  InputTextFlags = C.ImGuiInputTextFlags_NoHorizontalScroll
	InputTextFlagsNoUndoRedo          InputTextFlags = C.ImGuiInputTextFlags_NoUndoRedo
	InputTextFlagsElideLeft           InputTextFlags = C.ImGuiInputTextFlags_ElideLeft

	InputTextFlagsCallbackCompletion InputTextFlags = C.ImGuiInputTextFlags_CallbackCompletion
	InputTextFlagsCallbackHistory    InputTextFlags = C.ImGuiInputTextFlags_CallbackHistory
	InputTextFlagsCallbackAlways     InputTextFlags = C.ImGuiInputTextFlags_CallbackAlways
	InputTextFlagsCallbackCharFilter InputTextFlags = C.ImGuiInputTextFlags_CallbackCharFilter
	InputTextFlagsCallbackResize     InputTextFlags = C.ImGuiInputTextFlags_CallbackResize
	InputTextFlagsCallbackEdit       InputTextFlags = C.ImGuiInputTextFlags_CallbackEdit
)

// bufPtr returns a C char pointer to the first byte of buf, or nil if empty.
func bufPtr(buf []byte) *C.char {
	if len(buf) == 0 {
		return nil
	}
	return (*C.char)(unsafe.Pointer(&buf[0]))
}

// InputText edits the NUL-terminated text held in buf (cap len(buf)) and reports
// whether it changed. Callback-based variants are not yet exposed.
func InputText(label string, buf []byte, flags InputTextFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igInputText(clabel, bufPtr(buf), C.size_t(len(buf)), C.ImGuiInputTextFlags(flags), nil, nil))
}

// InputTextMultiline edits buf in a multi-line box of the given size.
func InputTextMultiline(label string, buf []byte, size Vec2, flags InputTextFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igInputTextMultiline(clabel, bufPtr(buf), C.size_t(len(buf)), size.c(), C.ImGuiInputTextFlags(flags), nil, nil))
}

// InputTextWithHint edits buf, showing hint when empty.
func InputTextWithHint(label, hint string, buf []byte, flags InputTextFlags) bool {
	clabel, chint := C.CString(label), C.CString(hint)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(chint))
	return bool(C.igInputTextWithHint(clabel, chint, bufPtr(buf), C.size_t(len(buf)), C.ImGuiInputTextFlags(flags), nil, nil))
}

// InputFloat edits a float in a box with optional step buttons.
func InputFloat(label string, v *float32, step, stepFast float32, format string, flags InputTextFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputFloat(clabel, (*C.float)(v), C.float(step), C.float(stepFast), cformat, C.ImGuiInputTextFlags(flags)))
}

// InputFloat2 edits a 2-component float bound to v.
func InputFloat2(label string, v *[2]float32, format string, flags InputTextFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputFloat2(clabel, (*C.float)(&v[0]), cformat, C.ImGuiInputTextFlags(flags)))
}

// InputFloat3 edits a 3-component float bound to v.
func InputFloat3(label string, v *[3]float32, format string, flags InputTextFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputFloat3(clabel, (*C.float)(&v[0]), cformat, C.ImGuiInputTextFlags(flags)))
}

// InputFloat4 edits a 4-component float bound to v.
func InputFloat4(label string, v *[4]float32, format string, flags InputTextFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputFloat4(clabel, (*C.float)(&v[0]), cformat, C.ImGuiInputTextFlags(flags)))
}

// InputInt edits an int in a box with optional step buttons.
func InputInt(label string, v *int32, step, stepFast int32, flags InputTextFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igInputInt(clabel, (*C.int)(v), C.int(step), C.int(stepFast), C.ImGuiInputTextFlags(flags)))
}

// InputInt2 edits a 2-component int bound to v.
func InputInt2(label string, v *[2]int32, flags InputTextFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igInputInt2(clabel, (*C.int)(&v[0]), C.ImGuiInputTextFlags(flags)))
}

// InputInt3 edits a 3-component int bound to v.
func InputInt3(label string, v *[3]int32, flags InputTextFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igInputInt3(clabel, (*C.int)(&v[0]), C.ImGuiInputTextFlags(flags)))
}

// InputInt4 edits a 4-component int bound to v.
func InputInt4(label string, v *[4]int32, flags InputTextFlags) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igInputInt4(clabel, (*C.int)(&v[0]), C.ImGuiInputTextFlags(flags)))
}

// InputDouble edits a double in a box with optional step buttons.
func InputDouble(label string, v *float64, step, stepFast float64, format string, flags InputTextFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputDouble(clabel, (*C.double)(v), C.double(step), C.double(stepFast), cformat, C.ImGuiInputTextFlags(flags)))
}

// InputScalar edits a single value of an arbitrary data type. pData, pStep and
// pStepFast point to values of the given DataType; pStep/pStepFast may be nil.
func InputScalar(label string, dataType DataType, pData, pStep, pStepFast unsafe.Pointer, format string, flags InputTextFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputScalar(clabel, C.ImGuiDataType(dataType), pData, pStep, pStepFast, cformat, C.ImGuiInputTextFlags(flags)))
}

// InputScalarN edits components values of an arbitrary data type stored
// contiguously at pData. pStep/pStepFast may be nil.
func InputScalarN(label string, dataType DataType, pData unsafe.Pointer, components int32, pStep, pStepFast unsafe.Pointer, format string, flags InputTextFlags) bool {
	clabel, cformat := C.CString(label), C.CString(format)
	defer C.free(unsafe.Pointer(clabel))
	defer C.free(unsafe.Pointer(cformat))
	return bool(C.igInputScalarN(clabel, C.ImGuiDataType(dataType), pData, C.int(components), pStep, pStepFast, cformat, C.ImGuiInputTextFlags(flags)))
}
