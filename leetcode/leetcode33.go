package main

import (
	"fmt"
)

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
你可以假设数组中不存在重复的元素。
你的算法时间复杂度必须是 O(log n) 级别。
 */
func main() {
	sources := []int{4}
	fmt.Println(BinarySearch(sources, 2))
}

/**
利用二分法找到无旋转区域，在该有序的数组段中搜索target
 */
func BinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := false
	if target >= nums[0] {
		left = true
	}
	//寻找旋转点
	maxSize := len(nums)
	low, high := 0, maxSize - 1
	mid := (low + high) / 2
	for mid >= 0 && mid <= maxSize - 1 {
		if left {
			if nums[mid] < nums[0] {
				high = mid - 1
			} else if nums[mid] > target {
				high = mid - 1
			} else if nums[mid] < target {
				low = mid + 1
			} else {
				return mid
			}
		} else {
			if nums[mid] >= nums[0] {
				low = mid + 1
			} else if nums[mid] > target {
				high = mid - 1
			} else if nums[mid] < target {
				low = mid + 1
			} else {
				return mid
			}
		}
		if (low + high) / 2 == mid {
			return -1
		}
		mid = (low + high) / 2
	}
	return -1
}