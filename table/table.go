// Package table provides tabular layout.
package table

import (
	"rodusek.dev/pkg/imgui"
	"rodusek.dev/pkg/imgui/internal/cimgui"
)

// Sizing selects a column's width policy. The zero value leaves the table's
// default sizing in effect.
type Sizing int

const (
	// SizingDefault uses the table's default column sizing.
	SizingDefault Sizing = iota
	// SizingStretch sizes the column to a proportion of the available width.
	SizingStretch
	// SizingFixed gives the column a fixed width.
	SizingFixed
)

func setTableFlag(f *cimgui.TableFlags, bit cimgui.TableFlags, on bool) {
	if on {
		*f |= bit
	} else {
		*f &^= bit
	}
}

func setColumnFlag(f *cimgui.TableColumnFlags, bit cimgui.TableColumnFlags, on bool) {
	if on {
		*f |= bit
	} else {
		*f &^= bit
	}
}

// Column describes a single table column.
type Column struct {
	Label string
	Width float32 // initial width or stretch weight; 0 = default
	flags cimgui.TableColumnFlags
}

// SetSizing selects the column's width policy.
func (c *Column) SetSizing(s Sizing) {
	c.flags &^= cimgui.TableColumnFlagsWidthStretch | cimgui.TableColumnFlagsWidthFixed
	switch s {
	case SizingStretch:
		c.flags |= cimgui.TableColumnFlagsWidthStretch
	case SizingFixed:
		c.flags |= cimgui.TableColumnFlagsWidthFixed
	}
}

// SetDefaultHidden hides the column by default (the user can still show it).
func (c *Column) SetDefaultHidden(on bool) {
	setColumnFlag(&c.flags, cimgui.TableColumnFlagsDefaultHide, on)
}

// SetResizable allows or prevents the user resizing the column (default allowed).
func (c *Column) SetResizable(on bool) {
	setColumnFlag(&c.flags, cimgui.TableColumnFlagsNoResize, !on)
}

// SetAngledHeader draws this column's header label rotated. The table also needs
// an angled-headers row, enabled automatically when any column requests it.
func (c *Column) SetAngledHeader(on bool) {
	setColumnFlag(&c.flags, cimgui.TableColumnFlagsAngledHeader, on)
}

// angled reports whether the column's header is drawn angled.
func (c *Column) angled() bool {
	return c.flags&cimgui.TableColumnFlagsAngledHeader != 0
}

// Table is a grid of cells. Columns describe the header; each row is a slice of
// cell widgets (one per column).
type Table struct {
	ID         string
	Size       imgui.Vec2
	Columns    []*Column
	Rows       [][]imgui.Widget
	FreezeCols int32 // leftmost columns kept fixed while scrolling horizontally
	FreezeRows int32 // topmost rows (incl. headers) kept fixed while scrolling vertically
	flags      cimgui.TableFlags
}

// New returns an empty table. ID must be unique. By default it draws borders and
// striped row backgrounds; use the setters to change this.
func New(id string) *Table {
	return &Table{ID: id, flags: cimgui.TableFlagsBorders | cimgui.TableFlagsRowBg}
}

// SetBorders draws or omits the table borders (default drawn).
func (t *Table) SetBorders(on bool) {
	setTableFlag(&t.flags, cimgui.TableFlagsBorders, on)
}

// SetRowBackground draws or omits the striped row backgrounds (default drawn).
func (t *Table) SetRowBackground(on bool) {
	setTableFlag(&t.flags, cimgui.TableFlagsRowBg, on)
}

// SetResizable allows the user to resize columns.
func (t *Table) SetResizable(on bool) {
	setTableFlag(&t.flags, cimgui.TableFlagsResizable, on)
}

// SetVerticalScroll gives the table its own vertical scroll region.
func (t *Table) SetVerticalScroll(on bool) {
	setTableFlag(&t.flags, cimgui.TableFlagsScrollY, on)
}

// SetSortable allows the user to sort by clicking column headers.
func (t *Table) SetSortable(on bool) {
	setTableFlag(&t.flags, cimgui.TableFlagsSortable, on)
}

// AddColumn appends a named column and returns it so it can be configured.
func (t *Table) AddColumn(label string) *Column {
	c := &Column{Label: label}
	t.Columns = append(t.Columns, c)
	return c
}

// AddRow appends a row of cell widgets.
func (t *Table) AddRow(cells ...imgui.Widget) { t.Rows = append(t.Rows, cells) }

// hasAngledHeaders reports whether any column requests an angled header.
func (t *Table) hasAngledHeaders() bool {
	for _, c := range t.Columns {
		if c.angled() {
			return true
		}
	}
	return false
}

// Display draws the table.
func (t *Table) Display() {
	cols := len(t.Columns)
	if cols == 0 {
		for _, r := range t.Rows {
			if len(r) > cols {
				cols = len(r)
			}
		}
	}
	if cols == 0 {
		return
	}
	if !cimgui.BeginTable(t.ID, int32(cols), t.flags, t.Size, 0) {
		return
	}
	for _, c := range t.Columns {
		cimgui.TableSetupColumn(c.Label, c.flags, c.Width, 0)
	}
	if t.FreezeCols > 0 || t.FreezeRows > 0 {
		cimgui.TableSetupScrollFreeze(t.FreezeCols, t.FreezeRows)
	}
	if len(t.Columns) > 0 {
		if t.hasAngledHeaders() {
			cimgui.TableAngledHeadersRow()
		}
		cimgui.TableHeadersRow()
	}
	for _, row := range t.Rows {
		cimgui.TableNextRow(0, 0)
		for _, cell := range row {
			cimgui.TableNextColumn()
			if cell != nil {
				cell.Display()
			}
		}
	}
	cimgui.EndTable()
}
