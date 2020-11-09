package cellwalker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCell(t *testing.T) {
	result := newCell(1, 1).String()

	assert.Equal(t, result, "A1")
}
