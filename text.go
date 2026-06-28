package imgui

import (
	"fmt"

	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// Text draws formatted text. With no args the format string is drawn verbatim;
// otherwise it is formatted with [fmt.Sprintf]. (Formatting happens in Go; the C
// varargs entry point is not used.)
func Text(format string, args ...any) {
	if len(args) == 0 {
		cimgui.TextUnformatted(format)
		return
	}
	cimgui.TextUnformatted(fmt.Sprintf(format, args...))
}

// TextUnformatted draws text verbatim, without any formatting.
func TextUnformatted(text string) { cimgui.TextUnformatted(text) }

// SeparatorText draws a horizontal separator with a centered label.
func SeparatorText(label string) { cimgui.SeparatorText(label) }
