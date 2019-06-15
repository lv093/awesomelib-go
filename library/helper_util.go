package library

import "math/rand"

func GenRandArray(n int, min int, max int) []int {
	if n <= 0 {
		return nil
	}
	arr := make([]int, n)
	i := 0
	for i < n {
		arr[i] = rand.Intn(max - min) + min
		i ++
	}
	return arr
}