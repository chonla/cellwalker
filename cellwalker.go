package cellwalker

// CellWalker struct
type CellWalker struct {
	position *Cell
	boundary *Range
}

func newCellWalker(cell *Cell, boundary *Range) *CellWalker {
	return &CellWalker{
		position: cell.Clone(),
		boundary: boundary.Clone(),
	}
}

// At initializes CellWalker by specify initial cell to start
func At(cellID string) *CellWalker {
	return newCellWalker(newCellByID(cellID), Sheet())
}

func (c *CellWalker) String() string {
	return c.position.String()
}

// Clone creates a clone of cellwalker
func (c *CellWalker) Clone() *CellWalker {
	return &CellWalker{
		position: c.position,
		boundary: c.boundary,
	}
}

// Above to move up one row
func (c *CellWalker) Above() *CellWalker {
	if c.CanMoveUp() {
		rowAbove := c.position.RowIndex() - 1
		return newCellWalker(newCell(c.position.ColumnIndex(), rowAbove), c.boundary)
	}
	return c.Clone()
}

// Below to move down one row
func (c *CellWalker) Below() *CellWalker {
	if c.CanMoveDown() {
		rowBeneath := c.position.RowIndex() + 1
		return newCellWalker(newCell(c.position.ColumnIndex(), rowBeneath), c.boundary)
	}
	return c.Clone()
}

// Right to move right one column
func (c *CellWalker) Right() *CellWalker {
	if c.CanMoveRight() {
		columnRight := c.position.ColumnIndex() + 1
		return newCellWalker(newCell(columnRight, c.position.RowIndex()), c.boundary)
	}
	return c.Clone()
}

// Left to move left one column
func (c *CellWalker) Left() *CellWalker {
	if c.CanMoveLeft() {
		columnLeft := c.position.ColumnIndex() - 1
		return newCellWalker(newCell(columnLeft, c.position.RowIndex()), c.boundary)
	}
	return c.Clone()
}

// LeftMost to move leftmost column
func (c *CellWalker) LeftMost() *CellWalker {
	if c.CanMoveLeft() {
		leftMost := c.boundary.LeftIndex()
		return newCellWalker(newCell(leftMost, c.position.RowIndex()), c.boundary)
	}
	return c.Clone()
}

// Rightmost to move rightmost column
func (c *CellWalker) RightMost() *CellWalker {
	if c.CanMoveRight() {
		rightMost := c.boundary.RightIndex()
		return newCellWalker(newCell(rightMost, c.position.RowIndex()), c.boundary)
	}
	return c.Clone()
}

// TopMost to move topmost column
func (c *CellWalker) TopMost() *CellWalker {
	if c.CanMoveUp() {
		topMost := c.boundary.TopIndex()
		return newCellWalker(newCell(c.position.ColumnIndex(), topMost), c.boundary)
	}
	return c.Clone()
}

// Bottommost to move bottommost column
func (c *CellWalker) BottomMost() *CellWalker {
	if c.CanMoveDown() {
		bottomMost := c.boundary.BottomIndex()
		return newCellWalker(newCell(c.position.ColumnIndex(), bottomMost), c.boundary)
	}
	return c.Clone()
}

// CanMoveLeft determines if it is at the left most cell
func (c *CellWalker) CanMoveLeft() bool {
	return c.position.ColumnIndex() > c.boundary.left
}

// CanMoveRight determines if it is at the right most cell
func (c *CellWalker) CanMoveRight() bool {
	return c.position.ColumnIndex() < c.boundary.right
}

// CanMoveUp determines if it is at the up most cell
func (c *CellWalker) CanMoveUp() bool {
	return c.position.RowIndex() > c.boundary.top
}

// CanMoveDown determines if it is at the bottom most cell
func (c *CellWalker) CanMoveDown() bool {
	return c.position.RowIndex() < c.boundary.bottom
}

// Column jumps to a given colName
func (c *CellWalker) Column(colName string) *CellWalker {
	colIndex := ColumnNameToIndex(colName)

	return newCellWalker(newCell(colIndex, c.position.RowIndex()), c.boundary)
}

// Row jumps to a given row
func (c *CellWalker) Row(row int) *CellWalker {
	return newCellWalker(newCell(c.position.ColumnIndex(), row), c.boundary)
}

// ColumnOffset returns a cell with a given offset distance to column
func (c *CellWalker) ColumnOffset(offset int) *CellWalker {
	return newCellWalker(newCell(c.position.ColumnIndex()+offset, c.position.RowIndex()), c.boundary)
}

// RowOffset returns a cell with a given offset distance to row
func (c *CellWalker) RowOffset(offset int) *CellWalker {
	return newCellWalker(newCell(c.position.ColumnIndex(), c.position.RowIndex()+offset), c.boundary)
}

// Tour traverses position to Right column first then first column of next row if hit the boundary edge.
// Return nil if cannot make a further move
func (c *CellWalker) Tour() *CellWalker {
	if c.CanMoveRight() {
		return c.Right()
	}
	if c.CanMoveDown() {
		return c.Below().LeftMost()
	}
	return nil
}

// IsAtTopBoundary determine if current position is at the top of range
func (c *CellWalker) IsAtTopBoundary() bool {
	return c.position.RowIndex() == c.boundary.TopIndex()
}

// IsAtBottomBoundary determine if current position is at the bottom of range
func (c *CellWalker) IsAtBottomBoundary() bool {
	return c.position.RowIndex() == c.boundary.BottomIndex()
}

// IsAtLeftBoundary determine if current position is at the left of range
func (c *CellWalker) IsAtLeftBoundary() bool {
	return c.position.ColumnIndex() == c.boundary.LeftIndex()
}

// IsAtRightBoundary determine if current position is at the right of range
func (c *CellWalker) IsAtRightBoundary() bool {
	return c.position.ColumnIndex() == c.boundary.RightIndex()
}

// ColumnIndex returns the current column number
func (c *CellWalker) ColumnIndex() int {
	return c.position.ColumnIndex()
}

// ColumnName returns the current column name
func (c *CellWalker) ColumnName() string {
	return ColumnIndexToName(c.position.ColumnIndex())
}

// RowIndex returns the current row number
func (c *CellWalker) RowIndex() int {
	return c.position.RowIndex()
}
