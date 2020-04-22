package main

import (
	"fmt"
	"awesomelib-go/algorithms/sort/quick"
	"awesomelib-go/algorithms/search"

	"strconv"
)

func main() {

	target := []int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}
	fmt.Println(quick.QuickSort(target, 0, len(target)-1))
	fmt.Println(threeSum(target))
}


func threeSum(nums []int) [][]int {
	l := len(nums)
	nums = quick.QuickSort(nums, 0, l - 1)
	res := make([][]int, 0)
	exist := make(map[string]int, 0)
	for i:=0; i < l; i++ {
		for j:=i+1; j < l-1; j++ {
			target := 0-nums[i]-nums[j]
			pos := search.BinarySearch(nums[j+1:], target)
			if pos >= 0 {
				key := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(target)
				if _, ok := exist[key]; ok {
					continue
				}
				exist[key] = 1
				res = append(res, []int{nums[i], nums[j], target})
			}
		}
	}
	return res
}