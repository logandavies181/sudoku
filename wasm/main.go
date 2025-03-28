//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
)

func addExportFunc(name string, fn func(js.Value, []js.Value) any) {
	js.Global().Set(name, js.FuncOf(fn))
}

func main() {
	addExportFunc("NewPuzzle", NewPuzzle)
	addExportFunc("UpdateCell", UpdateCell)
	addExportFunc("CheckIfCanUpdateCell", CheckIfCanUpdateCell)
	fmt.Println("sudoku wasm instantiated")
	select {}
}
