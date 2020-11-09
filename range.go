package cellwalker

import (
	"fmt"
	"math"
	"strings"
)

// Range is a boundary that cellwalker can move
type Range struct {
	top    int
	left   int
	bottom int
	right  int
}

func newRange(top, left, bottom, right int) *Range {
	return &Range{
		top:    top,
		left:   left,
		bottom: bottom,
		right:  right,
	}
}

// Within creates a new range
func Within(rangeID string) *Range {
	cellBoundary := strings.SplitN(rangeID, ":", 2)
	var topLeft *Cell
	var bottomRight *Cell
	if len(cellBoundary) == 1 {
		topLeft = newCellByID(rangeID)
		bottomRight = newCellByID(rangeID)
	} else {
		topLeft = newCellByID(cellBoundary[0])
		bottomRight = newCellByID(cellBoundary[1])
		minCol := int(math.Min(float64(topLeft.ColumnIndex()), float64(bottomRight.ColumnIndex())))
		maxCol := int(math.Max(float64(topLeft.ColumnIndex()), float64(bottomRight.ColumnIndex())))
		minRow := int(math.Min(float64(topLeft.RowIndex()), float64(bottomRight.RowIndex())))
		maxRow := int(math.Max(float64(topLeft.RowIndex()), float64(bottomRight.RowIndex())))
		topLeft = newCell(minCol, minRow)
		bottomRight = newCell(maxCol, maxRow)
	}
	return newRange(
		topLeft.RowIndex(),
		topLeft.ColumnIndex(),
		bottomRight.RowIndex(),
		bottomRight.ColumnIndex(),
	)
}

func (r *Range) String() string {
	topLeft := newCell(r.left, r.top)
	bottomRight := newCell(r.right, r.bottom)
	return fmt.Sprintf("%s:%s", topLeft.String(), bottomRight.String())
}

// At returns cell walker with range constraint
func (r *Range) At(cellID string) *CellWalker {
	cell := newCellByID(cellID)
	return newCellWalker(cell, r)
}

// Clone creates a copy of Range
func (r *Range) Clone() *Range {
	return newRange(r.top, r.left, r.bottom, r.right)
}

// LeftIndex returns column index of left boundary
func (r *Range) LeftIndex() int {
	return r.left
}

// TopIndex returns row index of top boundary
func (r *Range) TopIndex() int {
	return r.top
}

// RightIndex returns column index of right boundary
func (r *Range) RightIndex() int {
	return r.right
}

// BottomIndex returns row index of bottom boundary
func (r *Range) BottomIndex() int {
	return r.bottom
}
