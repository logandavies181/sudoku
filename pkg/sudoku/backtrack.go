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

func (p Puzzle) clone() Puzzle {
	q := make(Puzzle, 81)
	for i, v := range p {
		q[i] = v
	}
	return q
}

func (p Puzzle) cloneWithoutCandidates() Puzzle {
	q := make(Puzzle, 81)
	for i, v := range p {
		q[i] = cell.New(v.Value)
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
	p, ok, _ := p.backtrackingSolve(false)
	if !ok {
		return nil, fmt.Errorf("could not solve with backtracking")
	}
	return p, nil
}

// the count here assumes the puzzle is not already complete
func (p Puzzle) backtrackingSolve(shouldContinue bool) (Puzzle, bool, int) {
	var (
		backTrackSolutions puzzleSet
	)

	var _backTrackingSolve func(p Puzzle, shouldContinue bool) (Puzzle, bool, int)
	_backTrackingSolve = func(p Puzzle, shouldContinue bool) (Puzzle, bool, int) {
		guessIndex := p.getFirstUnsolvedIndex()
		switch guessIndex {
		case -1:
			backTrackSolutions = backTrackSolutions.add(p)
			return p, true, len(backTrackSolutions)
		case -2:
			return nil, false, 0
		}

		cans := p[guessIndex].ListCandidates()
		canOrder := rand.Perm(len(cans))
		for _, v := range canOrder {
			q := p.clone()
			can := cans[v]
			q.unsafeGuess(guessIndex, can)
			q.basicCheckCellsSeenBy(guessIndex)

			if !q.checkNoUnsolveableCells() {
				continue
			}

			if solution, solved, count := _backTrackingSolve(q, shouldContinue); solved {
				if !shouldContinue {
					return solution, true, count
				}
			}
		}

		if len(backTrackSolutions) > 0 {
			return backTrackSolutions[0], true, len(backTrackSolutions)
		}

		return nil, false, 0
	}

	return _backTrackingSolve(p, shouldContinue)
}
