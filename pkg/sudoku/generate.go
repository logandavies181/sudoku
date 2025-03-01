package sudoku

import (
	"math/rand"

	"github.com/logandavies181/sudoku/pkg/cell"
)

func generateRandomCandidates() []int {
	numbers := rand.Perm(9)
	for i := range numbers {
		numbers[i]++
	}
	return numbers
}

func newCompletedPuzzle() Puzzle {
	cells := make(Puzzle, 81)
	for i := range cells {
		cells[i] = cell.New(0)
	}

	firstRow := generateRandomCandidates()
	for i, v := range firstRow {
		cells[i] = cell.New(v)
	}

	cells.basicCheckCells()

	cells.backtrackingSolve()

	cells.Print()

	return cells
}

func (p Puzzle) getFirstUnsolvedIndex() int {
	for i, v := range p {
		if !v.Solved() {
			return i
		}
	}

	return -1
}

// returns true if no unsolved cells have no candidates.
// solves any cells that only have one candidate because why not.
func (p Puzzle) checkNoUnsolveableCells() bool {
	found := false
	p.foreachAllUnsolvedCells(func(id int) {
		candidates := p[id].ListCandidates()
		numCandidates := len(candidates)
		switch numCandidates {
		case 0:
			found = true
		case 1:
			p.solveCellAs(id, candidates[0])
		}
	})

	return !found
}

func (p Puzzle) backtrackingSolve() {
	guessIndex := p.getFirstUnsolvedIndex()
	if guessIndex > 0 {
		cans := p[guessIndex].ListCandidates()
		canOrder := rand.Perm(len(cans))
		for _, v := range canOrder {
			p.unsafeGuess(guessIndex, v)
			p.basicCheckCellsSeenBy(guessIndex)
			if p.checkNoUnsolveableCells() {
				p.backtrackingSolve()
			} else {
				
			}
		}
	}
}
