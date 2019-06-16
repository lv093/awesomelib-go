package quick

import (
)

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