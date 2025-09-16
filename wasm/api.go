//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/logandavies181/sudoku/pkg/cell"
	"github.com/logandavies181/sudoku/pkg/sudoku"
)

var state sudoku.Puzzle

func sudokuToIntArr(p sudoku.Puzzle) []any {
	ret := make([]any, len(p))
	for i, v := range p {
		ret[i] = v.Value
	}
	return ret
}

// (targetCount: number, maxTries: number) => number[]
func NewPuzzle(_ js.Value, args []js.Value) any {
	p := sudoku.NewPuzzle(args[0].Int(), args[1].Int())
	state = p
	return sudokuToIntArr(p)
}

// (index: number, value: number) => boolean
func UpdateCell(_ js.Value, args []js.Value) any {
	index := args[0].Int()
	value := args[1].Int()

	if value == 0 {
		state[index] = cell.New(0)
		return true
	}

	err := state[index].SolveAs(value)
	if err != nil {
		fmt.Println(err)
	}

	return err == nil
}

// (index: number, proposedValue: number) boolean
func CheckIfCanUpdateCell(_ js.Value, args []js.Value) any {
	index := args[0].Int()
	proposedValue := args[1].Int()

	if proposedValue == 0 {
		return true
	}

	found := false
	state.ForeachSeenBy(index, func(id int) {
		if state[id].Value == proposedValue {
			found = true
		}
	})

	return !found
}
