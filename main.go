package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/logandavies181/sudoku/cell"
)

var (
	cells []cell.Cell
	nums  = []int{}
)

func printPuzzle() {
	boxRowDivider := "-------------"
	for i, v := range cells {
		switch {
		case i%27 == 0:
			fmt.Println(boxRowDivider)
			fmt.Printf("|%d", v.Value)
		case i%27 == 26:
			fmt.Printf("%d|\n", v.Value)
		case i%9 == 0:
			fmt.Printf("|\n|%d", v.Value)
		case i%3 == 0:
			fmt.Printf("|%d", v.Value)
		default:
			fmt.Print(v.Value)
		}
	}
	fmt.Println(boxRowDivider)
}

func validatePuzzle() error {
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

func initializeFromFile(fname string) error {
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

func main() {
	err := mainE()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mainE() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("file name required")
	}

	err := initializeFromFile(os.Args[1])
	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	start := time.Now()

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

	printPuzzle()

	err = validatePuzzle()
	fmt.Println(time.Now().Sub(start))

	return err
}

func cellsFromInts(nums []int) []cell.Cell {
	cells := make([]cell.Cell, len(nums))
	for i, v := range nums {
		cells[i] = cell.New(v)
	}

	return cells
}
