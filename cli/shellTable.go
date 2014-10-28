package cli

import (
	"bytes"
	"strings"
)

func padder(size int) string {
	var pads []string
	for i := 0; i < size; i++ {
		pads = append(pads, " ")
	}
	return strings.Join(pads, "")
}

// ShellTable is a table to be rendered in the shell.
type ShellTable struct {
	separator       *string
	maxColumnWidths *[]int
	rows            []*ShellTableRow
}

// NewShellTable initialized and returns a new ShellTable
func NewShellTable(sep string) *ShellTable {
	return &ShellTable{separator: &sep}
}

// NewSharedShellTable initialized and returns a new ShellTable with shared ptrs
func NewSharedShellTable(sep *string, widPtr *[]int) *ShellTable {
	return &ShellTable{separator: sep, maxColumnWidths: widPtr}
}

// Row adds a new row to the shell table
func (t *ShellTable) Row() *ShellTableRow {
	newRow := &ShellTableRow{}
	t.rows = append(t.rows, newRow)
	return newRow
}

// String returns the string table
func (t *ShellTable) padded() *ShellTable {
	maxColumnWidths := t.MaxColumnWidths()
	for _, r := range t.rows {
		for i, c := range r.columns {
			c.WriteString(padder(maxColumnWidths[i] - len(c.String())))
		}
	}
	return t
}

func (t *ShellTable) String() string {
	var lines []string
	for _, r := range t.padded().rows {
		var cols []string
		for _, c := range r.columns {
			cols = append(cols, c.String())
		}
		lines = append(lines, strings.Join(cols, *t.separator))
	}
	return strings.Join(lines, "\n")
}

// ShellTableRow is a row
type ShellTableRow struct {
	columns []*ShellTableColumn
}

// MaxColumnWidths returns an array of max column widths for each column
func (t *ShellTable) MaxColumnWidths() []int {
	if t.maxColumnWidths == nil {
		t.maxColumnWidths = &[]int{}
	}
	for _, r := range t.rows {
		for i, c := range r.columns {
			if len(*t.maxColumnWidths) == i {
				*t.maxColumnWidths = append(*t.maxColumnWidths, 0)
			}
			colWidth := len(c.String())
			if colWidth > (*t.maxColumnWidths)[i] {
				(*t.maxColumnWidths)[i] = colWidth
			}
		}
	}
	return *t.maxColumnWidths
}

// Column creates a new column
func (r *ShellTableRow) Column(strs ...string) *ShellTableColumn {
	newCol := &ShellTableColumn{new(bytes.Buffer)}
	r.columns = append(r.columns, newCol)
	newCol.WriteString(strings.Join(strs, " "))
	return newCol
}

// ShellTableColumn is a buffer for a column
type ShellTableColumn struct {
	*bytes.Buffer
}
