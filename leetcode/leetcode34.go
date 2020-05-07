package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{1, 3, 5, 5, 6}, 9))
}

/**
两次调用二分查找法
 */
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums) - 1
	mid := (right + left + 1) / 2
	min := -1
	for true {
		if left > right {
			break
		}
		mid = (right + left + 1) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			min = mid
			right = mid - 1
		}
	}
	if min == -1 {
		return []int{-1, -1}
	}
	max := min
	left, right = min + 1, len(nums)-1
	for true {
		if left > right {
			break
		}
		mid = (right + left + 1) / 2
		if nums[mid] > target {
			right = mid - 1
		} else {
			max = mid
			left = mid + 1
		}
	}

	return []int{min, max}
}

