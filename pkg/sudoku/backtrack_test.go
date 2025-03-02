package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePuzzle(t *testing.T) {
	p := newCompletedPuzzle()

	assert.NotNil(t, p)
	assert.NoError(t, p.Validate())
	assert.Equal(t, -1, p.getFirstUnsolvedIndex())
}

func TestClone(t *testing.T) {
	p, err := InitializeFromFile("../../test/medium.txt")
	assert.NoError(t, err)

	q := p.clone()
	assert.Equal(t, p[0].ListCandidates(), q[0].ListCandidates())

	q.solveCellAs(0, 5)

	assert.NotEqual(t, p[0], q[0])
}
