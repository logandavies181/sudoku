package sudoku

func getAllSeenBy(src int) []int {
	ret := make([]int, 20)
	count := 0

	for i := range cells {
		if canSee(src, i) && !(i == src) {
			ret[count] = i
			count++
		}
	}

	return ret
}

func canSee(src, dst int) bool {
	return xPos(src) == xPos(dst) ||
		yPos(src) == yPos(dst) ||
		boxIndex(src) == boxIndex(dst)
}

func xPos(src int) int {
	return src % 9
}

func yPos(src int) int {
	return src / 9
}

func boxIndex(src int) int {
	xPos := xPos(src)
	yPos := yPos(src)

	return xPos/3 + (yPos/3)*3
}
