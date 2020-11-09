package cellwalker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOneCellRange(t *testing.T) {
	result := Within("A1")

	assert.Equal(t, "A1:A1", result.String())
}

func TestCreateRange(t *testing.T) {
	result := Within("A1:Z8")

	assert.Equal(t, "A1:Z8", result.String())
}

func TestCreateWithBottomRightFirst(t *testing.T) {
	result := Within("Z8:A1")

	assert.Equal(t, "A1:Z8", result.String())
}

func TestCreateRangeWithTopRightFirst(t *testing.T) {
	result := Within("Z1:A8")

	assert.Equal(t, "A1:Z8", result.String())
}

func TestCreateRangeWithBottomLeftFirst(t *testing.T) {
	result := Within("A8:Z1")

	assert.Equal(t, "A1:Z8", result.String())
}

func TestGetLeftIndexOfRange(t *testing.T) {
	result := Within("B4:AB300")

	assert.Equal(t, result.LeftIndex(), 2)
}

func TestGetTopIndexOfRange(t *testing.T) {
	result := Within("B4:AB300")

	assert.Equal(t, result.TopIndex(), 4)
}

func TestGetRightIndexOfRange(t *testing.T) {
	result := Within("B4:AB300")

	assert.Equal(t, result.RightIndex(), 28)
}

func TestGetBottomIndexOfRange(t *testing.T) {
	result := Within("B4:AB300")

	assert.Equal(t, result.BottomIndex(), 300)
}

func TestCreateCellWalkerWithinRange(t *testing.T) {
	result := Within("A8:Z1").At("C2")

	assert.Equal(t, "C2", result.String())
}
