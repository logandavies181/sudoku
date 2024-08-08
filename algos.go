package main

// remove any candidates in each cell whose values can be seen from that cell
func basicCheckCell() {
	foreachUnsolvedCells(func(i int, v cell) {
		seenByCellIds := getAllSeenBy(i)
		for _, seenByCellId := range seenByCellIds {
			v.removeCandidate(cells[seenByCellId].value)
		}
	})
}

// check only one valid spot in the r/b/c for a given number
func basicCheckRBCSingle() {
	foreachRBC(func(cellIds []int) {
		candidateCounts := newIntIntMap()

		foreachCellIds(cellIds, func(id int, v cell) {
			for _, can := range cells[id].candidates {
				candidateCounts.IncrementKey(can)
			}
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
			foreachCellIds(cellIds, func(id int, v cell) {
				if contains(v.candidates, can) {
					cells[id].solve()
				}
			})
		}
	})
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
