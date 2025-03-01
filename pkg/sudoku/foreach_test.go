package sudoku

import (
	"testing"

	"github.com/logandavies181/sudoku/pkg/cell"
	"github.com/stretchr/testify/assert"
)

func TestForeachCandidateInCell(t *testing.T) {
	c := cell.New(0)

	count := 0
	sum := 0
	foreachCandidateInCell(c, func(candidate int) {
		count++
		sum += candidate
	})

	assert.Equal(t, 9, count)
	assert.Equal(t, 45, sum)
}

func TestGetAllSeenBy(t *testing.T) {
	cellIds := getAllSeenBy(0)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 18, 19, 20, 27, 36, 45, 54, 63, 72}, cellIds)
}

func TestForeachSeenBy(t *testing.T) {
	err := InitializeFromFile("../../test/foreachseenbytest.txt")
	assert.NoError(t, err)

	cellIds := make([]int, 0)
	foreachSeenBy(0, func(id int) {
		cellIds = append(cellIds, id)
	})

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 18, 19, 20, 27, 36, 45, 54, 63, 72}, cellIds)
}
