package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	cells []cell
	nums = []int{}
	start = time.Now()
)


func printPuzzle() {
	boxRowDivider := "-------------"
	for i, v := range cells {
		switch {
		case i % 27 == 0:
			fmt.Println(boxRowDivider)
			fmt.Printf("|%d", v.value)
		case i % 27 == 26:
			fmt.Printf("%d|\n", v.value)
		case i % 9 == 0:
			fmt.Printf("|\n|%d", v.value)
		case i % 3 == 0:
			fmt.Printf("|%d", v.value)
		default:
			fmt.Print(v.value)
		}
	}
	fmt.Println(boxRowDivider)
}

func validatePuzzle() error {
	var err error

	foreachRBC(func(cellIds []int) {
		if err != nil {
			return
		}

		valueCounts := newIntIntMap()
		foreachCellIds(cellIds, func(id int, v cell) {
			valueCounts.IncrementKey(v.value)
		})

		for k, v := range valueCounts {
			if v > 1 {
				err = fmt.Errorf("too many of: %d\n", k)
				break
			}
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

	for {
		basicCheckCells()
		if
			! basicSolveRBCSingle() &&
			! checkBoxLinearCandidates() &&
			! updateSolvedCells() {

			break
		}
	}

	printPuzzle()

	err = validatePuzzle()
	if err != nil {
		return err
	}

	fmt.Println(time.Now().Sub(start))

	return nil
}
