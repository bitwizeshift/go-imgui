// Package text provides text-display widgets (labels, colored/disabled/wrapped
// text, bullets, separators and links).
package text

import (
	"fmt"

	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// Text displays a single run of text. At most one style applies, in the order
// Bullet, Color, Disabled, Wrapped; otherwise the text is drawn plain.
type Text struct {
	Content  string
	Color    *imgui.Color // when non-nil, draw in this color
	Disabled bool
	Wrapped  bool
	Bullet   bool
}

// New returns plain text.
func New(content string) *Text { return &Text{Content: content} }

// Newf returns plain text formatted with [fmt.Sprintf].
func Newf(format string, args ...any) *Text { return &Text{Content: fmt.Sprintf(format, args...)} }

// Label is an alias of [New].
func Label(content string) *Text { return New(content) }

// Labelf is an alias of [Newf].
func Labelf(format string, args ...any) *Text { return Newf(format, args...) }

// Colored returns text drawn in c.
func Colored(c imgui.Color, content string) *Text { return &Text{Content: content, Color: &c} }

// Disabled returns text in the disabled (greyed) color.
func Disabled(content string) *Text { return &Text{Content: content, Disabled: true} }

// Wrapped returns text that wraps at the window edge.
func Wrapped(content string) *Text { return &Text{Content: content, Wrapped: true} }

// Bullet returns text prefixed with a bullet.
func Bullet(content string) *Text { return &Text{Content: content, Bullet: true} }

// Display draws the text.
func (t *Text) Display() {
	switch {
	case t.Bullet:
		cimgui.BulletText(t.Content)
	case t.Color != nil:
		cimgui.TextColored(t.Color.Vec4(), t.Content)
	case t.Disabled:
		cimgui.TextDisabled(t.Content)
	case t.Wrapped:
		cimgui.TextWrapped(t.Content)
	default:
		cimgui.TextUnformatted(t.Content)
	}
}

// LabelText draws a label/value pair: Value on the left, Label on the right,
// aligned like the other labelled widgets.
type LabelText struct {
	Label string
	Value string
}

// NewLabelText returns a label/value pair.
func NewLabelText(label, value string) *LabelText {
	return &LabelText{Label: label, Value: value}
}

// NewLabelTextf returns a label/value pair with a [fmt.Sprintf]-formatted value.
func NewLabelTextf(label, format string, args ...any) *LabelText {
	return &LabelText{Label: label, Value: fmt.Sprintf(format, args...)}
}

// Display draws the label/value pair.
func (l *LabelText) Display() { cimgui.LabelText(l.Label, l.Value) }

var _ imgui.Widget = (*LabelText)(nil)

// Separator draws a horizontal separator, optionally with a centered label.
type Separator struct {
	Label string
}

// SeparatorText returns a labelled separator.
func SeparatorText(label string) *Separator { return &Separator{Label: label} }

// Display draws the separator.
func (s *Separator) Display() {
	if s.Label == "" {
		cimgui.Separator()
		return
	}
	cimgui.SeparatorText(s.Label)
}

// Link displays a clickable hyperlink. When URL is set the system browser opens
// it on click; otherwise OnClick (if set) is invoked. Use [Link.Clicked] to poll.
type Link struct {
	Label   string
	URL     string
	OnClick func()
	clicked bool
}

// NewLink returns a link with no URL (use OnClick).
func NewLink(label string) *Link { return &Link{Label: label} }

// NewLinkURL returns a link that opens url when clicked.
func NewLinkURL(label, url string) *Link { return &Link{Label: label, URL: url} }

// Display draws the link.
func (l *Link) Display() {
	if l.URL != "" {
		l.clicked = cimgui.TextLinkOpenURL(l.Label, l.URL)
	} else {
		l.clicked = cimgui.TextLink(l.Label)
	}
	if l.clicked && l.OnClick != nil {
		l.OnClick()
	}
}

// Clicked reports whether the link was clicked during the last [Link.Display].
func (l *Link) Clicked() bool { return l.clicked }
