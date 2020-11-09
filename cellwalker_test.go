package cellwalker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpecifyCellWithCellID(t *testing.T) {
	result := At("A1").String()

	assert.Equal(t, "A1", result)
}

func TestSpecifyCellWithOutOfBoundColumnShouldMoveToTheEdge(t *testing.T) {
	result := At("XFE1").String()

	assert.Equal(t, "XFD1", result)
}

func TestSpecifyCellWithRowOnly(t *testing.T) {
	result := At("B").String()

	assert.Equal(t, "B1", result)
}

func TestSpecifyCellWithLargeRowOnly(t *testing.T) {
	result := At("CCC").String()

	assert.Equal(t, "CCC1", result)
}

func TestMoveCurrentCellUpward(t *testing.T) {
	result := At("A2").Above().String()

	assert.Equal(t, "A1", result)
}

func TestMoveTopMostCellUpwardShouldGoNowhere(t *testing.T) {
	result := At("A1").Above().String()

	assert.Equal(t, "A1", result)
}

func TestMoveCurrentCellDownward(t *testing.T) {
	result := At("A2").Below().String()

	assert.Equal(t, "A3", result)
}

func TestMoveBottomMostCellDownwardShouldGoNowhere(t *testing.T) {
	result := At("A1048576").Below().String()

	assert.Equal(t, "A1048576", result)
}

func TestMoveCurrentCellRightward(t *testing.T) {
	result := At("A1").Right().String()

	assert.Equal(t, "B1", result)
}

func TestMoveRightMostCellRightwardShouldGoNowhere(t *testing.T) {
	result := At("XFD1").Right().String()

	assert.Equal(t, "XFD1", result)
}

func TestMoveCurrentCellLeftward(t *testing.T) {
	result := At("B1").Left().String()

	assert.Equal(t, "A1", result)
}

func TestMoveLeftMostCellLeftwardShouldGoNowhere(t *testing.T) {
	result := At("A1").Left().String()

	assert.Equal(t, "A1", result)
}

func TestCanMoveLeftShouldReturnTrueIfNotAtTheLeftMostCell(t *testing.T) {
	result := At("B1").CanMoveLeft()

	assert.True(t, result)
}

func TestCanMoveLeftShouldReturnFalseIfAtTheLeftMostCell(t *testing.T) {
	result := At("A1").CanMoveLeft()

	assert.False(t, result)
}

func TestCanMoveRightShouldReturnTrueIfNotAtTheRightMostCell(t *testing.T) {
	result := At("XFC1").CanMoveRight()

	assert.True(t, result)
}

func TestCanMoveRightShouldReturnFalseIfAtTheRightMostCell(t *testing.T) {
	result := At("XFD1").CanMoveRight()

	assert.False(t, result)
}

func TestCanMoveUpShouldReturnTrueIfNotAtTheTopMostCell(t *testing.T) {
	result := At("A2").CanMoveUp()

	assert.True(t, result)
}

func TestCanMoveUpShouldReturnFalseIfAtTheTopMostCell(t *testing.T) {
	result := At("A1").CanMoveUp()

	assert.False(t, result)
}

func TestCanMoveDownShouldReturnTrueIfNotAtTheDownMostCell(t *testing.T) {
	result := At("A1048575").CanMoveDown()

	assert.True(t, result)
}

func TestCanMoveDownShouldReturnFalseIfAtTheDownMostCell(t *testing.T) {
	result := At("A1048576").CanMoveDown()

	assert.False(t, result)
}

func TestJumpToAnyColumnInTheSameRow(t *testing.T) {
	result := At("A1").Column("X")

	assert.Equal(t, "X1", result.String())
}

func TestJumpToAnyRowInTheSameColumn(t *testing.T) {
	result := At("A1").Row(778)

	assert.Equal(t, "A778", result.String())
}

func TestPositiveColumnOffset(t *testing.T) {
	result := At("A1").ColumnOffset(4)

	assert.Equal(t, "E1", result.String())
}

func TestNegativeColumnOffset(t *testing.T) {
	result := At("AB1").ColumnOffset(-4)

	assert.Equal(t, "X1", result.String())
}

func TestVeryLargePositiveColumnOffset(t *testing.T) {
	result := At("A1").ColumnOffset(40000000)

	assert.Equal(t, "XFD1", result.String())
}

func TestVeryLargeNegativeColumnOffset(t *testing.T) {
	result := At("XFD1").ColumnOffset(-40000000)

	assert.Equal(t, "A1", result.String())
}

func TestPositiveRowOffset(t *testing.T) {
	result := At("A1").RowOffset(4)

	assert.Equal(t, "A5", result.String())
}

func TestNegativeRowOffset(t *testing.T) {
	result := At("AB10").RowOffset(-4)

	assert.Equal(t, "AB6", result.String())
}

func TestVeryLargePositiveRowOffset(t *testing.T) {
	result := At("A1").RowOffset(40000000)

	assert.Equal(t, "A1048576", result.String())
}

func TestVeryLargeNegativeRowOffset(t *testing.T) {
	result := At("A1048576").RowOffset(-40000000)

	assert.Equal(t, "A1", result.String())
}

func TestTour(t *testing.T) {
	result1 := Within("B3:E5").At("C4")
	result2 := result1.Tour()
	result3 := result2.Tour()
	result4 := result3.Tour()
	result5 := result4.Tour()
	result6 := result5.Tour()
	result7 := result6.Tour()
	result8 := result7.Tour()

	assert.NotNil(t, result1)
	assert.Equal(t, "C4", result1.String())
	assert.NotNil(t, result2)
	assert.Equal(t, "D4", result2.String())
	assert.NotNil(t, result3)
	assert.Equal(t, "E4", result3.String())
	assert.NotNil(t, result4)
	assert.Equal(t, "B5", result4.String())
	assert.NotNil(t, result5)
	assert.Equal(t, "C5", result5.String())
	assert.NotNil(t, result6)
	assert.Equal(t, "D5", result6.String())
	assert.NotNil(t, result7)
	assert.Equal(t, "E5", result7.String())
	assert.Nil(t, result8)
}
