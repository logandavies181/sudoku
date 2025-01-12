package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func toBinaryString(i uint16) string {
	return strconv.FormatInt(int64(i), 2)
}

func TestHasCandidate(t *testing.T) {
	c := cell{
		candidates: 0b1,
	}

	actual := c.hasCandidate(1)

	assert.True(t, actual, toBinaryString(c.candidates))
}

func TestAddCandidate(t *testing.T) {
	c := newCell(0)

	c.addCandidate(1)

	assert.True(t, c.hasCandidate(1), toBinaryString(c.candidates))
}
