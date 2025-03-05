//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/logandavies181/sudoku/pkg/sudoku"
)

var state []int

func sudokuToIntArr(p sudoku.Puzzle) []any {
	ret := make([]any, len(p))
	for i, v := range p {
		ret[i] = v.Value
	}
	return ret
}

func NewPuzzle(_ js.Value, args []js.Value) any {
	p := sudoku.NewPuzzle(args[0].Int(), args[1].Int())
	return sudokuToIntArr(p)
}
