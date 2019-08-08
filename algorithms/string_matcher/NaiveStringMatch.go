package main

import "fmt"

func BruteForceAlgo(text string, pattern string) int {
	n := len(text)
	m := len(pattern)

	for i := 0; i < n - m + 1; i++ {
		for j := 0; j < m; j++ {
			if text[j + i] != pattern[j] {
				break
			} else if j == m-1 {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Printf("match result:%d", BruteForceAlgo("hello", "el"))
}