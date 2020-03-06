package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("", "."))
}

/**
1.首先找到小于target的最大值下标pos
2.然后在[0. pos]间压缩寻找
 */
func twoSum(nums []int, target int) []int {
	max := len(nums) - 1
	origin := make(map[int]int, max + 1)
	another := make(map[int]int, max + 1)
	for i, value := range nums {
		if _, ok := origin[value]; ok {
			another[value] = i
		} else {
			origin[value] = i
		}

	}
	QuickSort(nums, 0, max)
	low, high := 0, max
	for low != high {
		if nums[low] + nums[high] > target {
			high--
		} else if nums[low] + nums[high] < target {
			low++
		} else {
			break
		}
	}
	fmt.Println(nums, low, high)
	res := make([]int, 2)
	if value, ok := origin[nums[low]]; ok {
		res[0] = value
	}
	if nums[low] == nums[high] {
		if value, ok := another[nums[high]]; ok {
			res[1] = value
		}
	} else {
		if value, ok := origin[nums[high]]; ok {
			res[1] = value
		}
	}
	return res
}

func partition(arr []int, p int, q int) int {
	temp := arr[q]
	i := p - 1
	j := p
	for ;j < q;j++ {
		if arr[j] <= temp {
			i = i + 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i + 1], arr[q] = arr[q], arr[i + 1]
	return i + 1
}

func QuickSort(arr []int, low int, high int) []int {
	if low < high {
		middle := partition(arr, low, high)
		QuickSort(arr, low, middle - 1)
		QuickSort(arr, middle + 1, high)
	}
	return arr
}