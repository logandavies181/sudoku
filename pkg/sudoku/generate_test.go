package sudoku

import (
	"math/rand"
	"testing"

	"github.com/logandavies181/sudoku/pkg/cell"
	"github.com/stretchr/testify/assert"
)

func TestNewCompletedPuzzle(t *testing.T) {
	p := newCompletedPuzzle()
	ind := rand.New(rand.NewSource(1)).Int() % 9
	p[ind] = cell.New(0)
	p[80-ind] = cell.New(0)
	p.Print()
}

func TestNewPuzzle(t *testing.T) {
	p := newPuzzle()
	assert.NotNil(t, p)

	//p.Print()
}
