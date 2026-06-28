package input

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// HistoryDir is the direction of a history-recall request delivered to a [Text]
// widget's OnHistory handler.
type HistoryDir int

const (
	// HistoryUp recalls the previous entry (Up arrow).
	HistoryUp HistoryDir = iota
	// HistoryDown recalls the next entry (Down arrow).
	HistoryDown
)

// CallbackData gives a [Text] callback access to and control over the live edit
// buffer. It is only valid for the duration of the callback it is passed to.
type CallbackData struct {
	c *cimgui.InputTextCallbackData
}

// Buf returns the current text in the edit buffer.
func (d *CallbackData) Buf() string {
	return d.c.Buf()
}

// CursorPos reports the byte offset of the cursor.
func (d *CallbackData) CursorPos() int {
	return d.c.CursorPos()
}

// SetCursorPos moves the cursor to the given byte offset.
func (d *CallbackData) SetCursorPos(pos int) {
	d.c.SetCursorPos(pos)
}

// SelectionStart reports the byte offset of the selection start.
func (d *CallbackData) SelectionStart() int {
	return d.c.SelectionStart()
}

// SelectionEnd reports the byte offset of the selection end.
func (d *CallbackData) SelectionEnd() int {
	return d.c.SelectionEnd()
}

// HasSelection reports whether any text is selected.
func (d *CallbackData) HasSelection() bool {
	return d.c.HasSelection()
}

// SelectAll selects all text in the buffer.
func (d *CallbackData) SelectAll() {
	d.c.SelectAll()
}

// ClearSelection clears the current selection.
func (d *CallbackData) ClearSelection() {
	d.c.ClearSelection()
}

// InsertChars inserts text at the given byte offset.
func (d *CallbackData) InsertChars(pos int, text string) {
	d.c.InsertChars(pos, text)
}

// DeleteChars deletes count bytes starting at the given byte offset.
func (d *CallbackData) DeleteChars(pos, count int) {
	d.c.DeleteChars(pos, count)
}
