package cell

import "fmt"

type Cell struct {
	Value      int
	candidates uint16
}

func New(n int) Cell {
	var candidates uint16
	if n == 0 {
		candidates = allCandidates()
	} else {
		candidates = 0
	}

	return Cell{
		Value:      n,
		candidates: candidates,
	}
}

func (c *Cell) Solved() bool {
	if (c.Value != 0) != (c.candidates == 0) {
		panic("must have either value or candidates")
	}

	return c.Value != 0 && c.candidates == 0
}

func (c *Cell) HasCandidate(n int) bool {
	return c.candidates&(1<<(n-1)) > 0
}

func (c *Cell) AddCandidate(n int) {
	c.candidates |= (1 << (n - 1))
}

func (c *Cell) RemoveCandidate(n int) bool {
	if n == 0 {
		return false
	}

	initial := c.candidates
	c.candidates &= ^(1 << (n - 1))
	return c.candidates != initial
}

func (c *Cell) NumCandidates() int {
	candidates := c.candidates
	count := 0
	for i := 0; i < 9; i++ {
		// todo use hasCandidates
		if candidates&1 == 1 {
			count++
		}
		candidates >>= 1
	}

	return count
}

func (c *Cell) ListCandidates() []int {
	ret := make([]int, 0)

	for i := 1; i < 10; i++ {
		if c.HasCandidate(i) {
			ret = append(ret, i)
		}
	}

	return ret
}

func (c *Cell) Solve() {
	// todo: handle an error here maybe?
	cds := c.ListCandidates()
	if len(cds) == 1 {
		c.SolveAs(cds[0])
	}
}

func (c *Cell) SolveAs(val int) {
	if !c.HasCandidate(val) {
		panic(fmt.Sprint("cell does not contain candidate ", val))
	}

	c.Value = val
	c.candidates = 0
}

func allCandidates() uint16 {
	return 0b111111111
}
