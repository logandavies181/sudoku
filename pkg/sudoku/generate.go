package sudoku

import (
	"fmt"
	"time"

	"github.com/logandavies181/sudoku/pkg/cell"
)

func newCompletedPuzzle() Puzzle {
	cells := make(Puzzle, 81)
	for i := range cells {
		cells[i] = cell.New(0)
	}

	firstRow := generateRandomCandidates()
	for i, v := range firstRow {
		cells[i] = cell.New(v)
	}

	start := time.Now()
	cells, _ = cells.BacktrackingSolve()

	cells.Print()

	fmt.Println(time.Now().Sub(start))

	return cells
}
