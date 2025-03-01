package sudoku

import (
	"github.com/logandavies181/sudoku/pkg/cell"
)

// remove any candidates in each cell whose values can be seen from that cell
func basicCheckCells() bool {
	found := false
	foreachAllUnsolvedCells(func(id int) {
		seenByCellIds := getAllSeenBy(id)
		for _, seenByCellId := range seenByCellIds {
			val := cells[seenByCellId].Value
			if val != 0 && cells[id].RemoveCandidate(val) {
				found = true
			}
		}
	})

	return found
}

// only check cells seen by a given cell
func basicCheckCellsSeenBy(id int) bool {
	found := false
	foreachUnsolvedSeenBy(id, func(id int) {
		seenByCellIds := getAllSeenBy(id)
		for _, seenByCellId := range seenByCellIds {
			val := cells[seenByCellId].Value
			if val != 0 && cells[id].RemoveCandidate(val) {
				found = true
			}
		}
	})

	return found
}

// check only one valid spot in the r/b/c for a given number
func basicSolveRBCSingle() bool {
	found := false
	foreachRBC(func(cellIds []int) {
		candidateCounts := getCandidateCounts(cellIds)

		foreachFilledCellIds(cellIds, func(id int, v cell.Cell) {
			delete(candidateCounts, v.Value)
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
			foreachEmptyCellIds(cellIds, func(id int, v cell.Cell) {
				if v.HasCandidate(can) {
					found = true
					solveCellAs(id, can)
					basicCheckCellsSeenBy(id)
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
		foreachUnsolvedCells(cellIds, func(id int) {
			candidateCounts := getCandidateCounts(cellIds)
			for candidate, count := range candidateCounts {
				if count > 3 || count < 2 {
					// 0 is solved, 4 is too many
					continue
				}

				locations := locateCandidates(cellIds, candidate)

				if allInSameRow(locations) {
					rowId := yPos(locations[0])

					// TODO: must be a better way
					numRemoved := removeCandidatesFromRow(rowId, candidate)
					numRestored := addCandidateToCells(locations, candidate)
					if numRemoved > numRestored {
						found = true
					}
				}

				if allInSameColumn(locations) {
					colId := xPos(locations[0])

					numRemoved := removeCandidatesFromColumn(colId, candidate)
					numRestored := addCandidateToCells(locations, candidate)
					if numRemoved > numRestored {
						found = true
					}
				}
			}
		})
	})

	return found
}

// update solve any cells that have only one candidate remaining
func updateSolvedCells() bool {
	found := false

	foreachAllUnsolvedCells(func(i int) {
		if cells[i].NumCandidates() == 1 {
			cans := cells[i].ListCandidates()
			solveCellAs(i, cans[0])
			found = true
		}
	})

	return found
}
