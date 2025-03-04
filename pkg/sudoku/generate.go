package sudoku

import (
	"fmt"
	"math/rand"
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

func newPuzzle() Puzzle {
	p := newCompletedPuzzle()
	removeOrder := rand.Perm(40)

	for _, v := range removeOrder {
		for _, w := range []int{v, 80 - v} {
			q := p.cloneWithoutCandidates()

			q[w] = cell.New(0)
			q.basicCheckCells()
			q.Print()
			_, solved, numSolns := q.backtrackingSolve(true)
			if !solved {
				panic("couldn't backtrack solve newPuzzle")
			}

			if numSolns > 1 {
				return p
			}

			p = q.clone()
		}
	}

	return nil
}
