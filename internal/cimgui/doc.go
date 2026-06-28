// package cimgui is the single cgo boundary for the Dear ImGui wrapper.
//
// It is the only package in the module that uses `import "C"`. It compiles, from
// the vendored sources copied into csources/ (see gencsources.go): the cimgui C
// API, its bundled Dear ImGui, the GLFW platform backend, the OpenGL3 renderer
// backend, and GLFW itself. On top of that it exposes Go functions that map 1:1
// onto the cimgui C API with Go types substituted for C ones.
//
// It is deliberately internal and intentionally un-ergonomic: the curated,
// idiomatic surface lives in the public packages that import this one.
//
// # Conventions
//
//   - Functions are named after their cimgui C counterpart with the "ig" prefix
//     stripped. Overloaded calls keep cimgui's underscore disambiguator verbatim
//     (igRadioButton_Bool becomes [RadioButton_Bool], igBeginChild_ID becomes
//     [BeginChild_ID]); calls with no overload get the bare name ([Button]).
//   - Each ImGui enum is exposed as a typed Go type backed by int32 with
//     idiomatic constant names ([WindowFlags], WindowFlagsNoTitleBar). The values
//     are taken directly from the C enum, never hardcoded.
//   - ImVec2/ImVec4 are represented by [Vec2]/[Vec4]; output pointers and scalars
//     are passed as Go pointers and copied back.
//
// # Callbacks across cgo
//
// Functions that take a C callback (InputText events, the _FnStrPtr/_FnFloatPtr
// getter overloads) route a Go closure across the boundary via [handle]: the
// closure is registered to obtain an opaque token passed as user_data, recovered
// in a //export trampoline (see callbacks.go), and released when the call
// returns. [InputTextResizable] and friends additionally grow a [TextBuffer] in
// place through the resize callback.
//
// # Not yet exposed
//
//   - printf-style variadic functions cannot be called through cgo. The common
//     ones are provided via small extern "C" shims in shims.cpp (see [TextColored],
//     [TextWrapped], [SetTooltip], ...) that forward a pre-built string.
//
// # Textures
//
// [Image], [ImageWithBg] and [ImageButton] take a [TextureRef]. Build one with
// [CreateTextureRGBA] (which uploads raw RGBA8 pixels to a GL texture via the
// extern "C" helpers in texture.cpp), [TextureRefFromID] for a texture you
// created yourself, or [FontAtlasTexRef] for Dear ImGui's font atlas.
//
// CIMGUI_DEFINE_ENUMS_AND_STRUCTS is defined for C compilation (the cgo preamble)
// so cimgui.h exposes the C structs and enums, but not for the C++ translation
// units, where cimgui.cpp must see the real Dear ImGui types.
package cimgui
