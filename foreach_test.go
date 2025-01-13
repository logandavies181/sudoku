package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForeachCandidateInCell(t *testing.T) {
	c := newCell(0)

	count := 0
	sum := 0
	foreachCandidateInCell(c, func(candidate int) {
		count++
		sum += candidate
	})

	assert.Equal(t, 9, count)
	assert.Equal(t, 45, sum)
}
