package imgui

import "github.com/bitwizeshift/go-imgui/internal/cimgui"

// TableOptions are the optional inputs to [Table]. A nil *TableOptions uses Dear
// ImGui's defaults; each field maps to an ImGuiTableFlags_ bit.
type TableOptions struct {
	Resizable                  bool // ImGuiTableFlags_Resizable
	Reorderable                bool // ImGuiTableFlags_Reorderable
	Hideable                   bool // ImGuiTableFlags_Hideable
	Sortable                   bool // ImGuiTableFlags_Sortable
	NoSavedSettings            bool // ImGuiTableFlags_NoSavedSettings
	ContextMenuInBody          bool // ImGuiTableFlags_ContextMenuInBody
	RowBg                      bool // ImGuiTableFlags_RowBg
	BordersInnerH              bool // ImGuiTableFlags_BordersInnerH
	BordersOuterH              bool // ImGuiTableFlags_BordersOuterH
	BordersInnerV              bool // ImGuiTableFlags_BordersInnerV
	BordersOuterV              bool // ImGuiTableFlags_BordersOuterV
	BordersH                   bool // ImGuiTableFlags_BordersH (composite)
	BordersV                   bool // ImGuiTableFlags_BordersV (composite)
	BordersInner               bool // ImGuiTableFlags_BordersInner (composite)
	BordersOuter               bool // ImGuiTableFlags_BordersOuter (composite)
	Borders                    bool // ImGuiTableFlags_Borders (composite)
	NoBordersInBody            bool // ImGuiTableFlags_NoBordersInBody
	NoBordersInBodyUntilResize bool // ImGuiTableFlags_NoBordersInBodyUntilResize
	SizingFixedFit             bool // ImGuiTableFlags_SizingFixedFit
	SizingFixedSame            bool // ImGuiTableFlags_SizingFixedSame
	SizingStretchProp          bool // ImGuiTableFlags_SizingStretchProp
	SizingStretchSame          bool // ImGuiTableFlags_SizingStretchSame
	NoHostExtendX              bool // ImGuiTableFlags_NoHostExtendX
	NoHostExtendY              bool // ImGuiTableFlags_NoHostExtendY
	NoKeepColumnsVisible       bool // ImGuiTableFlags_NoKeepColumnsVisible
	PreciseWidths              bool // ImGuiTableFlags_PreciseWidths
	NoClip                     bool // ImGuiTableFlags_NoClip
	PadOuterX                  bool // ImGuiTableFlags_PadOuterX
	NoPadOuterX                bool // ImGuiTableFlags_NoPadOuterX
	NoPadInnerX                bool // ImGuiTableFlags_NoPadInnerX
	ScrollX                    bool // ImGuiTableFlags_ScrollX
	ScrollY                    bool // ImGuiTableFlags_ScrollY
	SortMulti                  bool // ImGuiTableFlags_SortMulti
	SortTristate               bool // ImGuiTableFlags_SortTristate
	HighlightHoveredColumn     bool // ImGuiTableFlags_HighlightHoveredColumn
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *TableOptions) flags() cimgui.TableFlags {
	if o == nil {
		return cimgui.TableFlagsNone
	}
	var f cimgui.TableFlags
	if o.Resizable {
		f |= cimgui.TableFlagsResizable
	}
	if o.Reorderable {
		f |= cimgui.TableFlagsReorderable
	}
	if o.Hideable {
		f |= cimgui.TableFlagsHideable
	}
	if o.Sortable {
		f |= cimgui.TableFlagsSortable
	}
	if o.NoSavedSettings {
		f |= cimgui.TableFlagsNoSavedSettings
	}
	if o.ContextMenuInBody {
		f |= cimgui.TableFlagsContextMenuInBody
	}
	if o.RowBg {
		f |= cimgui.TableFlagsRowBg
	}
	if o.BordersInnerH {
		f |= cimgui.TableFlagsBordersInnerH
	}
	if o.BordersOuterH {
		f |= cimgui.TableFlagsBordersOuterH
	}
	if o.BordersInnerV {
		f |= cimgui.TableFlagsBordersInnerV
	}
	if o.BordersOuterV {
		f |= cimgui.TableFlagsBordersOuterV
	}
	if o.BordersH {
		f |= cimgui.TableFlagsBordersH
	}
	if o.BordersV {
		f |= cimgui.TableFlagsBordersV
	}
	if o.BordersInner {
		f |= cimgui.TableFlagsBordersInner
	}
	if o.BordersOuter {
		f |= cimgui.TableFlagsBordersOuter
	}
	if o.Borders {
		f |= cimgui.TableFlagsBorders
	}
	if o.NoBordersInBody {
		f |= cimgui.TableFlagsNoBordersInBody
	}
	if o.NoBordersInBodyUntilResize {
		f |= cimgui.TableFlagsNoBordersInBodyUntilResize
	}
	if o.SizingFixedFit {
		f |= cimgui.TableFlagsSizingFixedFit
	}
	if o.SizingFixedSame {
		f |= cimgui.TableFlagsSizingFixedSame
	}
	if o.SizingStretchProp {
		f |= cimgui.TableFlagsSizingStretchProp
	}
	if o.SizingStretchSame {
		f |= cimgui.TableFlagsSizingStretchSame
	}
	if o.NoHostExtendX {
		f |= cimgui.TableFlagsNoHostExtendX
	}
	if o.NoHostExtendY {
		f |= cimgui.TableFlagsNoHostExtendY
	}
	if o.NoKeepColumnsVisible {
		f |= cimgui.TableFlagsNoKeepColumnsVisible
	}
	if o.PreciseWidths {
		f |= cimgui.TableFlagsPreciseWidths
	}
	if o.NoClip {
		f |= cimgui.TableFlagsNoClip
	}
	if o.PadOuterX {
		f |= cimgui.TableFlagsPadOuterX
	}
	if o.NoPadOuterX {
		f |= cimgui.TableFlagsNoPadOuterX
	}
	if o.NoPadInnerX {
		f |= cimgui.TableFlagsNoPadInnerX
	}
	if o.ScrollX {
		f |= cimgui.TableFlagsScrollX
	}
	if o.ScrollY {
		f |= cimgui.TableFlagsScrollY
	}
	if o.SortMulti {
		f |= cimgui.TableFlagsSortMulti
	}
	if o.SortTristate {
		f |= cimgui.TableFlagsSortTristate
	}
	if o.HighlightHoveredColumn {
		f |= cimgui.TableFlagsHighlightHoveredColumn
	}
	return f
}

// TableColumnOptions are the optional inputs to [TableSetupColumn]. A nil pointer
// uses Dear ImGui's defaults; each field maps to an ImGuiTableColumnFlags_ bit.
// The Is* fields are status flags reported by Dear ImGui and have no effect when
// set as input.
type TableColumnOptions struct {
	Disabled             bool // ImGuiTableColumnFlags_Disabled
	DefaultHide          bool // ImGuiTableColumnFlags_DefaultHide
	DefaultSort          bool // ImGuiTableColumnFlags_DefaultSort
	WidthStretch         bool // ImGuiTableColumnFlags_WidthStretch
	WidthFixed           bool // ImGuiTableColumnFlags_WidthFixed
	NoResize             bool // ImGuiTableColumnFlags_NoResize
	NoReorder            bool // ImGuiTableColumnFlags_NoReorder
	NoHide               bool // ImGuiTableColumnFlags_NoHide
	NoClip               bool // ImGuiTableColumnFlags_NoClip
	NoSort               bool // ImGuiTableColumnFlags_NoSort
	NoSortAscending      bool // ImGuiTableColumnFlags_NoSortAscending
	NoSortDescending     bool // ImGuiTableColumnFlags_NoSortDescending
	NoHeaderLabel        bool // ImGuiTableColumnFlags_NoHeaderLabel
	NoHeaderWidth        bool // ImGuiTableColumnFlags_NoHeaderWidth
	PreferSortAscending  bool // ImGuiTableColumnFlags_PreferSortAscending
	PreferSortDescending bool // ImGuiTableColumnFlags_PreferSortDescending
	IndentEnable         bool // ImGuiTableColumnFlags_IndentEnable
	IndentDisable        bool // ImGuiTableColumnFlags_IndentDisable
	AngledHeader         bool // ImGuiTableColumnFlags_AngledHeader
	IsEnabled            bool // ImGuiTableColumnFlags_IsEnabled (status)
	IsVisible            bool // ImGuiTableColumnFlags_IsVisible (status)
	IsSorted             bool // ImGuiTableColumnFlags_IsSorted (status)
	IsHovered            bool // ImGuiTableColumnFlags_IsHovered (status)
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *TableColumnOptions) flags() cimgui.TableColumnFlags {
	if o == nil {
		return cimgui.TableColumnFlagsNone
	}
	var f cimgui.TableColumnFlags
	if o.Disabled {
		f |= cimgui.TableColumnFlagsDisabled
	}
	if o.DefaultHide {
		f |= cimgui.TableColumnFlagsDefaultHide
	}
	if o.DefaultSort {
		f |= cimgui.TableColumnFlagsDefaultSort
	}
	if o.WidthStretch {
		f |= cimgui.TableColumnFlagsWidthStretch
	}
	if o.WidthFixed {
		f |= cimgui.TableColumnFlagsWidthFixed
	}
	if o.NoResize {
		f |= cimgui.TableColumnFlagsNoResize
	}
	if o.NoReorder {
		f |= cimgui.TableColumnFlagsNoReorder
	}
	if o.NoHide {
		f |= cimgui.TableColumnFlagsNoHide
	}
	if o.NoClip {
		f |= cimgui.TableColumnFlagsNoClip
	}
	if o.NoSort {
		f |= cimgui.TableColumnFlagsNoSort
	}
	if o.NoSortAscending {
		f |= cimgui.TableColumnFlagsNoSortAscending
	}
	if o.NoSortDescending {
		f |= cimgui.TableColumnFlagsNoSortDescending
	}
	if o.NoHeaderLabel {
		f |= cimgui.TableColumnFlagsNoHeaderLabel
	}
	if o.NoHeaderWidth {
		f |= cimgui.TableColumnFlagsNoHeaderWidth
	}
	if o.PreferSortAscending {
		f |= cimgui.TableColumnFlagsPreferSortAscending
	}
	if o.PreferSortDescending {
		f |= cimgui.TableColumnFlagsPreferSortDescending
	}
	if o.IndentEnable {
		f |= cimgui.TableColumnFlagsIndentEnable
	}
	if o.IndentDisable {
		f |= cimgui.TableColumnFlagsIndentDisable
	}
	if o.AngledHeader {
		f |= cimgui.TableColumnFlagsAngledHeader
	}
	if o.IsEnabled {
		f |= cimgui.TableColumnFlagsIsEnabled
	}
	if o.IsVisible {
		f |= cimgui.TableColumnFlagsIsVisible
	}
	if o.IsSorted {
		f |= cimgui.TableColumnFlagsIsSorted
	}
	if o.IsHovered {
		f |= cimgui.TableColumnFlagsIsHovered
	}
	return f
}

// TableRowOptions are the optional inputs to [TableNextRow]. A nil pointer uses
// Dear ImGui's defaults.
type TableRowOptions struct {
	Headers bool // ImGuiTableRowFlags_Headers
}

// flags folds o into the cimgui flag set. A nil receiver yields no flags.
func (o *TableRowOptions) flags() cimgui.TableRowFlags {
	if o == nil || !o.Headers {
		return cimgui.TableRowFlagsNone
	}
	return cimgui.TableRowFlagsHeaders
}

// Table begins a table with the given number of columns. A zero outerSize
// auto-fits. It models ImGui::BeginTable. open reports whether the table is
// visible; the returned [EndFunc] (ImGui::EndTable) ends it only when open.
func Table(strID string, columns int32, outerSize Vec2, innerWidth float32, opts *TableOptions) (open bool, end EndFunc) {
	open = cimgui.BeginTable(strID, columns, opts.flags(), outerSize, innerWidth)
	if !open {
		return open, func() {}
	}
	return open, cimgui.EndTable
}

// TableNextRow advances to the next row. A zero minRowHeight uses the default. It
// models ImGui::TableNextRow.
func TableNextRow(minRowHeight float32, opts *TableRowOptions) {
	cimgui.TableNextRow(opts.flags(), minRowHeight)
}

// TableNextColumn advances to the next column (wrapping to a new row as needed)
// and reports whether the column is visible. It models ImGui::TableNextColumn.
func TableNextColumn() bool {
	return cimgui.TableNextColumn()
}

// TableSetColumnIndex moves to the given column and reports whether it is
// visible. It models ImGui::TableSetColumnIndex.
func TableSetColumnIndex(columnN int32) bool {
	return cimgui.TableSetColumnIndex(columnN)
}

// TableSetupColumn declares a column. initWidthOrWeight is interpreted per the
// column's sizing flag; userID may be 0. It models ImGui::TableSetupColumn.
func TableSetupColumn(label string, initWidthOrWeight float32, userID uint32, opts *TableColumnOptions) {
	cimgui.TableSetupColumn(label, opts.flags(), initWidthOrWeight, userID)
}

// TableSetupScrollFreeze locks the given number of columns and rows so they stay
// visible while scrolling. It models ImGui::TableSetupScrollFreeze.
func TableSetupScrollFreeze(cols, rows int32) {
	cimgui.TableSetupScrollFreeze(cols, rows)
}

// TableHeadersRow submits a row of headers using the labels from
// [TableSetupColumn]. It models ImGui::TableHeadersRow.
func TableHeadersRow() {
	cimgui.TableHeadersRow()
}

// TableAngledHeadersRow submits an angled-text headers row. It models
// ImGui::TableAngledHeadersRow.
func TableAngledHeadersRow() {
	cimgui.TableAngledHeadersRow()
}

// TableHeader submits a single header cell with the given label. It models
// ImGui::TableHeader.
func TableHeader(label string) {
	cimgui.TableHeader(label)
}
