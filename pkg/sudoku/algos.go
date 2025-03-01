package sudoku

import (
	"github.com/logandavies181/sudoku/pkg/cell"
)

// remove any candidates in each cell whose values can be seen from that cell
func (p Puzzle) basicCheckCells() bool {
	found := false
	p.foreachAllUnsolvedCells(func(id int) {
		seenByCellIds := p.getAllSeenBy(id)
		for _, seenByCellId := range seenByCellIds {
			val := p[seenByCellId].Value
			if val != 0 && p[id].RemoveCandidate(val) {
				found = true
			}
		}
	})

	return found
}

// only check cells seen by a given cell
func (p Puzzle) basicCheckCellsSeenBy(id int) bool {
	found := false
	p.foreachUnsolvedSeenBy(id, func(id int) {
		seenByCellIds := p.getAllSeenBy(id)
		for _, seenByCellId := range seenByCellIds {
			val := p[seenByCellId].Value
			if val != 0 && p[id].RemoveCandidate(val) {
				found = true
			}
		}
	})

	return found
}

// check only one valid spot in the r/b/c for a given number
func (p Puzzle) basicSolveRBCSingle() bool {
	found := false
	p.foreachRBC(func(cellIds []int) {
		candidateCounts := p.getCandidateCounts(cellIds)

		p.foreachFilledCellIds(cellIds, func(id int, v cell.Cell) {
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
			p.foreachEmptyCellIds(cellIds, func(id int, v cell.Cell) {
				if v.HasCandidate(can) {
					found = true
					p.solveCellAs(id, can)
					p.basicCheckCellsSeenBy(id)
				}
			})
		}
	})

	return found
}

// check for 2/3 lined up candidates within a box
func (p Puzzle) checkBoxLinearCandidates() bool {
	found := false
	p.foreachBox(func(cellIds []int) {
		p.foreachUnsolvedCells(cellIds, func(id int) {
			candidateCounts := p.getCandidateCounts(cellIds)
			for candidate, count := range candidateCounts {
				if count > 3 || count < 2 {
					// 0 is solved, 4 is too many
					continue
				}

				locations := p.locateCandidates(cellIds, candidate)

				if allInSameRow(locations) {
					rowId := yPos(locations[0])

					// TODO: must be a better way
					numRemoved := p.removeCandidatesFromRow(rowId, candidate)
					numRestored := p.addCandidateToCells(locations, candidate)
					if numRemoved > numRestored {
						found = true
					}
				}

				if allInSameColumn(locations) {
					colId := xPos(locations[0])

					numRemoved := p.removeCandidatesFromColumn(colId, candidate)
					numRestored := p.addCandidateToCells(locations, candidate)
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
func (p Puzzle) updateSolvedCells() bool {
	found := false

	p.foreachAllUnsolvedCells(func(i int) {
		if p[i].NumCandidates() == 1 {
			cans := p[i].ListCandidates()
			p.solveCellAs(i, cans[0])
			found = true
		}
	})

	return found
}
