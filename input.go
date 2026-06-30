package imgui

import (
	"unsafe"

	"github.com/bitwizeshift/go-imgui/internal/cimgui"
)

// TextBuffer is a growable, NUL-terminated C-backed text buffer used by the
// resizable input-text wrappers ([InputTextResizable] and friends). Build one
// with [NewTextBuffer]; its memory is released by a finalizer or eagerly via
// [TextBuffer.Free].
type TextBuffer = cimgui.TextBuffer

// InputTextCallback receives an event during an input-text widget, modelling the
// C ImGuiInputTextCallback. It returns 0 in the common case.
type InputTextCallback = cimgui.InputTextCallback

// InputTextCallbackData is a view over the live ImGuiInputTextCallbackData passed
// to an [InputTextCallback]. It is valid only for the duration of the callback.
type InputTextCallbackData = cimgui.InputTextCallbackData

// NewTextBuffer returns a [TextBuffer] seeded with s.
func NewTextBuffer(s string) *TextBuffer {
	return cimgui.NewTextBuffer(s)
}

// InputOptions are the optional inputs to the Input* widgets. A nil
// *InputOptions uses Dear ImGui's defaults. Format overrides the printf-style
// display format of the numeric inputs (empty selects each widget's default)
// and is ignored by the text inputs. The remaining fields map to
// ImGuiInputTextFlags_ bits.
type InputOptions struct {
	Format              string // numeric display format; empty uses the widget default
	CharsDecimal        bool   // ImGuiInputTextFlags_CharsDecimal
	CharsHexadecimal    bool   // ImGuiInputTextFlags_CharsHexadecimal
	CharsScientific     bool   // ImGuiInputTextFlags_CharsScientific
	CharsUppercase      bool   // ImGuiInputTextFlags_CharsUppercase
	CharsNoBlank        bool   // ImGuiInputTextFlags_CharsNoBlank
	AllowTabInput       bool   // ImGuiInputTextFlags_AllowTabInput
	EnterReturnsTrue    bool   // ImGuiInputTextFlags_EnterReturnsTrue
	EscapeClearsAll     bool   // ImGuiInputTextFlags_EscapeClearsAll
	CtrlEnterForNewLine bool   // ImGuiInputTextFlags_CtrlEnterForNewLine
	ReadOnly            bool   // ImGuiInputTextFlags_ReadOnly
	Password            bool   // ImGuiInputTextFlags_Password
	AlwaysOverwrite     bool   // ImGuiInputTextFlags_AlwaysOverwrite
	AutoSelectAll       bool   // ImGuiInputTextFlags_AutoSelectAll
	ParseEmptyRefVal    bool   // ImGuiInputTextFlags_ParseEmptyRefVal
	DisplayEmptyRefVal  bool   // ImGuiInputTextFlags_DisplayEmptyRefVal
	NoHorizontalScroll  bool   // ImGuiInputTextFlags_NoHorizontalScroll
	NoUndoRedo          bool   // ImGuiInputTextFlags_NoUndoRedo
	ElideLeft           bool   // ImGuiInputTextFlags_ElideLeft
	CallbackCompletion  bool   // ImGuiInputTextFlags_CallbackCompletion
	CallbackHistory     bool   // ImGuiInputTextFlags_CallbackHistory
	CallbackAlways      bool   // ImGuiInputTextFlags_CallbackAlways
	CallbackCharFilter  bool   // ImGuiInputTextFlags_CallbackCharFilter
	CallbackResize      bool   // ImGuiInputTextFlags_CallbackResize
	CallbackEdit        bool   // ImGuiInputTextFlags_CallbackEdit
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *InputOptions) flags() cimgui.InputTextFlags {
	if o == nil {
		return cimgui.InputTextFlagsNone
	}
	var f cimgui.InputTextFlags
	if o.CharsDecimal {
		f |= cimgui.InputTextFlagsCharsDecimal
	}
	if o.CharsHexadecimal {
		f |= cimgui.InputTextFlagsCharsHexadecimal
	}
	if o.CharsScientific {
		f |= cimgui.InputTextFlagsCharsScientific
	}
	if o.CharsUppercase {
		f |= cimgui.InputTextFlagsCharsUppercase
	}
	if o.CharsNoBlank {
		f |= cimgui.InputTextFlagsCharsNoBlank
	}
	if o.AllowTabInput {
		f |= cimgui.InputTextFlagsAllowTabInput
	}
	if o.EnterReturnsTrue {
		f |= cimgui.InputTextFlagsEnterReturnsTrue
	}
	if o.EscapeClearsAll {
		f |= cimgui.InputTextFlagsEscapeClearsAll
	}
	if o.CtrlEnterForNewLine {
		f |= cimgui.InputTextFlagsCtrlEnterForNewLine
	}
	if o.ReadOnly {
		f |= cimgui.InputTextFlagsReadOnly
	}
	if o.Password {
		f |= cimgui.InputTextFlagsPassword
	}
	if o.AlwaysOverwrite {
		f |= cimgui.InputTextFlagsAlwaysOverwrite
	}
	if o.AutoSelectAll {
		f |= cimgui.InputTextFlagsAutoSelectAll
	}
	if o.ParseEmptyRefVal {
		f |= cimgui.InputTextFlagsParseEmptyRefVal
	}
	if o.DisplayEmptyRefVal {
		f |= cimgui.InputTextFlagsDisplayEmptyRefVal
	}
	if o.NoHorizontalScroll {
		f |= cimgui.InputTextFlagsNoHorizontalScroll
	}
	if o.NoUndoRedo {
		f |= cimgui.InputTextFlagsNoUndoRedo
	}
	if o.ElideLeft {
		f |= cimgui.InputTextFlagsElideLeft
	}
	if o.CallbackCompletion {
		f |= cimgui.InputTextFlagsCallbackCompletion
	}
	if o.CallbackHistory {
		f |= cimgui.InputTextFlagsCallbackHistory
	}
	if o.CallbackAlways {
		f |= cimgui.InputTextFlagsCallbackAlways
	}
	if o.CallbackCharFilter {
		f |= cimgui.InputTextFlagsCallbackCharFilter
	}
	if o.CallbackResize {
		f |= cimgui.InputTextFlagsCallbackResize
	}
	if o.CallbackEdit {
		f |= cimgui.InputTextFlagsCallbackEdit
	}
	return f
}

// format returns o.Format, or def when o is nil or Format is empty.
func (o *InputOptions) format(def string) string {
	if o == nil || o.Format == "" {
		return def
	}
	return o.Format
}

// InputText edits the NUL-terminated text held in buf (capacity len(buf)) and
// reports whether it changed. It models ImGui::InputText. For an automatically
// growing buffer, use [InputTextResizable].
func InputText(label string, buf []byte, opts *InputOptions) bool {
	return cimgui.InputText(label, buf, opts.flags())
}

// InputTextMultiline edits buf in a multi-line box of the given size. It models
// ImGui::InputTextMultiline.
func InputTextMultiline(label string, buf []byte, size Vec2, opts *InputOptions) bool {
	return cimgui.InputTextMultiline(label, buf, size, opts.flags())
}

// InputTextWithHint edits buf, showing hint while empty. It models
// ImGui::InputTextWithHint.
func InputTextWithHint(label, hint string, buf []byte, opts *InputOptions) bool {
	return cimgui.InputTextWithHint(label, hint, buf, opts.flags())
}

// InputTextResizable edits buf, growing it automatically as text is entered, and
// reports whether it changed. cb, when non-nil, receives the callback events
// enabled in opts (the resize event is always handled internally). It models
// ImGui::InputText with a resize callback.
func InputTextResizable(label string, buf *TextBuffer, cb InputTextCallback, opts *InputOptions) bool {
	return cimgui.InputTextResizable(label, buf, opts.flags(), cb)
}

// InputTextMultilineResizable is [InputTextResizable] in a multi-line box of the
// given size. It models ImGui::InputTextMultiline with a resize callback.
func InputTextMultilineResizable(label string, buf *TextBuffer, size Vec2, cb InputTextCallback, opts *InputOptions) bool {
	return cimgui.InputTextMultilineResizable(label, buf, size, opts.flags(), cb)
}

// InputTextWithHintResizable is [InputTextResizable] showing hint while empty. It
// models ImGui::InputTextWithHint with a resize callback.
func InputTextWithHintResizable(label, hint string, buf *TextBuffer, cb InputTextCallback, opts *InputOptions) bool {
	return cimgui.InputTextWithHintResizable(label, hint, buf, opts.flags(), cb)
}

// InputFloat edits a float in a box with optional step buttons (a zero step
// hides them). It models ImGui::InputFloat.
func InputFloat(label string, v *float32, step, stepFast float32, opts *InputOptions) bool {
	return cimgui.InputFloat(label, v, step, stepFast, opts.format("%.3f"), opts.flags())
}

// InputFloat2 edits a 2-component float bound to v. It models ImGui::InputFloat2.
func InputFloat2(label string, v *[2]float32, opts *InputOptions) bool {
	return cimgui.InputFloat2(label, v, opts.format("%.3f"), opts.flags())
}

// InputFloat3 edits a 3-component float bound to v. It models ImGui::InputFloat3.
func InputFloat3(label string, v *[3]float32, opts *InputOptions) bool {
	return cimgui.InputFloat3(label, v, opts.format("%.3f"), opts.flags())
}

// InputFloat4 edits a 4-component float bound to v. It models ImGui::InputFloat4.
func InputFloat4(label string, v *[4]float32, opts *InputOptions) bool {
	return cimgui.InputFloat4(label, v, opts.format("%.3f"), opts.flags())
}

// InputInt edits an int in a box with optional step buttons (a zero step hides
// them). It models ImGui::InputInt.
func InputInt(label string, v *int32, step, stepFast int32, opts *InputOptions) bool {
	return cimgui.InputInt(label, v, step, stepFast, opts.flags())
}

// InputInt2 edits a 2-component int bound to v. It models ImGui::InputInt2.
func InputInt2(label string, v *[2]int32, opts *InputOptions) bool {
	return cimgui.InputInt2(label, v, opts.flags())
}

// InputInt3 edits a 3-component int bound to v. It models ImGui::InputInt3.
func InputInt3(label string, v *[3]int32, opts *InputOptions) bool {
	return cimgui.InputInt3(label, v, opts.flags())
}

// InputInt4 edits a 4-component int bound to v. It models ImGui::InputInt4.
func InputInt4(label string, v *[4]int32, opts *InputOptions) bool {
	return cimgui.InputInt4(label, v, opts.flags())
}

// InputDouble edits a double in a box with optional step buttons. It models
// ImGui::InputDouble.
func InputDouble(label string, v *float64, step, stepFast float64, opts *InputOptions) bool {
	return cimgui.InputDouble(label, v, step, stepFast, opts.format("%.6f"), opts.flags())
}

// InputScalar edits a single value of an arbitrary data type. pData, pStep and
// pStepFast point to values of dataType; pStep and pStepFast may be nil. It
// models ImGui::InputScalar.
func InputScalar(label string, dataType DataType, pData, pStep, pStepFast unsafe.Pointer, opts *InputOptions) bool {
	return cimgui.InputScalar(label, dataType, pData, pStep, pStepFast, opts.format(""), opts.flags())
}

// InputScalarN edits components values of dataType stored contiguously at pData.
// pStep and pStepFast may be nil. It models ImGui::InputScalarN.
func InputScalarN(label string, dataType DataType, pData unsafe.Pointer, components int32, pStep, pStepFast unsafe.Pointer, opts *InputOptions) bool {
	return cimgui.InputScalarN(label, dataType, pData, components, pStep, pStepFast, opts.format(""), opts.flags())
}
