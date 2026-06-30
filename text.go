package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// TextUnformatted draws text verbatim, with no printf-style formatting applied.
// It models ImGui::TextUnformatted.
func TextUnformatted(text string) {
	cimgui.TextUnformatted(text)
}

// TextColored draws text in the given RGBA color. It models ImGui::TextColored.
func TextColored(col Vec4, text string) {
	cimgui.TextColored(col, text)
}

// TextDisabled draws text using the disabled text color. It models
// ImGui::TextDisabled.
func TextDisabled(text string) {
	cimgui.TextDisabled(text)
}

// TextWrapped draws text, wrapping at the window's right edge. It models
// ImGui::TextWrapped.
func TextWrapped(text string) {
	cimgui.TextWrapped(text)
}

// LabelText draws a value on the left and a label on the right. It models
// ImGui::LabelText.
func LabelText(label, text string) {
	cimgui.LabelText(label, text)
}

// BulletText draws text prefixed with a bullet. It models ImGui::BulletText.
func BulletText(text string) {
	cimgui.BulletText(text)
}

// SeparatorText draws a horizontal separator with a centered label. It models
// ImGui::SeparatorText.
func SeparatorText(label string) {
	cimgui.SeparatorText(label)
}

// TextLink draws text styled as a hyperlink and reports whether it was clicked.
// It models ImGui::TextLink.
func TextLink(label string) bool {
	return cimgui.TextLink(label)
}

// TextLinkOpenURL draws a hyperlink that opens url when clicked, reporting
// whether it was clicked. It models ImGui::TextLinkOpenURL.
func TextLinkOpenURL(label, url string) bool {
	return cimgui.TextLinkOpenURL(label, url)
}
