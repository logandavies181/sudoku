package main

import "fmt"

type cell struct {
	value int
	candidates []int
}

func newCell(n int) cell {
	var candidates []int
	if n == 0 {
		candidates = allCandidates()
	} else {
		candidates = nil
	}

	return cell{
		value: n,
		candidates: candidates,
	}
}

func (c *cell) removeCandidate(n int) {
	c.candidates[n] = 0
}

func (c *cell) numCandidates() int {
	count := 0
	for _, v := range c.candidates[1:] {
		if v == 1 {
			count++
		}
	}

	return count
}

func (c *cell) listCandidates() []int {
	ret := make([]int, 0)

	for i, v := range c.candidates {
		if v == 1 && i != 0 {
			ret = append(ret, i)
		}
	}

	return ret
}

func (c *cell) solve() {
	// todo: handle an error here maybe?
	cds := c.listCandidates()
	if len(cds) == 1 {
		c.value = cds[0]
		c.candidates = nil
	}
}

func (c *cell) solveAs(val int) {
	if ! containsCandidate(c.candidates, val) {
		panic(fmt.Sprint("cell does not contain candidate ", val))
	}

	c.value = val
	c.candidates = nil
}

func cellsFromInts(nums []int) []cell {
	cells := make([]cell, len(nums))
	for i, v := range nums {
		cells[i] = newCell(v)
	}

	return cells
}

func allCandidates() []int {
	ret := make([]int, 10) // one extra space so we're not juggling the off-by-one constantly
	for i := range ret {
		ret[i] = 1
	}
	ret[0] = 0

	return ret
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

	return xPos / 3 + (yPos / 3) * 3
}
