package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSee_CanSee(t *testing.T) {
	assert.True(t, canSee(3, 57))
	assert.True(t, canSee(3, 4))
	assert.True(t, canSee(4, 3))
	assert.True(t, canSee(74, 65))
	assert.False(t, canSee(11, 3))
}

func TestCanSee_BoxIndex(t *testing.T) {
	assert.Equal(t, 1, boxIndex(3))
	assert.Equal(t, 5, boxIndex(43))
}

func TestCanSee_XPos(t *testing.T) {
	assert.Equal(t, 8, xPos(8))
	assert.Equal(t, 7, xPos(43))
}

func TestCanSee_YPos(t *testing.T) {
	assert.Equal(t, 0, yPos(8))
	assert.Equal(t, 4, yPos(43))
}

func Test_CheckBoxLinearCandidates_Basic(t *testing.T) {
	err := InitializeFromFile("../../test/checkBoxLinearCandidates.txt")
	assert.NoError(t, err)

	updated := checkBoxLinearCandidates()

	assert.True(t, updated)
	assert.True(t, !cells[3].HasCandidate(7))
}

func Test_CheckBoxLinearCandidates_NoopCompletedPuzzle(t *testing.T) {
	err := InitializeFromFile("../../test/completed.txt")
	assert.NoError(t, err)

	updated := checkBoxLinearCandidates()

	assert.False(t, updated)
}

func TestGetCandidateCounts_NoneCompletedPuzzle(t *testing.T) {
	err := InitializeFromFile("../../test/completed.txt")
	assert.NoError(t, err)

	// TODO
}
