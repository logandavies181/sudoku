package sudoku

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/logandavies181/sudoku/pkg/cell"
)

var (
	cells []cell.Cell
)

func solveCellAs(id int, val int) {
	cells[id].SolveAs(val)
	/*
	stack.Push(stack.SolveStackItem{
		Index: id,
		Value: val,
		Guess: false,
	})
	*/
}

// todo: make this safe
func unsafeGuess(id int, val int) {
	cells[id].SolveAs(val)
	/*
	stack.Push(stack.SolveStackItem{
		Index: id,
		Value: val,
		Guess: true,
	})
	*/
}

func Solve() error {
	for {
		shouldBreak := true
		for _, alg := range []func() bool{
			basicCheckCells,
			basicSolveRBCSingle,
			checkBoxLinearCandidates,
			updateSolvedCells,
		} {
			if alg() {
				shouldBreak = false
			}
		}

		if shouldBreak {
			break
		}
	}

	return ValidatePuzzle()
}

func InitializeFromFile(fname string) error {
	f, err := os.Open(fname)
	if err != nil {
		return err
	}

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	nums := make([]int, 81)
	count := 0
	for _, line := range lines {
		for _, v := range line {
			nums[count], err = strconv.Atoi(v)
			if err != nil {
				return err
			}

			count++
		}
	}

	if count != 81 {
		return fmt.Errorf("bad input file. incorrect number of cells")
	}

	cells = cellsFromInts(nums)

	basicCheckCells()

	return nil
}

func ValidatePuzzle() error {
	var err error

	// check for duplicates in RBC
	foreachRBC(func(cellIds []int) {
		if err != nil {
			return
		}

		valueCounts := newIntIntMap()
		foreachCellIds(cellIds, func(id int, v cell.Cell) {
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
	foreachAllUnsolvedCells(func(id int) {
		if err != nil {
			return
		}

		found := false
		c := cells[id]
		foreachCandidateInCell(c, func(candidate int) {
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

func cellsFromInts(nums []int) []cell.Cell {
	cells := make([]cell.Cell, len(nums))
	for i, v := range nums {
		cells[i] = cell.New(v)
	}

	return cells
}
