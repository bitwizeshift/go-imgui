// Package tooltip attaches hover tooltips to other widgets.
package tooltip

import (
	"fmt"

	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// Tooltip wraps a widget and shows Text when that widget is hovered. It is a
// widget itself, so it composes anywhere a widget is expected.
type Tooltip struct {
	Target imgui.Widget
	Text   string
}

// For returns a tooltip showing text when target is hovered.
func For(target imgui.Widget, text string) *Tooltip {
	return &Tooltip{Target: target, Text: text}
}

// Forf is [For] with a [fmt.Sprintf]-formatted text.
func Forf(target imgui.Widget, format string, args ...any) *Tooltip {
	return &Tooltip{Target: target, Text: fmt.Sprintf(format, args...)}
}

// Display draws the target and its tooltip.
func (t *Tooltip) Display() {
	if t.Target != nil {
		t.Target.Display()
	}
	if cimgui.IsItemHovered(cimgui.HoveredFlagsForTooltip) {
		cimgui.SetTooltip(t.Text)
	}
}
