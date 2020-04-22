package main

import (
	"strconv"
)

func main() {
	println(reverse(-123))
}

func reverse(x int) int {
	max := 2147483647
	min := -2147483648
	neg := false
	if x < 0 {
		x = 0 - x
		neg = true
	}
	init := strconv.Itoa(x)
	res := ""
	l := len(init)
	for i:=0; i < l; i ++ {
		res = res + init[l-i-1:l-i]
	}
	println(res)
	for i:=0; i < l; i++ {
		println(i)
		if res[i:i+1] != "0" {
			res = res[i:]
			break
		}
	}
	result, _ := strconv.Atoi(res)
	if neg {
		result = 0 - result
	}
	if result > max {
		return 0
	}
	if result < min {
		return 0
	}

	return result
}