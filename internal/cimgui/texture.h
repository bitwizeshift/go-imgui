// Hand-written extern "C" helpers for creating renderer textures from raw RGBA
// pixels. They keep all OpenGL calls on the C side (like backend.cpp) so Go code
// can produce a texture for the image widgets without any GL bindings.
#ifndef CIMGUI_TEXTURE_H
#define CIMGUI_TEXTURE_H

#ifdef __cplusplus
extern "C" {
#endif

// backendCreateTextureRGBA uploads width*height tightly-packed RGBA8 pixels to a
// new GL texture and returns its name (id). A current GL context is required.
unsigned int backendCreateTextureRGBA(const unsigned char *pixels, int width, int height);

// backendDeleteTexture deletes a texture created by backendCreateTextureRGBA.
void backendDeleteTexture(unsigned int tex);

#ifdef __cplusplus
}
#endif

#endif // CIMGUI_TEXTURE_H
