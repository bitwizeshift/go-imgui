package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
// #include "texture.h"
import "C"

import "unsafe"

// TextureRef refers to a texture the renderer backend can draw. It wraps Dear
// ImGui's ImTextureRef; for the OpenGL3 backend the underlying ID is a GL
// texture name.
type TextureRef struct {
	ref C.ImTextureRef_c
}

// TextureRefFromID builds a [TextureRef] from a backend texture identifier.
func TextureRefFromID(id uint64) TextureRef {
	var r C.ImTextureRef_c
	r._TexID = C.ImTextureID(id)
	return TextureRef{ref: r}
}

// ID returns the backend texture identifier behind the ref (the GL texture name
// for the OpenGL3 backend).
func (t TextureRef) ID() uint64 { return uint64(t.ref._TexID) }

// CreateTextureRGBA uploads width*height tightly-packed RGBA8 pixels (4 bytes
// per pixel, len(pixels) == width*height*4) to a new GPU texture and returns a
// [TextureRef] for it. A GL context must be current (it is during an app.Run
// frame). Free it with [DeleteTexture] when no longer needed.
func CreateTextureRGBA(width, height int, pixels []byte) TextureRef {
	var p *C.uchar
	if len(pixels) > 0 {
		p = (*C.uchar)(unsafe.Pointer(&pixels[0]))
	}
	id := C.backendCreateTextureRGBA(p, C.int(width), C.int(height))
	return TextureRefFromID(uint64(id))
}

// DeleteTexture deletes a texture created by [CreateTextureRGBA].
func DeleteTexture(ref TextureRef) { C.backendDeleteTexture(C.uint(ref.ID())) }

// FontAtlasTexRef returns the [TextureRef] of the current context's font atlas.
// It is handy for exercising the image widgets without loading an external
// image, and is valid once the backend has uploaded the atlas (after the first
// frame). It must be called with an active context.
func FontAtlasTexRef() TextureRef {
	io := C.igGetIO_Nil()
	return TextureRef{ref: io.Fonts.TexRef}
}

// Image draws texRef as an image. uv0 and uv1 are the texture coordinates of the
// top-left and bottom-right corners ({0,0} and {1,1} for the whole texture).
func Image(texRef TextureRef, size, uv0, uv1 Vec2) {
	C.igImage(texRef.ref, size.c(), uv0.c(), uv1.c())
}

// ImageWithBg draws texRef over a background color and tinted by tintCol.
func ImageWithBg(texRef TextureRef, size, uv0, uv1 Vec2, bgCol, tintCol Vec4) {
	C.igImageWithBg(texRef.ref, size.c(), uv0.c(), uv1.c(), bgCol.c(), tintCol.c())
}

// ImageButton draws a clickable image button and reports whether it was clicked.
func ImageButton(strID string, texRef TextureRef, size, uv0, uv1 Vec2, bgCol, tintCol Vec4) bool {
	cid := C.CString(strID)
	defer C.free(unsafe.Pointer(cid))
	return bool(C.igImageButton(cid, texRef.ref, size.c(), uv0.c(), uv1.c(), bgCol.c(), tintCol.c()))
}
