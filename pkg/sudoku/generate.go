package sudoku

import (
	"fmt"
	"math/rand"
	"time"

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

	start := time.Now()
	cells, _ = cells.BacktrackingSolve()

	cells.Print()

	fmt.Println(time.Now().Sub(start))

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

func (p Puzzle) BacktrackingSolve() (Puzzle, error) {
	p.basicCheckCells()
	p, ok := p.backtrackingSolve()
	if !ok {
		p.Print()
		return nil, fmt.Errorf("could not solve with backtracking")
	}
	return p, nil
}

func (p Puzzle) backtrackingSolve() (Puzzle, bool) {
	guessIndex := p.getFirstUnsolvedIndex()
	switch guessIndex {
	case -1:
		return p, true
	case -2:
		return nil, false
	}

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

	return nil, false
}
