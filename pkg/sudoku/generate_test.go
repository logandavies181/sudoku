package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPuzzle(t *testing.T) {
	p := NewPuzzle(25, 10_000)
	assert.NotNil(t, p)

	p.Print()

	assert.Equal(t, 25, p.clueCount())
}
