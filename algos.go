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
		candidateCounts := getCandidateCounts(cellIds)

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

// check for 2/3 lined up candidates within a box
func checkBoxLinearCandidates() bool {
	found := false

	foreachBox(func(cellIds []int) {
		foreachEmptyCellIds(cellIds, func(id int, v cell) {
			candidateCounts := getCandidateCounts(cellIds)
			for candidate, count := range candidateCounts {
				if count > 3 && count < 2 {
					// 1 is solved, 4 is too many
					continue
				}

				found = true

				locations := locateCandidates(cellIds, candidate)

				if allInSameRow(locations) {
					rowId := yPos(locations[0])
					removeCandidatesFromRow(rowId, candidate)
					addCandidateToCells(locations, candidate)
				}

				if allInSameColumn(locations) {
					colId := xPos(locations[0])
					removeCandidatesFromColumn(colId, candidate)
					addCandidateToCells(locations, candidate)
				}
			}
		})
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
