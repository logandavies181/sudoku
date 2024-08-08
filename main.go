package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	start = time.Now()
	nums = []int{}
	cells = cellsFromInts(nums)
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

func main() {
	err := mainE()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mainE() error {
	if len(os.Args) == 2 {
		f, err := os.Open(os.Args[1])
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
			for i, v := range line {
				nums[i], err = strconv.Atoi(v)
				if err != nil {
					return err
				}

				count++
			}
		}

	} else {
		return fmt.Errorf("file name required")
	}

	atLeastOneSolved := true
	for atLeastOneSolved {
		atLeastOneSolved = false
		
		basicCheckCell()
		basicCheckRBCSingle()

		atLeastOneSolved = updateSolvedCells()
	}

	printPuzzle()

	fmt.Println(time.Now().Sub(start))

	return nil
}
