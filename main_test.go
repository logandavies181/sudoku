package main

import(
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
