package sudoku

type puzzleSet []Puzzle

func (ps puzzleSet) add(p Puzzle) puzzleSet {
	if !ps.has(p) {
		ps = append(ps, p)
	}
	return ps
}

func (ps puzzleSet) has(q Puzzle) bool {
	for _, p := range ps {
		matches := true
		for i := range p {
			if p[i].Value != q[i].Value {
				matches = false
				break
			}
		}
		if matches {
			return true
		}
	}

	return false
}
