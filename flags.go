package imgui

// WindowFlags configures a window opened with [BeginV].
type WindowFlags int

// Window flags. See Dear ImGui's ImGuiWindowFlags_ for the full set.
const (
	WindowFlagsNone                WindowFlags = 0
	WindowFlagsNoTitleBar          WindowFlags = 1 << 0
	WindowFlagsNoResize            WindowFlags = 1 << 1
	WindowFlagsNoMove              WindowFlags = 1 << 2
	WindowFlagsNoScrollbar         WindowFlags = 1 << 3
	WindowFlagsNoScrollWithMouse   WindowFlags = 1 << 4
	WindowFlagsNoCollapse          WindowFlags = 1 << 5
	WindowFlagsAlwaysAutoResize    WindowFlags = 1 << 6
	WindowFlagsNoBackground        WindowFlags = 1 << 7
	WindowFlagsNoSavedSettings     WindowFlags = 1 << 8
	WindowFlagsNoMouseInputs       WindowFlags = 1 << 9
	WindowFlagsMenuBar             WindowFlags = 1 << 10
	WindowFlagsHorizontalScrollbar WindowFlags = 1 << 11
	WindowFlagsNoFocusOnAppearing  WindowFlags = 1 << 12
	WindowFlagsNoNavInputs         WindowFlags = 1 << 16
)

// ChildFlags configures a child region opened with [BeginChildV].
type ChildFlags int

// Child flags. See Dear ImGui's ImGuiChildFlags_ for the full set.
const (
	ChildFlagsNone        ChildFlags = 0
	ChildFlagsBorders     ChildFlags = 1 << 0
	ChildFlagsAutoResizeX ChildFlags = 1 << 4
	ChildFlagsAutoResizeY ChildFlags = 1 << 5
)

// Cond selects when a state-setting call applies (e.g. [SetNextWindowSizeV]).
type Cond int

// Conditions. See Dear ImGui's ImGuiCond_.
const (
	CondNone         Cond = 0
	CondAlways       Cond = 1 << 0
	CondOnce         Cond = 1 << 1
	CondFirstUseEver Cond = 1 << 2
	CondAppearing    Cond = 1 << 3
)

// SliderFlags configures sliders and drags (e.g. [SliderFloatV]).
type SliderFlags int

// Slider flags. See Dear ImGui's ImGuiSliderFlags_.
const (
	SliderFlagsNone        SliderFlags = 0
	SliderFlagsLogarithmic SliderFlags = 1 << 5
)

// InputTextFlags configures input widgets (e.g. [InputIntV]).
type InputTextFlags int

// Input flags. See Dear ImGui's ImGuiInputTextFlags_ for the full set.
const (
	InputTextFlagsNone     InputTextFlags = 0
	InputTextFlagsReadOnly InputTextFlags = 1 << 14
)

// ComboFlags configures a combo box opened with [BeginComboV].
type ComboFlags int

// Combo flags. See Dear ImGui's ImGuiComboFlags_ for the full set.
const (
	ComboFlagsNone           ComboFlags = 0
	ComboFlagsPopupAlignLeft ComboFlags = 1 << 0
)

// ColorEditFlags configures color editors (e.g. [ColorEdit3V]).
type ColorEditFlags int

// Color-edit flags. See Dear ImGui's ImGuiColorEditFlags_ for the full set.
const (
	ColorEditFlagsNone       ColorEditFlags = 0
	ColorEditFlagsNoAlpha    ColorEditFlags = 1 << 1
	ColorEditFlagsDisplayHSV ColorEditFlags = 1 << 21
	ColorEditFlagsFloat      ColorEditFlags = 1 << 24
)

// TreeNodeFlags configures tree nodes and headers (e.g. [CollapsingHeaderV]).
type TreeNodeFlags int

// Tree-node flags. See Dear ImGui's ImGuiTreeNodeFlags_ for the full set.
const (
	TreeNodeFlagsNone        TreeNodeFlags = 0
	TreeNodeFlagsSelected    TreeNodeFlags = 1 << 0
	TreeNodeFlagsFramed      TreeNodeFlags = 1 << 1
	TreeNodeFlagsDefaultOpen TreeNodeFlags = 1 << 5
)

// SelectableFlags configures selectables (e.g. [SelectableV]).
type SelectableFlags int

// Selectable flags. See Dear ImGui's ImGuiSelectableFlags_ for the full set.
const (
	SelectableFlagsNone              SelectableFlags = 0
	SelectableFlagsNoAutoClosePopups SelectableFlags = 1 << 0
)
