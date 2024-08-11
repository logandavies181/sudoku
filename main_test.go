package main

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
	err := initializeFromFile("test/checkBoxLinearCandidates.txt")
	assert.Nil(t, err)

	updated := checkBoxLinearCandidates()

	assert.True(t, updated)
	assert.True(t, cells[3].candidates[7] == 0)
}

func Test_CheckBoxLinearCandidates_NoopCompletedPuzzle(t *testing.T) {
	err := initializeFromFile("test/completed.txt")
	assert.Nil(t, err)

	updated := checkBoxLinearCandidates()

	assert.False(t, updated)
}
