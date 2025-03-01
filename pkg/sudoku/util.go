package sudoku

import (
	"fmt"

	"github.com/logandavies181/sudoku/pkg/cell"
)

type intIntMap map[int]int

func newIntIntMap() intIntMap {
	return make(intIntMap)
}

func (iim intIntMap) IncrementKey(i int) {
	v, ok := iim[i]
	if ok {
		iim[i] = v + 1
	} else {
		iim[i] = 1
	}
}

func (p Puzzle) getCandidateCounts(cellIds []int) intIntMap {
	candidateCounts := newIntIntMap()
	p.foreachEmptyCellIds(cellIds, func(id int, v cell.Cell) {
		p.foreachCandidateInCell(p[id], func(candidate int) {
			candidateCounts.IncrementKey(candidate)
		})
	})

	return candidateCounts
}

func (p Puzzle) locateCandidates(cellIds []int, candidate int) []int {
	ret := make([]int, 0)
	for _, v := range cellIds {
		if p[v].HasCandidate(candidate) {
			ret = append(ret, v)
		}
	}

	return ret
}

func compareCells(cellIds []int, f func(int, int) bool) bool {
	prev := cellIds[0]
	for _, v := range cellIds {
		if !f(prev, v) {
			return false
		}
		prev = v
	}

	return true
}

func allInSameColumn(cellIds []int) bool {
	return compareCells(cellIds, func(i1, i2 int) bool {
		return xPos(i1) == xPos(i2)
	})
}

func allInSameRow(cellIds []int) bool {
	return compareCells(cellIds, func(i1, i2 int) bool {
		return yPos(i1) == yPos(i2)
	})
}

func (p Puzzle) removeCandidatesFromRow(yPos int, candidate int) int {
	count := 0
	for _, v := range rows[yPos] {
		if p[v].RemoveCandidate(candidate) {
			count++
		}
	}
	return count
}

func (p Puzzle) removeCandidatesFromColumn(xPos int, candidate int) int {
	count := 0
	for _, v := range columns[xPos] {
		if p[v].RemoveCandidate(candidate) {
			count++
		}
	}
	return count
}

func (p Puzzle) addCandidateToCells(cellIds []int, candidate int) int {
	count := 0
	for _, v := range cellIds {
		p[v].AddCandidate(candidate)
		count++
	}
	return count
}

func (p Puzzle) Print() {
	boxRowDivider := "-------------"
	for i, v := range p {
		switch {
		case i%27 == 0:
			fmt.Println(boxRowDivider)
			fmt.Printf("|%d", v.Value)
		case i%27 == 26:
			fmt.Printf("%d|\n", v.Value)
		case i%9 == 0:
			fmt.Printf("|\n|%d", v.Value)
		case i%3 == 0:
			fmt.Printf("|%d", v.Value)
		default:
			fmt.Print(v.Value)
		}
	}
	fmt.Println(boxRowDivider)
}
