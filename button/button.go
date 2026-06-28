// Package button provides button widgets.
package button

import (
	"github.com/bitwizeshift/go-imgui"
	"github.com/bitwizeshift/go-imgui/internal/cimgui"
)

// Button is a clickable button. OnClick (if set) fires when pressed; the press
// state can also be polled with [Button.Pressed].
type Button struct {
	Label   string
	Size    imgui.Vec2 // zero auto-fits the label
	Small   bool       // draw without frame padding
	OnClick func()
	pressed bool
}

// New returns a button labelled label.
func New(label string) *Button { return &Button{Label: label} }

// WithDimensions returns a button of the given size.
func WithDimensions(label string, width, height float32) *Button {
	return &Button{Label: label, Size: imgui.Vec2{X: width, Y: height}}
}

// Small returns a button drawn without frame padding.
func Small(label string) *Button {
	return &Button{Label: label, Small: true}
}

// Display draws the button.
func (b *Button) Display() {
	if b.Small {
		b.pressed = cimgui.SmallButton(b.Label)
	} else {
		b.pressed = cimgui.Button(b.Label, b.Size)
	}
	if b.pressed && b.OnClick != nil {
		b.OnClick()
	}
}

// Pressed reports whether the button was pressed during the last [Button.Display].
func (b *Button) Pressed() bool { return b.pressed }

// Direction is the arrow direction of an [Arrow] button.
type Direction = cimgui.Dir

// Arrow directions.
const (
	Left  Direction = cimgui.DirLeft
	Right Direction = cimgui.DirRight
	Up    Direction = cimgui.DirUp
	Down  Direction = cimgui.DirDown
)

// Arrow is a square button containing a direction arrow. ID must be unique.
type Arrow struct {
	ID      string
	Dir     Direction
	OnClick func()
	pressed bool
}

// NewArrow returns an arrow button.
func NewArrow(id string, dir Direction) *Arrow {
	return &Arrow{ID: id, Dir: dir}
}

// LeftArrow returns a left-pointing arrow button.
func LeftArrow(id string) *Arrow {
	return NewArrow(id, Left)
}

// RightArrow returns a right-pointing arrow button.
func RightArrow(id string) *Arrow {
	return NewArrow(id, Right)
}

// UpArrow returns an up-pointing arrow button.
func UpArrow(id string) *Arrow {
	return NewArrow(id, Up)
}

// DownArrow returns a down-pointing arrow button.
func DownArrow(id string) *Arrow {
	return NewArrow(id, Down)
}

// Display draws the arrow button.
func (a *Arrow) Display() {
	a.pressed = cimgui.ArrowButton(a.ID, a.Dir)
	if a.pressed && a.OnClick != nil {
		a.OnClick()
	}
}

// Pressed reports whether the arrow was pressed during the last [Arrow.Display].
func (a *Arrow) Pressed() bool {
	return a.pressed
}

// Invisible is a sized button that draws nothing but reports hover and click,
// useful as a custom interactive region. ID must be unique.
type Invisible struct {
	ID      string
	Size    imgui.Vec2
	OnClick func()
	pressed bool
}

// NewInvisible returns an invisible button of the given size.
func NewInvisible(id string, size imgui.Vec2) *Invisible {
	return &Invisible{ID: id, Size: size}
}

// Display draws the invisible button.
func (b *Invisible) Display() {
	b.pressed = cimgui.InvisibleButton(b.ID, b.Size, cimgui.ButtonFlagsMouseButtonLeft)
	if b.pressed && b.OnClick != nil {
		b.OnClick()
	}
}

// Pressed reports whether the button was pressed during the last [Invisible.Display].
func (b *Invisible) Pressed() bool {
	return b.pressed
}

var _ imgui.Widget = (*Invisible)(nil)

// Bullet draws a small bullet glyph and stays on the same line, so the next
// widget renders beside it. For a bulleted line of text use text.Bullet instead.
type Bullet struct{}

// NewBullet returns a bullet glyph.
func NewBullet() *Bullet {
	return &Bullet{}
}

// Display draws the bullet.
func (b *Bullet) Display() {
	cimgui.Bullet()
}

var _ imgui.Widget = (*Bullet)(nil)

// ProgressBar shows progress in the range 0..1. A zero Size auto-fits; Overlay,
// when set, is centered text drawn over the bar.
type ProgressBar struct {
	Fraction float32
	Size     imgui.Vec2
	Overlay  string
}

// NewProgressBar returns a progress bar filled to fraction.
func NewProgressBar(fraction float32) *ProgressBar {
	return &ProgressBar{Fraction: fraction}
}

// Display draws the progress bar.
func (p *ProgressBar) Display() {
	cimgui.ProgressBar(p.Fraction, p.Size, p.Overlay)
}
