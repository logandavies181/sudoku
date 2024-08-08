package main

// remove any candidates in each cell whose values can be seen from that cell
func basicCheckCells() {
	foreachUnsolvedCells(func(id int, v cell) {
		seenByCellIds := getAllSeenBy(id)
		for _, seenByCellId := range seenByCellIds {
			v.removeCandidate(cells[seenByCellId].value)
		}
	})
}

// check only one valid spot in the r/b/c for a given number
func basicSolveRBCSingle() bool {
	found := false
	foreachRBC(func(cellIds []int) {
		candidateCounts := newIntIntMap()

		foreachEmptyCellIds(cellIds, func(id int, v cell) {
			for val, exists := range cells[id].candidates {
				if exists == 1 {
					candidateCounts.IncrementKey(val)
				}
			}
		})

		foreachFilledCellIds(cellIds, func(id int, v cell) {
			delete(candidateCounts, v.value)
		})

		singleCandidates := make([]int, 0)
		for k, v := range candidateCounts {
			if v == 1 {
				singleCandidates = append(singleCandidates, k)
			}
		}

		if len(singleCandidates) == 0 {
			return
		}

		for _, can := range singleCandidates {
			can := can
			foreachEmptyCellIds(cellIds, func(id int, v cell) {
				if containsCandidate(v.candidates, can) {
					found = true
					cells[id].solveAs(can)
					basicCheckCells()
				}
			})
		}
	})

	return found
}

// update solve any cells that have only one candidate remaining
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
