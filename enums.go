package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// Dir is a cardinal direction, modelling ImGuiDir. It selects the arrow drawn by
// [ArrowButton] and the direction of various layout helpers.
type Dir = cimgui.Dir

// Direction values for [Dir].
const (
	DirNone  = cimgui.DirNone  // ImGuiDir_None
	DirLeft  = cimgui.DirLeft  // ImGuiDir_Left
	DirRight = cimgui.DirRight // ImGuiDir_Right
	DirUp    = cimgui.DirUp    // ImGuiDir_Up
	DirDown  = cimgui.DirDown  // ImGuiDir_Down
)

// Cond selects when a state-setting call (such as [SetNextWindowPos]) applies.
// It models ImGuiCond.
type Cond = cimgui.Cond

// Condition values for [Cond].
const (
	CondNone         = cimgui.CondNone         // ImGuiCond_None
	CondAlways       = cimgui.CondAlways       // ImGuiCond_Always
	CondOnce         = cimgui.CondOnce         // ImGuiCond_Once
	CondFirstUseEver = cimgui.CondFirstUseEver // ImGuiCond_FirstUseEver
	CondAppearing    = cimgui.CondAppearing    // ImGuiCond_Appearing
)

// Col identifies a styleable interface color, modelling ImGuiCol. It indexes the
// style-color stack used by [StyleColor] and [GetStyleColorVec4].
type Col = cimgui.Col

// Color identifiers for [Col].
const (
	ColText             = cimgui.ColText             // ImGuiCol_Text
	ColTextDisabled     = cimgui.ColTextDisabled     // ImGuiCol_TextDisabled
	ColWindowBg         = cimgui.ColWindowBg         // ImGuiCol_WindowBg
	ColChildBg          = cimgui.ColChildBg          // ImGuiCol_ChildBg
	ColPopupBg          = cimgui.ColPopupBg          // ImGuiCol_PopupBg
	ColBorder           = cimgui.ColBorder           // ImGuiCol_Border
	ColFrameBg          = cimgui.ColFrameBg          // ImGuiCol_FrameBg
	ColFrameBgHovered   = cimgui.ColFrameBgHovered   // ImGuiCol_FrameBgHovered
	ColFrameBgActive    = cimgui.ColFrameBgActive    // ImGuiCol_FrameBgActive
	ColTitleBg          = cimgui.ColTitleBg          // ImGuiCol_TitleBg
	ColTitleBgActive    = cimgui.ColTitleBgActive    // ImGuiCol_TitleBgActive
	ColTitleBgCollapsed = cimgui.ColTitleBgCollapsed // ImGuiCol_TitleBgCollapsed
	ColMenuBarBg        = cimgui.ColMenuBarBg        // ImGuiCol_MenuBarBg
	ColCheckMark        = cimgui.ColCheckMark        // ImGuiCol_CheckMark
	ColSliderGrab       = cimgui.ColSliderGrab       // ImGuiCol_SliderGrab
	ColSliderGrabActive = cimgui.ColSliderGrabActive // ImGuiCol_SliderGrabActive
	ColButton           = cimgui.ColButton           // ImGuiCol_Button
	ColButtonHovered    = cimgui.ColButtonHovered    // ImGuiCol_ButtonHovered
	ColButtonActive     = cimgui.ColButtonActive     // ImGuiCol_ButtonActive
	ColHeader           = cimgui.ColHeader           // ImGuiCol_Header
	ColHeaderHovered    = cimgui.ColHeaderHovered    // ImGuiCol_HeaderHovered
	ColHeaderActive     = cimgui.ColHeaderActive     // ImGuiCol_HeaderActive
	ColSeparator        = cimgui.ColSeparator        // ImGuiCol_Separator
	ColTab              = cimgui.ColTab              // ImGuiCol_Tab
	ColTabHovered       = cimgui.ColTabHovered       // ImGuiCol_TabHovered
	ColTabSelected      = cimgui.ColTabSelected      // ImGuiCol_TabSelected
	ColPlotLines        = cimgui.ColPlotLines        // ImGuiCol_PlotLines
	ColPlotHistogram    = cimgui.ColPlotHistogram    // ImGuiCol_PlotHistogram
	ColTextSelectedBg   = cimgui.ColTextSelectedBg   // ImGuiCol_TextSelectedBg
	ColModalWindowDimBg = cimgui.ColModalWindowDimBg // ImGuiCol_ModalWindowDimBg
)

// StyleVar identifies a styleable layout variable, modelling ImGuiStyleVar. It
// indexes the style-variable stack used by [PushStyleVar] and [StyleVarScope].
type StyleVar = cimgui.StyleVar

// Style-variable identifiers for [StyleVar].
const (
	StyleVarAlpha            = cimgui.StyleVarAlpha            // ImGuiStyleVar_Alpha
	StyleVarDisabledAlpha    = cimgui.StyleVarDisabledAlpha    // ImGuiStyleVar_DisabledAlpha
	StyleVarWindowPadding    = cimgui.StyleVarWindowPadding    // ImGuiStyleVar_WindowPadding
	StyleVarWindowRounding   = cimgui.StyleVarWindowRounding   // ImGuiStyleVar_WindowRounding
	StyleVarWindowBorderSize = cimgui.StyleVarWindowBorderSize // ImGuiStyleVar_WindowBorderSize
	StyleVarFramePadding     = cimgui.StyleVarFramePadding     // ImGuiStyleVar_FramePadding
	StyleVarFrameRounding    = cimgui.StyleVarFrameRounding    // ImGuiStyleVar_FrameRounding
	StyleVarFrameBorderSize  = cimgui.StyleVarFrameBorderSize  // ImGuiStyleVar_FrameBorderSize
	StyleVarItemSpacing      = cimgui.StyleVarItemSpacing      // ImGuiStyleVar_ItemSpacing
	StyleVarItemInnerSpacing = cimgui.StyleVarItemInnerSpacing // ImGuiStyleVar_ItemInnerSpacing
	StyleVarIndentSpacing    = cimgui.StyleVarIndentSpacing    // ImGuiStyleVar_IndentSpacing
	StyleVarScrollbarSize    = cimgui.StyleVarScrollbarSize    // ImGuiStyleVar_ScrollbarSize
	StyleVarGrabMinSize      = cimgui.StyleVarGrabMinSize      // ImGuiStyleVar_GrabMinSize
	StyleVarGrabRounding     = cimgui.StyleVarGrabRounding     // ImGuiStyleVar_GrabRounding
	StyleVarTabRounding      = cimgui.StyleVarTabRounding      // ImGuiStyleVar_TabRounding
)

// DataType identifies the element type for the scalar widgets ([InputScalar],
// [SliderScalar], [DragScalar] and their N variants). It models ImGuiDataType.
type DataType = cimgui.DataType

// Element types for [DataType].
const (
	DataTypeS8     = cimgui.DataTypeS8     // ImGuiDataType_S8
	DataTypeU8     = cimgui.DataTypeU8     // ImGuiDataType_U8
	DataTypeS16    = cimgui.DataTypeS16    // ImGuiDataType_S16
	DataTypeU16    = cimgui.DataTypeU16    // ImGuiDataType_U16
	DataTypeS32    = cimgui.DataTypeS32    // ImGuiDataType_S32
	DataTypeU32    = cimgui.DataTypeU32    // ImGuiDataType_U32
	DataTypeS64    = cimgui.DataTypeS64    // ImGuiDataType_S64
	DataTypeU64    = cimgui.DataTypeU64    // ImGuiDataType_U64
	DataTypeFloat  = cimgui.DataTypeFloat  // ImGuiDataType_Float
	DataTypeDouble = cimgui.DataTypeDouble // ImGuiDataType_Double
	DataTypeBool   = cimgui.DataTypeBool   // ImGuiDataType_Bool
)

// MouseButton identifies a mouse button, modelling ImGuiMouseButton. It selects
// the button queried by [IsItemClicked].
type MouseButton = cimgui.MouseButton

// Mouse buttons for [MouseButton].
const (
	MouseButtonLeft   = cimgui.MouseButtonLeft   // ImGuiMouseButton_Left
	MouseButtonRight  = cimgui.MouseButtonRight  // ImGuiMouseButton_Right
	MouseButtonMiddle = cimgui.MouseButtonMiddle // ImGuiMouseButton_Middle
)

// Key identifies a keyboard key, modelling ImGuiKey.
type Key = cimgui.Key

// Keys for [Key].
const (
	KeyUpArrow   = cimgui.KeyUpArrow   // ImGuiKey_UpArrow
	KeyDownArrow = cimgui.KeyDownArrow // ImGuiKey_DownArrow
)
