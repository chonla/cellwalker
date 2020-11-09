package cellwalker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSheetRange(t *testing.T) {
	result := Sheet().String()

	assert.Equal(t, "A1:XFD1048576", result)
}

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
