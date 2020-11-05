package cellwalker

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// RowsLimit https://support.microsoft.com/en-us/office/excel-specifications-and-limits-1672b34d-7043-467e-8e27-269d656771c3
const (
	RowsLimit    = 1048576
	ColumnsLimit = 16384
)

// CellWalker struct
type CellWalker struct {
	column int
	row    int
}

func newCellWalker(col int, row int) *CellWalker {
	if col > ColumnsLimit {
		col = ColumnsLimit
	}
	if row < 1 {
		row = 1
	} else {
		if row > RowsLimit {
			row = RowsLimit
		}
	}
	return &CellWalker{
		column: col,
		row:    row,
	}
}

// At initializes CellWalker by specify initial cell to start
func At(cellID string) *CellWalker {
	cleanCellID := strings.ToUpper(cellID)
	re := regexp.MustCompile(`^([A-Z]+)([0-9]*)$`)
	match := re.FindStringSubmatch(cleanCellID)

	col := match[1]
	row, err := strconv.ParseInt(fmt.Sprintf("0%s", match[2]), 10, 32)
	if err != nil || row == 0 {
		row = 1
	}

	return newCellWalker(ColumnNameToIndex(col), int(row))
}

// ColumnIndexToName converts column index to default excel name
func ColumnIndexToName(id int) string {
	name := ""
	dividend := id
	modulo := 0

	for dividend > 0 {
		modulo = (dividend - 1) % 26
		name = fmt.Sprintf("%c%s", rune(modulo+'A'), name)
		dividend = (dividend - modulo) / 26
	}
	return name
}

// ColumnNameToIndex converts default excel column name to index, 1-based index
// name must be uppercase start from A, B, C, ..., Z, AA, AB, ... ZZ, AAA, ..., ZZZ, ...
func ColumnNameToIndex(name string) int {
	index := 0
	for colCharIndex, colCharLen := 0, len(name); colCharIndex < colCharLen; colCharIndex++ {
		charNum := int(name[colCharIndex]-'A') + 1
		digitNum := (colCharLen - (colCharIndex + 1))
		columnWeight := int(math.Pow(26.0, float64(digitNum)))
		columnValue := charNum * columnWeight
		index += columnValue
	}
	return index
}

// String representation of Cell
func (c *CellWalker) String() string {
	return fmt.Sprintf("%s%d", ColumnIndexToName(c.column), c.row)
}

// Above to move up one row
func (c *CellWalker) Above() *CellWalker {
	rowAbove := c.row - 1
	if rowAbove < 1 {
		rowAbove = 1
	}
	return newCellWalker(c.column, rowAbove)
}

// Below to move down one row
func (c *CellWalker) Below() *CellWalker {
	rowBeneath := c.row + 1
	if rowBeneath > RowsLimit {
		rowBeneath = RowsLimit
	}
	return newCellWalker(c.column, rowBeneath)
}

// Right to move right one column
func (c *CellWalker) Right() *CellWalker {
	rowRight := c.column + 1
	if rowRight > ColumnsLimit {
		rowRight = ColumnsLimit
	}
	return newCellWalker(rowRight, c.row)
}

// Left to move left one column
func (c *CellWalker) Left() *CellWalker {
	rowLeft := c.column - 1
	if rowLeft < 1 {
		rowLeft = 1
	}
	return newCellWalker(rowLeft, c.row)
}

// CanMoveLeft determines if it is at the left most cell
func (c *CellWalker) CanMoveLeft() bool {
	return c.column > 1
}

// CanMoveRight determines if it is at the right most cell
func (c *CellWalker) CanMoveRight() bool {
	return c.column < ColumnsLimit
}

// CanMoveUp determines if it is at the up most cell
func (c *CellWalker) CanMoveUp() bool {
	return c.row > 1
}

// CanMoveDown determines if it is at the bottom most cell
func (c *CellWalker) CanMoveDown() bool {
	return c.row < RowsLimit
}

// Column jump to a given colName
func (c *CellWalker) Column(colName string) *CellWalker {
	colIndex := ColumnNameToIndex(colName)

	return newCellWalker(colIndex, c.row)
}

// Row jump to a given row
func (c *CellWalker) Row(row int) *CellWalker {
	return newCellWalker(c.column, row)
}
