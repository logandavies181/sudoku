package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzleSetHas_Empty(t *testing.T) {
	p := newEmptyPuzzle()
	ps := puzzleSet{p}

	assert.True(t, ps.has(p))
}

func TestPuzzleSetHas_Complete(t *testing.T) {
	p := newCompletedPuzzle()
	ps := puzzleSet{p}

	assert.True(t, ps.has(p))
	assert.False(t, ps.has(newEmptyPuzzle()))
}

func TestPuzzleSetHas_Multiple(t *testing.T) {
	p := newCompletedPuzzle()
	q := newEmptyPuzzle()
	ps := puzzleSet{p, q}

	assert.True(t, ps.has(p))
}

func TestPuzzleSetAdd(t *testing.T) {
	p := newCompletedPuzzle()
	q := newEmptyPuzzle()
	ps := puzzleSet{}

	ps = ps.add(p).add(q)
	qs := ps.add(p)

	assert.Equal(t, ps, qs)
}
