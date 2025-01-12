package main

import "fmt"

type cell struct {
	value      int
	candidates uint16
}

func newCell(n int) cell {
	var candidates uint16
	if n == 0 {
		candidates = allCandidates()
	} else {
		candidates = 0
	}

	return cell{
		value:      n,
		candidates: candidates,
	}
}

func (c *cell) hasCandidate(n int) bool {
	return c.candidates & (1 << n - 1) > 0
}

func (c *cell) addCandidate(n int) {
	c.candidates |= (1 << n - 1)
}

func (c *cell) removeCandidate(n int) bool {
	initial := c.candidates
	c.candidates ^= (1 << n - 1)
	return c.candidates != initial
}

func (c *cell) numCandidates() int {
	candidates := c.candidates
	count := 0
	for i := 0; i < 9; i++ {
		// todo use hasCandidates
		if candidates & 1 == 1 {
			count++
		}
		candidates >>= 1
	}

	return count
}

func (c *cell) listCandidates() []int {
	candidates := c.candidates
	ret := make([]int, 0)

	for i := 0; i < 9; i++ {
		// todo use hasCandidates
		if candidates & 1 == 1 {
			ret = append(ret, i+1)
		}
	}

	return ret
}

func (c *cell) solve() {
	// todo: handle an error here maybe?
	cds := c.listCandidates()
	if len(cds) == 1 {
		c.value = cds[0]
		c.candidates = 0
	}
}

func (c *cell) solveAs(val int) {
	if !c.hasCandidate(val) {
		panic(fmt.Sprint("cell does not contain candidate ", val))
	}

	c.value = val
	c.candidates = 0
}

func cellsFromInts(nums []int) []cell {
	cells := make([]cell, len(nums))
	for i, v := range nums {
		cells[i] = newCell(v)
	}

	return cells
}

func allCandidates() uint16 {
	return 0b111111111
}

func getAllSeenBy(src int) []int {
	ret := make([]int, 20)
	count := 0

	for i := range cells {
		if canSee(src, i) && !(i == src) {
			ret[count] = i
			count++
		}
	}

	return ret
}

func canSee(src, dst int) bool {
	return xPos(src) == xPos(dst) ||
		yPos(src) == yPos(dst) ||
		boxIndex(src) == boxIndex(dst)
}

func xPos(src int) int {
	return src % 9
}

func yPos(src int) int {
	return src / 9
}

func boxIndex(src int) int {
	xPos := xPos(src)
	yPos := yPos(src)

	return xPos/3 + (yPos/3)*3
}
