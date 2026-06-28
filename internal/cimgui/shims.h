// Hand-written extern "C" shims for Dear ImGui functions that are variadic
// (printf-style). cgo cannot call C variadic functions directly, so each shim
// forwards a single pre-built string through a "%s" format, preserving the
// exact rendering behaviour of the original call.
#ifndef CIMGUI_SHIMS_H
#define CIMGUI_SHIMS_H

#ifdef __cplusplus
extern "C" {
#endif

// shimTextColored renders s in the given RGBA color (mirrors ImGui::TextColored).
void shimTextColored(float r, float g, float b, float a, const char *s);

// shimTextDisabled renders s using the disabled text color.
void shimTextDisabled(const char *s);

// shimTextWrapped renders s, wrapping at the window's edge.
void shimTextWrapped(const char *s);

// shimLabelText renders a label/value pair (value on the left, label on the right).
void shimLabelText(const char *label, const char *s);

// shimBulletText renders s prefixed with a bullet.
void shimBulletText(const char *s);

// shimSetTooltip sets the contents of the current tooltip to s.
void shimSetTooltip(const char *s);

#ifdef __cplusplus
}
#endif

#endif // CIMGUI_SHIMS_H
