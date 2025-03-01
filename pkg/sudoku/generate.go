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

	// TODO: this is inefficient because we're checking cells
	// that we don't intend to update.
	basicCheckCells()

	rows := 2
	for rows <= 3{ 
		count := 0
		offset := 9*(rows-1)
		secondRowNums := generateRandomCandidates()
		for _, num := range secondRowNums {
			for i := range cells[offset:offset+9] {
				if cells[i+offset].HasCandidate(num) {
					cells[i+offset].SolveAs(num)
					count++
					break
				}
			}
		}
		if count == offset {
			rows++
			continue
		}
		for i := range offset {
			cells[i+offset] = cell.New(0)
		}
		basicCheckCells()
	}

	PrintPuzzle()
}
