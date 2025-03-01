package sudoku

import (
	"testing"

	"github.com/logandavies181/sudoku/pkg/cell"
	"github.com/stretchr/testify/assert"
)

func TestGetCandidateCounts(t *testing.T) {
	p := Puzzle{cell.New(0)}
	actual := p.getCandidateCounts([]int{0})

	expected := intIntMap{
		1: 1,
		2: 1,
		3: 1,
		4: 1,
		5: 1,
		6: 1,
		7: 1,
		8: 1,
		9: 1,
	}

	assert.Equal(t, expected, actual)
}

func TestBasicCheckCellsSeenBy(t *testing.T) {
	p, err := InitializeFromFile("../../test/basiccheckseenby.txt")
	assert.NoError(t, err)

	assert.True(t, p[5].HasCandidate(5))

	p[0].Solve()

	p.basicCheckCellsSeenBy(0)

	assert.False(t, p[5].HasCandidate(5))
}
