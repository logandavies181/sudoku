package main

import "github.com/logandavies181/sudoku/pkg/sudoku"

var state []int

func sudokuToIntArr(p sudoku.Puzzle) []int {
	ret := make([]int, len(p))
	for i, v := range p {
		ret[i] = v.Value
	}
	return ret
}

func NewPuzzle(targetClueCount, maxTries int) []int {
	p := sudoku.NewPuzzle(targetClueCount, maxTries)
	return sudokuToIntArr(p)
}
