package main

import "fmt"

func GetNext(text string) []int {
	next := make([]int,len(text))
	next[0], next[1] = -1, 0
	j, k := 1, 0

	for j < len(text) - 1 {
		if k == -1 || text[j] == text[k] {
			next[j+1] = k + 1
			j++
			k++
		} else {
			k = next[k]
		}
	}
	return next
}

func KMP(pattern string, text string) int {
	i, j := 0, 0
	next := GetNext(pattern)

	for i < len(text) && j < len(pattern) {
		if j == -1 || text[i] == pattern[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(pattern) {
		return i - j
	} else {
		return -1
	}
}

func main() {
	pattern := "ABAB"
	next := GetNext(pattern)
	fmt.Printf("next===>%v\n", next)
	fmt.Printf("RESULT ===>%d", KMP(pattern, "ddABABCdABAB"))
}

