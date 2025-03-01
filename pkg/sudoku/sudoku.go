package sudoku

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/logandavies181/sudoku/pkg/cell"
)

type Puzzle []cell.Cell

func (p Puzzle) solveCellAs(id int, val int) {
	p[id].SolveAs(val)
	/*
		stack.Push(stack.SolveStackItem{
			Index: id,
			Value: val,
			Guess: false,
		})
	*/
}

// todo: make this safe
func (p Puzzle) unsafeGuess(id int, val int) {
	p[id].SolveAs(val)
	/*
		stack.Push(stack.SolveStackItem{
			Index: id,
			Value: val,
			Guess: true,
		})
	*/
}

func (p Puzzle) Solve() error {
	for {
		shouldBreak := true
		for _, alg := range []func() bool{
			p.basicCheckCells,
			p.basicSolveRBCSingle,
			p.checkBoxLinearCandidates,
			p.updateSolvedCells,
		} {
			if alg() {
				shouldBreak = false
			}
		}

		if shouldBreak {
			break
		}
	}

	return p.Validate()
}

func InitializeFromFile(fname string) (*Puzzle, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	nums := make([]int, 81)
	count := 0
	for _, line := range lines {
		for _, v := range line {
			nums[count], err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}

			count++
		}
	}

	if count != 81 {
		return nil, fmt.Errorf("bad input file. incorrect number of cells")
	}

	p := newFromInts(nums)

	p.basicCheckCells()

	return &p, nil
}

func (p Puzzle) Validate() error {
	var err error

	// check for duplicates in RBC
	p.foreachRBC(func(cellIds []int) {
		if err != nil {
			return
		}

		valueCounts := newIntIntMap()
		p.foreachCellIds(cellIds, func(id int, v cell.Cell) {
			valueCounts.IncrementKey(v.Value)
		})

		for k, v := range valueCounts {
			if k > 0 && v > 1 {
				err = fmt.Errorf("too many of: %d\n", k)
				break
			}
		}
	})

	// check for cells with no candidates
	// fixme: why tho?
	p.foreachAllUnsolvedCells(func(id int) {
		if err != nil {
			return
		}

		found := false
		c := p[id]
		p.foreachCandidateInCell(c, func(candidate int) {
			if c.HasCandidate(candidate) {
				found = true
			}
		})

		if !found {
			err = fmt.Errorf("no available candidates in cell: %d", id)
		}
	})

	return err
}

func newFromInts(nums []int) Puzzle {
	cells := make([]cell.Cell, len(nums))
	for i, v := range nums {
		cells[i] = cell.New(v)
	}

	return cells
}
