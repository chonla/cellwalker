package cellwalker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertColumnIndexToSingleColumnName(t *testing.T) {
	result := ColumnIndexToName(1)

	assert.Equal(t, "A", result)
}

func TestConvertLastSingleColumnIndexToSingleColumnName(t *testing.T) {
	result := ColumnIndexToName(26)

	assert.Equal(t, "Z", result)
}

func TestConvertColumnIndexTo2CharactersColumnName(t *testing.T) {
	result := ColumnIndexToName(27)

	assert.Equal(t, "AA", result)
}

func TestConvertLastColumnIndexTo2CharactersColumnName(t *testing.T) {
	result := ColumnIndexToName(52)

	assert.Equal(t, "AZ", result)
}

func TestConvertComplexLastColumnIndexTo2CharactersColumnName(t *testing.T) {
	result := ColumnIndexToName(702)

	assert.Equal(t, "ZZ", result)
}

func TestColumnIndexTo3CharactersColumnName(t *testing.T) {
	result := ColumnIndexToName(703)

	assert.Equal(t, "AAA", result)
}

func TestColumnNameToSingleDigitIndex(t *testing.T) {
	result := ColumnNameToIndex("A")

	assert.Equal(t, 1, result)
}

func TestLast1CharColumnNameToIndex(t *testing.T) {
	result := ColumnNameToIndex("Z")

	assert.Equal(t, 26, result)
}

func TestFirst2CharsColumnNameToIndex(t *testing.T) {
	result := ColumnNameToIndex("AA")

	assert.Equal(t, 27, result)
}

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
