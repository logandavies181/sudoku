package sudoku

import (
	"fmt"
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

	cells, _ = cells.backtrackingSolve()

	cells.Print()

	return cells
}

func (p Puzzle) clone() Puzzle {
	q := make(Puzzle, 81)
	for i, v := range p {
		q[i] = v
	}
	return q
}

// gets first unsolved index. returns -1 if none are found,
// or -2 if an unsolveable index is found
func (p Puzzle) getFirstUnsolvedIndex() int {
	for i, v := range p {
		solved, err := v.Solved()
		if err != nil {
			return -2
		}
		if !solved {
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
			p.basicCheckCellsSeenBy(id)
		}
	})

	return !found
}

func (p Puzzle) backtrackingSolve() (Puzzle, bool) {
	defer func() {
		if r := recover(); r != nil {
			p.Print()
			fmt.Println(r)
		}
	}()

	guessIndex := p.getFirstUnsolvedIndex()
	switch guessIndex {
	case -1:
		return p, true
	case -2:
		return nil, false
	}
	if guessIndex > 0 {
		cans := p[guessIndex].ListCandidates()
		canOrder := rand.Perm(len(cans))
		for _, v := range canOrder {
			q := p.clone()
			can := cans[v]
			q.unsafeGuess(guessIndex, can)
			q.basicCheckCellsSeenBy(guessIndex)
			if q.checkNoUnsolveableCells() {
				if solution, solved := q.backtrackingSolve(); solved {
					return solution, true
				}
			}
		}
	}

	return nil, false
}
