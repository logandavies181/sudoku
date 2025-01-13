package cell

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func toBinaryString(i uint16) string {
	return strconv.FormatInt(int64(i), 2)
}

func TestHasCandidate(t *testing.T) {
	c := Cell{
		candidates: 0b1,
	}

	actual := c.HasCandidate(1)

	assert.True(t, actual, toBinaryString(c.candidates))
}

func TestAddCandidate(t *testing.T) {
	c := Cell{
		candidates: 0b1,
	}

	c.AddCandidate(1)

	assert.True(t, c.HasCandidate(1), toBinaryString(c.candidates))
}

func TestRemoveCandidate(t *testing.T) {
	c := New(0)

	changed := c.RemoveCandidate(1)

	assert.True(t, changed)
	assert.False(t, c.HasCandidate(1))
	assert.Equal(t, uint16(0b111111110), c.candidates, toBinaryString(c.candidates))
}

func TestRemoveCandidate_NoChangeIfAlreadyRemoved(t *testing.T) {
	c := New(0)

	changedFirst := c.RemoveCandidate(1)
	changedSecond := c.RemoveCandidate(1)

	assert.True(t, changedFirst)
	assert.False(t, changedSecond)
	assert.Equal(t, uint16(0b111111110), c.candidates, toBinaryString(c.candidates))
}

func TestNumCandidates(t *testing.T) {
	c := New(0)

	actual := c.NumCandidates()

	assert.Equal(t, 9, actual)
}

func TestNumCandidates_SomeRemoved(t *testing.T) {
	c := New(0)
	c.RemoveCandidate(1)
	c.RemoveCandidate(2)

	actual := c.NumCandidates()

	assert.Equal(t, 7, actual)
}

func TestNumCandidates_NewSolved(t *testing.T) {
	c := New(1)

	actual := c.NumCandidates()

	assert.Equal(t, 0, actual)
}

func TestNumCandidates_SolvedLater(t *testing.T) {
	c := New(0)
	c.SolveAs(1)

	actual := c.NumCandidates()

	assert.Equal(t, 0, actual)
}

func TestListCandidates(t *testing.T) {
	c := New(0)

	actual := c.ListCandidates()

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, actual)
}

func TestListCandidates_SomeRemoved(t *testing.T) {
	c := New(0)
	c.RemoveCandidate(1)
	c.RemoveCandidate(2)

	actual := c.ListCandidates()

	assert.Equal(t, []int{3, 4, 5, 6, 7, 8, 9}, actual)
}
