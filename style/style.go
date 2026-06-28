package style

import (
	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// Col identifies a styleable interface color.
type Col = cimgui.Col

// Common interface colors.
const (
	ColText           Col = cimgui.ColText
	ColWindowBg       Col = cimgui.ColWindowBg
	ColChildBg        Col = cimgui.ColChildBg
	ColPopupBg        Col = cimgui.ColPopupBg
	ColBorder         Col = cimgui.ColBorder
	ColFrameBg        Col = cimgui.ColFrameBg
	ColFrameBgHovered Col = cimgui.ColFrameBgHovered
	ColTitleBg        Col = cimgui.ColTitleBg
	ColTitleBgActive  Col = cimgui.ColTitleBgActive
	ColButton         Col = cimgui.ColButton
	ColButtonHovered  Col = cimgui.ColButtonHovered
	ColButtonActive   Col = cimgui.ColButtonActive
	ColHeader         Col = cimgui.ColHeader
	ColText2          Col = cimgui.ColTextDisabled
	ColCheckMark      Col = cimgui.ColCheckMark
	ColPlotLines      Col = cimgui.ColPlotLines
)

// Var identifies a styleable layout variable.
type Var = cimgui.StyleVar

// Common layout variables.
const (
	VarAlpha          Var = cimgui.StyleVarAlpha
	VarWindowRounding Var = cimgui.StyleVarWindowRounding
	VarWindowPadding  Var = cimgui.StyleVarWindowPadding
	VarFrameRounding  Var = cimgui.StyleVarFrameRounding
	VarFramePadding   Var = cimgui.StyleVarFramePadding
	VarItemSpacing    Var = cimgui.StyleVarItemSpacing
	VarGrabRounding   Var = cimgui.StyleVarGrabRounding
	VarTabRounding    Var = cimgui.StyleVarTabRounding
)

// Dark applies the built-in dark preset to the global style.
func Dark() {
	cimgui.StyleColorsDark()
}

// Light applies the built-in light preset to the global style.
func Light() {
	cimgui.StyleColorsLight()
}

// Classic applies the built-in classic preset to the global style.
func Classic() {
	cimgui.StyleColorsClassic()
}

// SetColor sets a single global interface color.
func SetColor(col Col, c imgui.Color) {
	cimgui.GetStyle().SetColor(col, c.Vec4())
}

// Theme is a snapshot of the configurable layout options. Read the current
// values with [Current], adjust the fields, then apply them with [Theme.Apply].
// Colors, when set, are written as part of [Theme.Apply].
type Theme struct {
	Alpha          float32
	DisabledAlpha  float32
	WindowRounding float32
	FrameRounding  float32
	GrabRounding   float32
	TabRounding    float32
	WindowPadding  imgui.Vec2
	FramePadding   imgui.Vec2
	ItemSpacing    imgui.Vec2
	Colors         map[Col]imgui.Color
}

// Current returns the configurable options of the global style.
func Current() Theme {
	s := cimgui.GetStyle()
	return Theme{
		Alpha:          s.Alpha(),
		DisabledAlpha:  s.DisabledAlpha(),
		WindowRounding: s.WindowRounding(),
		FrameRounding:  s.FrameRounding(),
		GrabRounding:   s.GrabRounding(),
		TabRounding:    s.TabRounding(),
		WindowPadding:  s.WindowPadding(),
		FramePadding:   s.FramePadding(),
		ItemSpacing:    s.ItemSpacing(),
	}
}

// Apply writes the theme's options into the global style.
func (t Theme) Apply() {
	s := cimgui.GetStyle()
	s.SetAlpha(t.Alpha)
	s.SetDisabledAlpha(t.DisabledAlpha)
	s.SetWindowRounding(t.WindowRounding)
	s.SetFrameRounding(t.FrameRounding)
	s.SetGrabRounding(t.GrabRounding)
	s.SetTabRounding(t.TabRounding)
	s.SetWindowPadding(t.WindowPadding)
	s.SetFramePadding(t.FramePadding)
	s.SetItemSpacing(t.ItemSpacing)
	for col, c := range t.Colors {
		s.SetColor(col, c.Vec4())
	}
}

// Override is a single style-variable override applied within a [Scoped] block.
// Build one with [FloatVar] or [Vec2Var].
type Override struct {
	v      Var
	vec2   bool
	f      float32
	vector imgui.Vec2
}

// FloatVar overrides a scalar style variable.
func FloatVar(v Var, val float32) Override {
	return Override{v: v, f: val}
}

// Vec2Var overrides a [imgui.Vec2] style variable.
func Vec2Var(v Var, val imgui.Vec2) Override {
	return Override{v: v, vec2: true, vector: val}
}

// push pushes the override onto the style-variable stack.
func (o Override) push() {
	if o.vec2 {
		cimgui.PushStyleVar_Vec2(o.v, o.vector)
		return
	}
	cimgui.PushStyleVar_Float(o.v, o.f)
}

// Scoped applies color and variable overrides while displaying its children,
// then restores the previous style.
type Scoped struct {
	Colors  map[Col]imgui.Color
	Vars    []Override
	Widgets []imgui.Widget
}

var _ imgui.Widget = (*Scoped)(nil)

// NewScoped returns a scoped style block wrapping widgets.
func NewScoped(widgets ...imgui.Widget) *Scoped {
	return &Scoped{Widgets: widgets}
}

// AddWidget appends child widgets.
func (s *Scoped) AddWidget(ws ...imgui.Widget) {
	s.Widgets = append(s.Widgets, ws...)
}

// Display pushes the overrides, draws the children, then pops the overrides.
func (s *Scoped) Display() {
	for col, c := range s.Colors {
		cimgui.PushStyleColor_Vec4(col, c.Vec4())
	}
	for _, v := range s.Vars {
		v.push()
	}
	for _, w := range s.Widgets {
		if w != nil {
			w.Display()
		}
	}
	if len(s.Vars) > 0 {
		cimgui.PopStyleVar(int32(len(s.Vars)))
	}
	if len(s.Colors) > 0 {
		cimgui.PopStyleColor(int32(len(s.Colors)))
	}
}
