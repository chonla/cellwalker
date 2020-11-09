package cellwalker

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Cell represents a cell in Excel
type Cell struct {
	column int
	row    int
}

func newCell(col, row int) *Cell {
	if col > ColumnsLimit {
		col = ColumnsLimit
	} else {
		if col < 1 {
			col = 1
		}
	}
	if row < 1 {
		row = 1
	} else {
		if row > RowsLimit {
			row = RowsLimit
		}
	}
	return &Cell{
		column: col,
		row:    row,
	}
}

func newCellByID(cellID string) *Cell {
	cleanCellID := strings.ToUpper(cellID)
	re := regexp.MustCompile(`^([A-Z]+)([0-9]*)$`)
	match := re.FindStringSubmatch(cleanCellID)

	col := match[1]
	row, err := strconv.ParseInt(fmt.Sprintf("0%s", match[2]), 10, 32)
	if err != nil || row == 0 {
		row = 1
	}

	return newCell(ColumnNameToIndex(col), int(row))
}

// String representation of Cell
func (c *Cell) String() string {
	return fmt.Sprintf("%s%d", ColumnIndexToName(c.column), c.row)
}

// Clone creates a copy of Cell
func (c *Cell) Clone() *Cell {
	return newCell(c.column, c.row)
}

// ColumnIndex returns column number
func (c *Cell) ColumnIndex() int {
	return c.column
}

// RowIndex returns row number
func (c *Cell) RowIndex() int {
	return c.row
}
