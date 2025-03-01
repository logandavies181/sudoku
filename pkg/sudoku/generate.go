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

func generateThreeBoxes() {
	cells = make([]cell.Cell, 81)
	for i := range cells {
		cells[i] = cell.New(0)
	}

	firstRow := generateRandomCandidates()
	for i, v := range firstRow {
		cells[i] = cell.New(v)
	}

	basicCheckCells()

	backtrackingSolve()

	PrintPuzzle()
}

func getFirstUnsolvedIndex() int {
	for i, v := range cells {
		if !v.Solved() {
			return i
		}
	}

	return -1
}

// returns true if no unsolved cells have no candidates.
// solves any cells that only have one candidate because why not.
func checkNoUnsolveableCells() bool {
	found := false
	foreachAllUnsolvedCells(func(id int) {
		candidates := cells[id].ListCandidates()
		numCandidates := len(candidates)
		switch numCandidates {
		case 0:
			found = true
		case 1:
			solveCellAs(id, candidates[0])
		}
	})

	return !found
}

func backtrackingSolve() {
	guessIndex := getFirstUnsolvedIndex()
	if guessIndex > 0 {
		cans := cells[guessIndex].ListCandidates()
		canOrder := rand.Perm(len(cans))
		for _, v := range canOrder {
			unsafeGuess(guessIndex, v)
			basicCheckCellsSeenBy(guessIndex)
			if checkNoUnsolveableCells() {
				backtrackingSolve()
			} else {
				
			}
		}
	}
}
