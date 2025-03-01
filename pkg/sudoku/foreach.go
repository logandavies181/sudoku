package sudoku

import "github.com/logandavies181/sudoku/pkg/cell"

var (
	boxes = [][]int{
		{
			0, 1, 2, 9, 10, 11, 18, 19, 20,
		},
		{
			3, 4, 5, 12, 13, 14, 21, 22, 23,
		},
		{
			6, 7, 8, 15, 16, 17, 24, 25, 26,
		},
		{
			27, 28, 29, 36, 37, 38, 45, 46, 47,
		},
		{
			30, 31, 32, 39, 40, 41, 48, 49, 50,
		},
		{
			33, 34, 35, 42, 43, 44, 51, 52, 53,
		},
		{
			54, 55, 56, 63, 64, 65, 72, 73, 74,
		},
		{
			57, 58, 59, 66, 67, 68, 75, 76, 77,
		},
		{
			60, 61, 62, 69, 70, 71, 78, 79, 80,
		},
	}

	rows = [][]int{
		{
			0, 1, 2, 3, 4, 5, 6, 7, 8,
		},
		{
			9, 10, 11, 12, 13, 14, 15, 16, 17,
		},
		{
			18, 19, 20, 21, 22, 23, 24, 25, 26,
		},
		{
			27, 28, 29, 30, 31, 32, 33, 34, 35,
		},
		{
			36, 37, 38, 39, 40, 41, 42, 43, 44,
		},
		{
			45, 46, 47, 48, 49, 50, 51, 52, 53,
		},
		{
			54, 55, 56, 57, 58, 59, 60, 61, 62,
		},
		{
			63, 64, 65, 66, 67, 68, 69, 70, 71,
		},
		{
			72, 73, 74, 75, 76, 77, 78, 79, 80,
		},
	}

	columns = [][]int{
		{
			0, 9, 18, 27, 36, 45, 54, 63, 72,
		},
		{
			1, 10, 19, 28, 37, 46, 55, 64, 73,
		},
		{
			2, 11, 20, 29, 38, 47, 56, 65, 74,
		},
		{
			3, 12, 21, 30, 39, 48, 57, 66, 75,
		},
		{
			4, 13, 22, 31, 40, 49, 58, 67, 76,
		},
		{
			5, 14, 23, 32, 41, 50, 59, 68, 77,
		},
		{
			6, 15, 24, 33, 42, 51, 60, 69, 78,
		},
		{
			7, 16, 25, 34, 43, 52, 61, 70, 79,
		},
		{
			8, 17, 26, 35, 44, 53, 62, 71, 80,
		},
	}
)

func (p Puzzle) foreachUnsolvedCells(cellIds []int, f func(id int)) {
	for _, v := range cellIds {
		c := p[v]
		if !c.Solved() {
			f(v)
		}
	}
}

func (p Puzzle) foreachAllUnsolvedCells(f func(id int)) {
	for i, c := range p {
		if !c.Solved() {
			f(i)
		}
	}
}

func (p Puzzle) foreachSeenBy(id int, f func(id int)) {
	seenByCells := p.getAllSeenBy(id)
	for _, v := range seenByCells {
		f(v)
	}
}

func (p Puzzle) foreachUnsolvedSeenBy(id int, f func(id int)) {
	seenByCells := p.getAllSeenBy(id)
	for _, v := range seenByCells {
		if !p[v].Solved() {
			f(v)
		}
	}
}

func (p Puzzle) foreachEmptyCellIds(cellIds []int, f func(id int, v cell.Cell)) {
	for _, id := range cellIds {
		if p[id].Value == 0 {
			f(id, p[id])
		}
	}
}

func (p Puzzle) foreachFilledCellIds(cellIds []int, f func(id int, v cell.Cell)) {
	for _, id := range cellIds {
		if p[id].Value == 0 {
			f(id, p[id])
		}
	}
}

func (p Puzzle) foreachCellIds(cellIds []int, f func(id int, v cell.Cell)) {
	for _, id := range cellIds {
		f(id, p[id])
	}
}

func (p Puzzle) foreach(iterator [][]int, f func(cellIds []int)) {
	for _, item := range iterator {
		f(item)
	}
}

func (p Puzzle) foreachRow(f func(cellIds []int)) {
	p.foreach(rows, f)
}

func (p Puzzle) foreachBox(f func(cellIds []int)) {
	p.foreach(boxes, f)
}

func (p Puzzle) foreachColumn(f func(cellIds []int)) {
	p.foreach(columns, f)
}

func (p Puzzle) foreachRBC(f func(cellIds []int)) {
	for _, fe := range []func(func(cellIds []int)){p.foreachRow, p.foreachBox, p.foreachColumn} {
		fe(f)
	}
}

func foreachCandidateInCell(c cell.Cell, f func(candidate int)) {
	for i := 1; i < 10; i++ {
		if c.HasCandidate(i) {
			f(i)
		}
	}
}
