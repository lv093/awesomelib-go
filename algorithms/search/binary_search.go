package search


/**
二分查找的四种变体——几种变体的解决方法类似
变体一：查找第一个值等于给定值的元素
变体二：查找最后一个值等于给定值的元素
变体三：查找第一个大于等于给定值的元素
变体四：查找最后一个小于等于给定值的元素
 */

/**
二分查找标准
 */
func BinarySearch(sources []int, target int) int {
	if len(sources) <= 0 {
		return -1
	}
	low, high := 0, len(sources) - 1
	mid := (low + high) / 2
	for low <= high {
		if sources[mid] > target {
			high = mid - 1
		} else if sources[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
		mid = (low + high) / 2
	}
	return -1
}

/**
变体一：查找第一个值等于给定值的元素
 */
func BinarySearchFirstEquals(sources []int, target int) int {
	if len(sources) <= 0 {
		return -1
	}
	low, high := 0, len(sources) - 1
	mid := (low + high) / 2
	for low <= high {
		if sources[mid] > target {
			high = mid - 1
		} else if sources[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || sources[mid-1] != target {
				return mid
			} else {
				high = mid - 1
			}
		}
		mid = (low + high) / 2
	}
	return -1
}
