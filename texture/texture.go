// Package texture uploads images to the GPU and draws them. It interoperates
// with the standard library's image package: see [FromImage].
package texture

import (
	"image"
	"image/draw"
	"io"

	"github.com/bitwizeshift/go-imgui/internal/cimgui"
)

// Texture is a GPU texture handle.
type Texture struct {
	ref cimgui.TextureRef
}

// FromImage uploads any image.Image as a texture, converting it to RGBA as needed. A
// GL context must be current (it is during an app.Run frame).
func FromImage(img image.Image) *Texture {
	b := img.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), img, b.Min, draw.Src)
	return NewRGBA(b.Dx(), b.Dy(), dst.Pix)
}

// FromReader decodes an image from r and uploads it as a texture. A GL context
// must be current (it is during an app.Run frame).
func FromReader(r io.Reader) (*Texture, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return FromImage(img), nil
}

// NewRGBA uploads tightly-packed RGBA8 pixels (len(pix) == width*height*4) as a
// texture. A GL context must be current.
func NewRGBA(width, height int, pix []byte) *Texture {
	ref := cimgui.CreateTextureRGBA(width, height, pix)
	t := &Texture{ref: ref}
	return t
}

// Close releases the texture immediately. It is optional; an unused texture is
// freed automatically. Close is safe to call more than once.
func (t *Texture) Close() error {
	if t == nil {
		return nil
	}
	cimgui.DeleteTexture(t.ref)
	return nil
}
