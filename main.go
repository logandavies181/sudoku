package main

import (
	"fmt"
	"os"
	"time"

	"github.com/logandavies181/sudoku/pkg/sudoku"
)


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

	puzzle, err := sudoku.InitializeFromFile(os.Args[1])
	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	start := time.Now()

	err = puzzle.Solve()

	fmt.Println(time.Now().Sub(start))

	puzzle.Print()

	return err
}
