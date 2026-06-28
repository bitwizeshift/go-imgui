package texture

import (
	"github.com/bitwizeshift/go-imgui"
	"github.com/bitwizeshift/go-imgui/internal/cimgui"
)

// Image draws a [Texture] at the given size (zero uses the texture's full size).
type Image struct {
	Tex  *Texture
	Size imgui.Vec2
}

// NewImage returns an image widget for tex.
func NewImage(tex *Texture, size imgui.Vec2) *Image {
	return &Image{Tex: tex, Size: size}
}

// Display draws the image.
func (i *Image) Display() {
	if i.Tex == nil {
		return
	}
	cimgui.Image(i.Tex.ref, i.Size, imgui.Vec2{}, imgui.Vec2{X: 1, Y: 1})
}

// Button draws a clickable image button. ID must be unique.
type Button struct {
	ID      string
	Tex     *Texture
	Size    imgui.Vec2
	OnClick func()
	pressed bool
}

// NewButton returns an image button for tex.
func NewButton(id string, tex *Texture, size imgui.Vec2) *Button {
	return &Button{ID: id, Tex: tex, Size: size}
}

// Display draws the image button.
func (b *Button) Display() {
	if b.Tex == nil {
		return
	}
	b.pressed = cimgui.ImageButton(
		b.ID, b.Tex.ref, b.Size,
		imgui.Vec2{},
		imgui.Vec2{X: 1, Y: 1},
		imgui.Vec4{},
		imgui.Vec4{X: 1, Y: 1, Z: 1, W: 1},
	)
	if b.pressed && b.OnClick != nil {
		b.OnClick()
	}
}

// Pressed reports whether the button was clicked during the last Display.
func (b *Button) Pressed() bool { return b.pressed }
