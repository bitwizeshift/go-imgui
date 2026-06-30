package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// TextureRef refers to a texture the renderer backend can draw, wrapping Dear
// ImGui's ImTextureRef. It is the texture identifier accepted by [Image],
// [ImageWithBg] and [ImageButton].
type TextureRef = cimgui.TextureRef

// TextureRefFromID builds a [TextureRef] from a backend texture identifier. It
// models constructing an ImTextureRef from an ImTextureID.
func TextureRefFromID(id uint64) TextureRef {
	return cimgui.TextureRefFromID(id)
}

// FontAtlasTexRef returns the [TextureRef] of the current context's font atlas,
// valid once the backend has uploaded the atlas. It models reading
// ImGui::GetIO().Fonts->TexRef.
func FontAtlasTexRef() TextureRef {
	return cimgui.FontAtlasTexRef()
}

// Image draws texRef as an image. uv0 and uv1 are the texture coordinates of the
// top-left and bottom-right corners ({0,0} and {1,1} for the whole texture). It
// models ImGui::Image.
func Image(texRef TextureRef, size, uv0, uv1 Vec2) {
	cimgui.Image(texRef, size, uv0, uv1)
}

// ImageWithBg draws texRef over the background color bgCol and tinted by tintCol.
// It models ImGui::ImageWithBg.
func ImageWithBg(texRef TextureRef, size, uv0, uv1 Vec2, bgCol, tintCol Vec4) {
	cimgui.ImageWithBg(texRef, size, uv0, uv1, bgCol, tintCol)
}

// ImageButton draws a clickable image button and reports whether it was clicked.
// It models ImGui::ImageButton.
func ImageButton(strID string, texRef TextureRef, size, uv0, uv1 Vec2, bgCol, tintCol Vec4) bool {
	return cimgui.ImageButton(strID, texRef, size, uv0, uv1, bgCol, tintCol)
}
