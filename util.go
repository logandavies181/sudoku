package main

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

func getCandidateCounts(cellIds []int) intIntMap {
	candidateCounts := newIntIntMap()
	foreachEmptyCellIds(cellIds, func(id int, v cell) {
		foreachCandidateInCell(cells[id], func(candidate int) {
			candidateCounts.IncrementKey(candidate)
		})
	})

	return candidateCounts
}

func locateCandidates(cellIds []int, candidate int) []int {
	ret := make([]int, 0)
	for _, v := range cellIds {
		if cells[v].hasCandidate(candidate) {
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

func removeCandidatesFromRow(yPos int, candidate int) int {
	count := 0
	for _, v := range rows[yPos] {
		if cells[v].removeCandidate(candidate) {
			count++
		}
	}
	return count
}

func removeCandidatesFromColumn(xPos int, candidate int) int {
	count := 0
	for _, v := range columns[xPos] {
		if cells[v].removeCandidate(candidate) {
			count++
		}
	}
	return count
}

func addCandidateToCells(cellIds []int, candidate int) int {
	count := 0
	for _, v := range cellIds {
		cells[v].addCandidate(candidate)
		count++
	}
	return count
}
