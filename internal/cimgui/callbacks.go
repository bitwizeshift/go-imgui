package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import (
	"runtime"
	"unsafe"

	"github.com/bitwizeshift/go-imgui/internal/handle"
)

// Key identifies a keyboard key. Mirrors a subset of ImGuiKey relevant to input
// callbacks.
type Key int32

const (
	KeyUpArrow   Key = C.ImGuiKey_UpArrow
	KeyDownArrow Key = C.ImGuiKey_DownArrow
)

// InputTextCallback receives an event during an input-text widget. It returns 0
// in the common case; a non-zero return is event-specific and mirrors the C
// ImGuiInputTextCallback contract.
type InputTextCallback func(data *InputTextCallbackData) int32

// InputTextCallbackData is a view over the live ImGuiInputTextCallbackData passed
// to an [InputTextCallback]. It is only valid for the duration of the callback.
type InputTextCallbackData struct {
	c *C.ImGuiInputTextCallbackData
}

// EventFlag reports which callback event is being delivered (one of the
// InputTextFlagsCallback* values).
func (d *InputTextCallbackData) EventFlag() InputTextFlags {
	return InputTextFlags(d.c.EventFlag)
}

// EventChar is the character about to be inserted, valid during a
// CharFilter event. Setting it to 0 discards the character.
func (d *InputTextCallbackData) EventChar() rune {
	return rune(d.c.EventChar)
}

// SetEventChar replaces the character being filtered; 0 discards it.
func (d *InputTextCallbackData) SetEventChar(r rune) {
	d.c.EventChar = C.ImWchar(r)
}

// EventKey is the key that triggered a History event (KeyUpArrow or
// KeyDownArrow).
func (d *InputTextCallbackData) EventKey() Key {
	return Key(d.c.EventKey)
}

// Buf returns the current text in the edit buffer.
func (d *InputTextCallbackData) Buf() string {
	if d.c.Buf == nil {
		return ""
	}
	return C.GoStringN(d.c.Buf, d.c.BufTextLen)
}

// CursorPos reports the byte offset of the cursor.
func (d *InputTextCallbackData) CursorPos() int {
	return int(d.c.CursorPos)
}

// SetCursorPos moves the cursor to the given byte offset.
func (d *InputTextCallbackData) SetCursorPos(pos int) {
	d.c.CursorPos = C.int(pos)
}

// SelectionStart reports the byte offset of the selection start.
func (d *InputTextCallbackData) SelectionStart() int {
	return int(d.c.SelectionStart)
}

// SelectionEnd reports the byte offset of the selection end.
func (d *InputTextCallbackData) SelectionEnd() int {
	return int(d.c.SelectionEnd)
}

// HasSelection reports whether any text is selected.
func (d *InputTextCallbackData) HasSelection() bool {
	return bool(C.ImGuiInputTextCallbackData_HasSelection(d.c))
}

// SelectAll selects all text in the buffer.
func (d *InputTextCallbackData) SelectAll() {
	C.ImGuiInputTextCallbackData_SelectAll(d.c)
}

// ClearSelection clears the current selection.
func (d *InputTextCallbackData) ClearSelection() {
	C.ImGuiInputTextCallbackData_ClearSelection(d.c)
}

// InsertChars inserts text at the given byte offset.
func (d *InputTextCallbackData) InsertChars(pos int, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.ImGuiInputTextCallbackData_InsertChars(d.c, C.int(pos), ctext, nil)
}

// DeleteChars deletes count bytes starting at the given byte offset.
func (d *InputTextCallbackData) DeleteChars(pos, count int) {
	C.ImGuiInputTextCallbackData_DeleteChars(d.c, C.int(pos), C.int(count))
}

// inputTextBinding carries the per-widget state needed by an input-text callback
// across the cgo boundary. buf, when set, is resized automatically on the
// CallbackResize event; callback receives every other registered event.
type inputTextBinding struct {
	buf      *TextBuffer
	callback InputTextCallback
}

//export goInputTextCallbackTrampoline
func goInputTextCallbackTrampoline(cdata *C.ImGuiInputTextCallbackData) C.int {
	b, _ := handle.Restore(unsafe.Pointer(cdata.UserData)).(*inputTextBinding)
	if b == nil {
		return 0
	}
	if b.buf != nil && InputTextFlags(cdata.EventFlag) == InputTextFlagsCallbackResize {
		b.buf.resize(int(cdata.BufSize))
		cdata.Buf = (*C.char)(b.buf.data)
		return 0
	}
	if b.callback != nil {
		return C.int(b.callback(&InputTextCallbackData{c: cdata}))
	}
	return 0
}

// strGetter resolves an item label by index for the _FnStrPtr widget overloads.
type strGetter func(idx int32) string

// strGetterBinding holds the Go getter plus the C strings it has handed out, so
// they can be freed once the owning widget call returns.
type strGetterBinding struct {
	getter strGetter
	cstrs  []*C.char
}

func (b *strGetterBinding) free() {
	for _, p := range b.cstrs {
		C.free(unsafe.Pointer(p))
	}
}

//export goStrGetterTrampoline
func goStrGetterTrampoline(userData unsafe.Pointer, idx C.int) *C.char {
	b, _ := handle.Restore(userData).(*strGetterBinding)
	if b == nil {
		return nil
	}
	cstr := C.CString(b.getter(int32(idx)))
	b.cstrs = append(b.cstrs, cstr)
	return cstr
}

// floatGetter resolves a plot sample by index for the _FnFloatPtr overloads.
type floatGetter func(idx int32) float32

//export goFloatGetterTrampoline
func goFloatGetterTrampoline(userData unsafe.Pointer, idx C.int) C.float {
	g, _ := handle.Restore(userData).(floatGetter)
	if g == nil {
		return 0
	}
	return C.float(g(int32(idx)))
}

// TextBuffer is a growable, NUL-terminated C-backed text buffer for use with the
// resizable input-text wrappers. Its backing memory is freed by a finalizer, or
// eagerly with [TextBuffer.Free].
type TextBuffer struct {
	data unsafe.Pointer
	size int
}

// NewTextBuffer returns a buffer seeded with s.
func NewTextBuffer(s string) *TextBuffer {
	b := &TextBuffer{}
	b.Set(s)
	runtime.SetFinalizer(b, (*TextBuffer).Free)
	return b
}

// Set replaces the buffer contents with s, growing the backing memory if needed.
func (b *TextBuffer) Set(s string) {
	b.resize(len(s) + 1)
	dst := unsafe.Slice((*byte)(b.data), b.size)
	n := copy(dst, s)
	dst[n] = 0
}

// String returns the buffer contents up to the first NUL.
func (b *TextBuffer) String() string {
	if b.data == nil {
		return ""
	}
	return C.GoString((*C.char)(b.data))
}

// Grow ensures the backing memory holds at least n bytes, leaving the contents
// unchanged. It is a no-op when the buffer is already that large.
func (b *TextBuffer) Grow(n int) {
	b.resize(n)
}

// Free releases the backing memory. It is a no-op if already freed.
func (b *TextBuffer) Free() {
	if b.data != nil {
		C.free(b.data)
		b.data = nil
		b.size = 0
	}
	runtime.SetFinalizer(b, nil)
}

// resize grows the backing memory to at least n bytes.
func (b *TextBuffer) resize(n int) {
	if n <= b.size {
		return
	}
	p := C.realloc(b.data, C.size_t(n))
	if p == nil {
		panic("cimgui: realloc returned nil")
	}
	b.data = p
	b.size = n
}
