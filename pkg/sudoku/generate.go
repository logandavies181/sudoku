package sudoku

import (
	"math/rand"

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

	cells, _ = cells.BacktrackingSolve()

	return cells
}

func NewPuzzle(targetClueCount, maxTries int) Puzzle {
	p := newCompletedPuzzle()

	for range maxTries {
		removeOrder := rand.Perm(40)

		var soln Puzzle
		for _, v := range removeOrder {
			for i, w := range []int{v, 80 - v} {
				q := p.cloneWithoutCandidates()

				q[w] = cell.New(0)
				q.basicCheckCells()
				_, solved, numSolns := q.backtrackingSolve(true)
				if !solved {
					panic("couldn't backtrack solve newPuzzle")
				}

				if numSolns > 1 && q.clueCount() < targetClueCount {
					// Ensure the puzzle is always symmetrical by requiring that an even number of
					// iterations has happened. The centre cell gets removed twice, so it's still
					// an even number of iterations if removed.
					if (i == 1) {
						return p
					} else {
						return soln
					}
				}

				p = q.clone()
				if (i == 1) {
					soln = p.clone()
				}
			}
		}
	}

	return nil
}

func (p Puzzle) clueCount() int {
	count := 0
	for _, v := range p {
		if v.Value != 0 {
			count++
		}
	}
	return count
}
