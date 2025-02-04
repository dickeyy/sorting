package lib

import "golang.org/x/exp/rand"

func MakeRandomArray(size int, min, max int) []int {
	items := make([]int, size)

	for i := 0; i < size; i++ {
		items[i] = rand.Intn(max-min) + min
	}

	return items
}
