package main

type intIntMap map[int]int

func newIntIntMap() intIntMap {
	return make(intIntMap)
}

func (iim intIntMap) IncrementKey(i int) {
	v, ok := iim[i]
	if ok {
		iim[i] = v + 1
	} else {
		iim[i] = 1
	}
}

func contains[T comparable](s []T, t T) bool {
	for _, v := range s {
		if v == t {
			return true
		}
	}

	return false
}
