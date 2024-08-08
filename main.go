package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	start = time.Now()
	nums = []int{}
	cells = cellsFromInts(nums)
)

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

func printPuzzle() {
	boxRowDivider := "-------------"
	for i, v := range cells {
		switch {
		case i % 27 == 0:
			fmt.Println(boxRowDivider)
			fmt.Printf("|%d", v.value)
		case i % 27 == 26:
			fmt.Printf("%d|\n", v.value)
		case i % 9 == 0:
			fmt.Printf("|\n|%d", v.value)
		case i % 3 == 0:
			fmt.Printf("|%d", v.value)
		default:
			fmt.Print(v.value)
		}
	}
	fmt.Println(boxRowDivider)
}

func basicCheckRowsColsBoxes() {
	foreachUnsolvedCells(func(i int, v cell) {
		seenByCellIds := getAllSeenBy(i)
		for _, seenByCellId := range seenByCellIds {
			v.removeCandidate(cells[seenByCellId].value)
		}
	})
}

func updateSolvedCells() bool {
	found := false

	foreachUnsolvedCells(func(i int, v cell) {
		if cells[i].numCandidates() == 1 && cells[i].value == 0 {
			cells[i].solve()
			found = true
		}
	})

	return found
}

func main() {
	err := mainE()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mainE() error {
	if len(os.Args) == 2 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			return err
		}

		lines, err := csv.NewReader(f).ReadAll()
		if err != nil {
			return err
		}

		nums := make([]int, 81)
		count := 0
		for _, line := range lines {
			for i, v := range line {
				nums[i], err = strconv.Atoi(v)
				if err != nil {
					return err
				}

				count++
			}
		}

	} else {
		return fmt.Errorf("file name required")
	}

	atLeastOneSolved := true
	for atLeastOneSolved {
		atLeastOneSolved = false
		
		basicCheckRowsColsBoxes()

		atLeastOneSolved = updateSolvedCells()
	}

	printPuzzle()

	fmt.Println(time.Now().Sub(start))

	return nil
}
