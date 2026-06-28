package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
// #include "shims.h"
import "C"

import "unsafe"

// TextUnformatted draws text verbatim (no printf formatting).
func TextUnformatted(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.igTextUnformatted(ctext, nil)
}

// TextColored draws text in the given RGBA color.
func TextColored(col Vec4, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.shimTextColored(C.float(col.X), C.float(col.Y), C.float(col.Z), C.float(col.W), ctext)
}

// TextDisabled draws text using the disabled text color.
func TextDisabled(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.shimTextDisabled(ctext)
}

// TextWrapped draws text, wrapping at the window's right edge.
func TextWrapped(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.shimTextWrapped(ctext)
}

// LabelText draws a value on the left and a label on the right.
func LabelText(label, text string) {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.shimLabelText(clabel, ctext)
}

// BulletText draws text prefixed with a bullet.
func BulletText(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.shimBulletText(ctext)
}

// SeparatorText draws a horizontal separator with a centered label.
func SeparatorText(label string) {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	C.igSeparatorText(clabel)
}

// TextLink draws text styled as a hyperlink and reports whether it was clicked.
func TextLink(label string) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	return bool(C.igTextLink(clabel))
}

// TextLinkOpenURL draws a hyperlink that opens url when clicked, and reports
// whether it was clicked.
func TextLinkOpenURL(label, url string) bool {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	var curl *C.char
	if url != "" {
		curl = C.CString(url)
		defer C.free(unsafe.Pointer(curl))
	}
	return bool(C.igTextLinkOpenURL(clabel, curl))
}
