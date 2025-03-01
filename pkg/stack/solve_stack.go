package stack

var _solveStack solveStack

type SolveStackItem struct {
	Index int
	Value int
	Guess bool
}

type solveStack []SolveStackItem

func (s solveStack) Push(item SolveStackItem) {
	s = append(s, item)
}

func (s solveStack) Pop() SolveStackItem {
	ret := s[len(s)-1]
	s = s[:len(s)-1]
	return ret
}

func Push(item SolveStackItem) {
	_solveStack.Push(item)
}

func Pop() SolveStackItem {
	return _solveStack.Pop()
}
