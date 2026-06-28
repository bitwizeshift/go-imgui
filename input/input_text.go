package input

import (
	"github.com/bitwizeshift/go-imgui"
	"github.com/bitwizeshift/go-imgui/internal/cimgui"
)

// Text edits a bound string. It keeps an internal buffer that grows automatically
// as the text does, so callers work purely in terms of a *string. The callback
// fields, when set, give fine-grained control over editing; their event flags are
// enabled automatically.
type Text struct {
	Label     string
	Value     *string
	Hint      string     // placeholder shown when empty (single-line only)
	MaxLen    int        // initial buffer capacity hint in bytes; default 256
	Multiline bool       // draw a multi-line box of Size
	Size      imgui.Vec2 // multiline box size; zero auto-fits
	OnChange  func(string)
	// OnEdit fires whenever the text is edited, before OnChange.
	OnEdit func()
	// OnCompletion fires on a completion request (Tab).
	OnCompletion func(*CallbackData)
	// OnHistory fires on an Up/Down history request.
	OnHistory func(*CallbackData, HistoryDir)
	// OnCharFilter, when set, is called for each typed rune; it returns the rune
	// to insert, or 0 to discard the character.
	OnCharFilter func(rune) rune
	textFlags
	changed bool
	buf     *cimgui.TextBuffer
	scratch string
}

// NewText returns a text input bound to value.
func NewText(label string, value *string) *Text {
	return &Text{Label: label, Value: value}
}

// Display draws the text input.
func (t *Text) Display() {
	val := t.Value
	if val == nil {
		val = &t.scratch
	}
	t.syncBuffer(*val)

	flags := t.effectiveFlags()
	cb := t.callback()
	switch {
	case t.Multiline:
		t.changed = cimgui.InputTextMultilineResizable(t.Label, t.buf, t.Size, flags, cb)
	case t.Hint != "":
		t.changed = cimgui.InputTextWithHintResizable(t.Label, t.Hint, t.buf, flags, cb)
	default:
		t.changed = cimgui.InputTextResizable(t.Label, t.buf, flags, cb)
	}

	if t.changed {
		s := t.buf.String()
		*val = s
		if t.OnChange != nil {
			t.OnChange(s)
		}
	}
}

// syncBuffer seeds the internal buffer from the current value, creating it on
// first use and honoring MaxLen as an initial capacity hint.
func (t *Text) syncBuffer(val string) {
	if t.buf == nil {
		t.buf = cimgui.NewTextBuffer(val)
		max := t.MaxLen
		if max <= 0 {
			max = 256
		}
		t.buf.Grow(max)
		return
	}
	t.buf.Set(val)
}

// effectiveFlags ORs in the callback-event flags implied by the handler fields.
func (t *Text) effectiveFlags() cimgui.InputTextFlags {
	f := t.flags
	if t.OnCompletion != nil {
		f |= cimgui.InputTextFlagsCallbackCompletion
	}
	if t.OnHistory != nil {
		f |= cimgui.InputTextFlagsCallbackHistory
	}
	if t.OnEdit != nil {
		f |= cimgui.InputTextFlagsCallbackEdit
	}
	if t.OnCharFilter != nil {
		f |= cimgui.InputTextFlagsCallbackCharFilter
	}
	return f
}

// callback returns the cimgui callback dispatching to the handler fields, or nil
// when none are set.
func (t *Text) callback() cimgui.InputTextCallback {
	if t.OnCompletion == nil && t.OnHistory == nil && t.OnEdit == nil && t.OnCharFilter == nil {
		return nil
	}
	return func(data *cimgui.InputTextCallbackData) int32 {
		switch data.EventFlag() {
		case cimgui.InputTextFlagsCallbackCharFilter:
			r := t.OnCharFilter(data.EventChar())
			if r == 0 {
				return 1 // discard the character
			}
			data.SetEventChar(r)
		case cimgui.InputTextFlagsCallbackCompletion:
			t.OnCompletion(&CallbackData{c: data})
		case cimgui.InputTextFlagsCallbackHistory:
			dir := HistoryUp
			if data.EventKey() == cimgui.KeyDownArrow {
				dir = HistoryDown
			}
			t.OnHistory(&CallbackData{c: data}, dir)
		case cimgui.InputTextFlagsCallbackEdit:
			t.OnEdit()
		}
		return 0
	}
}

// Changed reports whether the text changed during the last Display.
func (t *Text) Changed() bool {
	return t.changed
}
