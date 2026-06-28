package input

import (
	"bytes"

	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// Text edits a bound string. It keeps an internal byte buffer (capacity MaxLen,
// default 256) synced with *Value each frame, so callers work purely in terms of
// a *string.
type Text struct {
	Label     string
	Value     *string
	Hint      string     // placeholder shown when empty (single-line only)
	MaxLen    int        // buffer capacity in bytes; default 256
	Multiline bool       // draw a multi-line box of Size
	Size      imgui.Vec2 // multiline box size; zero auto-fits
	OnChange  func(string)
	textFlags
	changed   bool
	buf       []byte
	scratch   string
}

// NewText returns a text input bound to value.
func NewText(label string, value *string) *Text { return &Text{Label: label, Value: value} }

// Display draws the text input.
func (t *Text) Display() {
	val := t.Value
	if val == nil {
		val = &t.scratch
	}
	max := t.MaxLen
	if max <= 0 {
		max = 256
	}
	if cap(t.buf) < max {
		t.buf = make([]byte, max)
	} else {
		t.buf = t.buf[:max]
	}
	// Seed the buffer from the current value, NUL-terminated.
	n := copy(t.buf, *val)
	if n >= len(t.buf) {
		n = len(t.buf) - 1
	}
	t.buf[n] = 0

	switch {
	case t.Multiline:
		t.changed = cimgui.InputTextMultiline(t.Label, t.buf, t.Size, t.flags)
	case t.Hint != "":
		t.changed = cimgui.InputTextWithHint(t.Label, t.Hint, t.buf, t.flags)
	default:
		t.changed = cimgui.InputText(t.Label, t.buf, t.flags)
	}

	if t.changed {
		end := bytes.IndexByte(t.buf, 0)
		if end < 0 {
			end = len(t.buf)
		}
		s := string(t.buf[:end])
		*val = s
		if t.OnChange != nil {
			t.OnChange(s)
		}
	}
}

// Changed reports whether the text changed during the last Display.
func (t *Text) Changed() bool { return t.changed }
