package stack

type SolveStackItem struct {
	Index int
	Value int
	Guess bool
}

type solveStack []SolveStackItem

func (s *solveStack) Push(item SolveStackItem) {
	*s = append(*s, item)
}

func (s *solveStack) Pop() *SolveStackItem {
	if len(*s) == 0 {
		return nil
	}

	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return &ret
}

func New() solveStack {
	return make(solveStack, 0)
}
