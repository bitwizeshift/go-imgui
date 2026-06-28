// Package tooltip attaches hover tooltips to other widgets.
package tooltip

import (
	"fmt"

	"github.com/bitwizeshift/go-imgui"
	"github.com/bitwizeshift/go-imgui/internal/cimgui"
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

// Rich wraps a target widget and shows arbitrary child widgets in a tooltip when
// the target is hovered, for content richer than a single string.
type Rich struct {
	Target  imgui.Widget
	Widgets []imgui.Widget
}

// ForWidgets returns a rich tooltip showing widgets when target is hovered.
func ForWidgets(target imgui.Widget, widgets ...imgui.Widget) *Rich {
	return &Rich{Target: target, Widgets: widgets}
}

// AddWidget appends tooltip-body widgets.
func (r *Rich) AddWidget(ws ...imgui.Widget) { r.Widgets = append(r.Widgets, ws...) }

// Display draws the target and, when hovered, its rich tooltip.
func (r *Rich) Display() {
	if r.Target != nil {
		r.Target.Display()
	}
	if cimgui.BeginItemTooltip() {
		displayAll(r.Widgets)
		cimgui.EndTooltip()
	}
}

var _ imgui.Widget = (*Rich)(nil)

// Panel is a bare tooltip: when displayed it shows its child widgets as a tooltip
// unconditionally, leaving any hover gating to the caller. Most uses want [For] or
// [Rich] instead.
type Panel struct {
	Widgets []imgui.Widget
}

// NewPanel returns a bare tooltip panel wrapping widgets.
func NewPanel(widgets ...imgui.Widget) *Panel { return &Panel{Widgets: widgets} }

// AddWidget appends tooltip-body widgets.
func (p *Panel) AddWidget(ws ...imgui.Widget) { p.Widgets = append(p.Widgets, ws...) }

// Display draws the tooltip body.
func (p *Panel) Display() {
	if cimgui.BeginTooltip() {
		displayAll(p.Widgets)
		cimgui.EndTooltip()
	}
}

var _ imgui.Widget = (*Panel)(nil)

func displayAll(ws []imgui.Widget) {
	for _, w := range ws {
		if w != nil {
			w.Display()
		}
	}
}
