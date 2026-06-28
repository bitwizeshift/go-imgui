package cimgui

// #include <stdlib.h>
// #include "cimgui.h"
import "C"

import "unsafe"

// TableFlags configures a table opened with [BeginTable]. Mirrors the public
// ImGuiTableFlags_.
type TableFlags int32

const (
	TableFlagsNone                       TableFlags = C.ImGuiTableFlags_None
	TableFlagsResizable                  TableFlags = C.ImGuiTableFlags_Resizable
	TableFlagsReorderable                TableFlags = C.ImGuiTableFlags_Reorderable
	TableFlagsHideable                   TableFlags = C.ImGuiTableFlags_Hideable
	TableFlagsSortable                   TableFlags = C.ImGuiTableFlags_Sortable
	TableFlagsNoSavedSettings            TableFlags = C.ImGuiTableFlags_NoSavedSettings
	TableFlagsContextMenuInBody          TableFlags = C.ImGuiTableFlags_ContextMenuInBody
	TableFlagsRowBg                      TableFlags = C.ImGuiTableFlags_RowBg
	TableFlagsBordersInnerH              TableFlags = C.ImGuiTableFlags_BordersInnerH
	TableFlagsBordersOuterH              TableFlags = C.ImGuiTableFlags_BordersOuterH
	TableFlagsBordersInnerV              TableFlags = C.ImGuiTableFlags_BordersInnerV
	TableFlagsBordersOuterV              TableFlags = C.ImGuiTableFlags_BordersOuterV
	TableFlagsBordersH                   TableFlags = C.ImGuiTableFlags_BordersH
	TableFlagsBordersV                   TableFlags = C.ImGuiTableFlags_BordersV
	TableFlagsBordersInner               TableFlags = C.ImGuiTableFlags_BordersInner
	TableFlagsBordersOuter               TableFlags = C.ImGuiTableFlags_BordersOuter
	TableFlagsBorders                    TableFlags = C.ImGuiTableFlags_Borders
	TableFlagsNoBordersInBody            TableFlags = C.ImGuiTableFlags_NoBordersInBody
	TableFlagsNoBordersInBodyUntilResize TableFlags = C.ImGuiTableFlags_NoBordersInBodyUntilResize
	TableFlagsSizingFixedFit             TableFlags = C.ImGuiTableFlags_SizingFixedFit
	TableFlagsSizingFixedSame            TableFlags = C.ImGuiTableFlags_SizingFixedSame
	TableFlagsSizingStretchProp          TableFlags = C.ImGuiTableFlags_SizingStretchProp
	TableFlagsSizingStretchSame          TableFlags = C.ImGuiTableFlags_SizingStretchSame
	TableFlagsNoHostExtendX              TableFlags = C.ImGuiTableFlags_NoHostExtendX
	TableFlagsNoHostExtendY              TableFlags = C.ImGuiTableFlags_NoHostExtendY
	TableFlagsNoKeepColumnsVisible       TableFlags = C.ImGuiTableFlags_NoKeepColumnsVisible
	TableFlagsPreciseWidths              TableFlags = C.ImGuiTableFlags_PreciseWidths
	TableFlagsNoClip                     TableFlags = C.ImGuiTableFlags_NoClip
	TableFlagsPadOuterX                  TableFlags = C.ImGuiTableFlags_PadOuterX
	TableFlagsNoPadOuterX                TableFlags = C.ImGuiTableFlags_NoPadOuterX
	TableFlagsNoPadInnerX                TableFlags = C.ImGuiTableFlags_NoPadInnerX
	TableFlagsScrollX                    TableFlags = C.ImGuiTableFlags_ScrollX
	TableFlagsScrollY                    TableFlags = C.ImGuiTableFlags_ScrollY
	TableFlagsSortMulti                  TableFlags = C.ImGuiTableFlags_SortMulti
	TableFlagsSortTristate               TableFlags = C.ImGuiTableFlags_SortTristate
	TableFlagsHighlightHoveredColumn     TableFlags = C.ImGuiTableFlags_HighlightHoveredColumn
)

// TableColumnFlags configures a column set up with [TableSetupColumn]. Mirrors
// the public ImGuiTableColumnFlags_.
type TableColumnFlags int32

const (
	TableColumnFlagsNone                 TableColumnFlags = C.ImGuiTableColumnFlags_None
	TableColumnFlagsDisabled             TableColumnFlags = C.ImGuiTableColumnFlags_Disabled
	TableColumnFlagsDefaultHide          TableColumnFlags = C.ImGuiTableColumnFlags_DefaultHide
	TableColumnFlagsDefaultSort          TableColumnFlags = C.ImGuiTableColumnFlags_DefaultSort
	TableColumnFlagsWidthStretch         TableColumnFlags = C.ImGuiTableColumnFlags_WidthStretch
	TableColumnFlagsWidthFixed           TableColumnFlags = C.ImGuiTableColumnFlags_WidthFixed
	TableColumnFlagsNoResize             TableColumnFlags = C.ImGuiTableColumnFlags_NoResize
	TableColumnFlagsNoReorder            TableColumnFlags = C.ImGuiTableColumnFlags_NoReorder
	TableColumnFlagsNoHide               TableColumnFlags = C.ImGuiTableColumnFlags_NoHide
	TableColumnFlagsNoClip               TableColumnFlags = C.ImGuiTableColumnFlags_NoClip
	TableColumnFlagsNoSort               TableColumnFlags = C.ImGuiTableColumnFlags_NoSort
	TableColumnFlagsNoSortAscending      TableColumnFlags = C.ImGuiTableColumnFlags_NoSortAscending
	TableColumnFlagsNoSortDescending     TableColumnFlags = C.ImGuiTableColumnFlags_NoSortDescending
	TableColumnFlagsNoHeaderLabel        TableColumnFlags = C.ImGuiTableColumnFlags_NoHeaderLabel
	TableColumnFlagsNoHeaderWidth        TableColumnFlags = C.ImGuiTableColumnFlags_NoHeaderWidth
	TableColumnFlagsPreferSortAscending  TableColumnFlags = C.ImGuiTableColumnFlags_PreferSortAscending
	TableColumnFlagsPreferSortDescending TableColumnFlags = C.ImGuiTableColumnFlags_PreferSortDescending
	TableColumnFlagsIndentEnable         TableColumnFlags = C.ImGuiTableColumnFlags_IndentEnable
	TableColumnFlagsIndentDisable        TableColumnFlags = C.ImGuiTableColumnFlags_IndentDisable
	TableColumnFlagsAngledHeader         TableColumnFlags = C.ImGuiTableColumnFlags_AngledHeader
	TableColumnFlagsIsEnabled            TableColumnFlags = C.ImGuiTableColumnFlags_IsEnabled
	TableColumnFlagsIsVisible            TableColumnFlags = C.ImGuiTableColumnFlags_IsVisible
	TableColumnFlagsIsSorted             TableColumnFlags = C.ImGuiTableColumnFlags_IsSorted
	TableColumnFlagsIsHovered            TableColumnFlags = C.ImGuiTableColumnFlags_IsHovered
)

// TableRowFlags configures a row started with [TableNextRow]. Mirrors
// ImGuiTableRowFlags_.
type TableRowFlags int32

const (
	TableRowFlagsNone    TableRowFlags = C.ImGuiTableRowFlags_None
	TableRowFlagsHeaders TableRowFlags = C.ImGuiTableRowFlags_Headers
)

// BeginTable opens a table with the given number of columns. A zero outerSize
// auto-fits. Call [EndTable] only if it returns true.
func BeginTable(strID string, columns int32, flags TableFlags, outerSize Vec2, innerWidth float32) bool {
	cid := C.CString(strID)
	defer C.free(unsafe.Pointer(cid))
	return bool(C.igBeginTable(cid, C.int(columns), C.ImGuiTableFlags(flags), outerSize.c(), C.float(innerWidth)))
}

// EndTable closes the table opened by [BeginTable].
func EndTable() { C.igEndTable() }

// TableNextRow advances to the next row. A zero minRowHeight uses the default.
func TableNextRow(rowFlags TableRowFlags, minRowHeight float32) {
	C.igTableNextRow(C.ImGuiTableRowFlags(rowFlags), C.float(minRowHeight))
}

// TableNextColumn advances to the next column (wrapping to a new row as needed)
// and reports whether the column is visible.
func TableNextColumn() bool { return bool(C.igTableNextColumn()) }

// TableSetColumnIndex moves to the given column and reports whether it is visible.
func TableSetColumnIndex(columnN int32) bool {
	return bool(C.igTableSetColumnIndex(C.int(columnN)))
}

// TableSetupColumn declares a column. initWidthOrWeight is interpreted per the
// column's sizing flag; userID may be 0.
func TableSetupColumn(label string, flags TableColumnFlags, initWidthOrWeight float32, userID uint32) {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	C.igTableSetupColumn(clabel, C.ImGuiTableColumnFlags(flags), C.float(initWidthOrWeight), C.ImGuiID(userID))
}

// TableSetupScrollFreeze locks the given number of columns and rows so they stay
// visible while scrolling.
func TableSetupScrollFreeze(cols, rows int32) {
	C.igTableSetupScrollFreeze(C.int(cols), C.int(rows))
}

// TableHeadersRow submits a row of headers using the labels from [TableSetupColumn].
func TableHeadersRow() { C.igTableHeadersRow() }

// TableAngledHeadersRow submits an angled-text headers row.
func TableAngledHeadersRow() { C.igTableAngledHeadersRow() }

// TableHeader submits a single header cell with the given label.
func TableHeader(label string) {
	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	C.igTableHeader(clabel)
}
