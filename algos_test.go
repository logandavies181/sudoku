package main

import (
	"testing"

	"github.com/logandavies181/sudoku/cell"
	"github.com/stretchr/testify/assert"
)

func withMockedCells(mockedCells []cell.Cell, f func()) {
	cellsBefore := cells
	cells = mockedCells
	f()
	cells = cellsBefore
}

func TestGetCandidateCounts(t *testing.T) {
	withMockedCells([]cell.Cell{cell.New(0)}, func() {
		actual := getCandidateCounts([]int{0})

		expected := intIntMap{
			1: 1,
			2: 1,
			3: 1,
			4: 1,
			5: 1,
			6: 1,
			7: 1,
			8: 1,
			9: 1,
		}

		assert.Equal(t, expected, actual)
	})
}
