package main

import "fmt"

func HeapSort(arr []int) []int {
	BuildMaxHeap(arr, len(arr))
	heapSize := len(arr)
	for i:= len(arr) - 1; i>=1; i-- {
		fmt.Printf("before buildmaxheap,pos:%d,arr:%v\n", i,arr)
		arr[i], arr[0] = arr[0], arr[i]
		heapSize--;
		maxHeapify(arr, 0, heapSize)
	}
	return  arr
}

func BuildMaxHeap(arr []int, heapSize int) []int {
	for i:= heapSize/2; i >= 0; i-- {
		maxHeapify(arr, i, heapSize)
	}
	return arr
}

/**
维护堆：对数组arr进行维护堆处理
 */
func maxHeapify(arr []int, pos int, heapSize int) {
	left := 2 * pos + 1
	right := 2 * pos + 2

	largest := pos
	if left < heapSize && arr[left] > arr[pos] {
		largest = left
	}
	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}
	if largest != pos {
		arr[largest], arr[pos] = arr[pos], arr[largest]
		maxHeapify(arr, largest, heapSize)
	}
}

func main() {
	res := HeapSort([]int{1,3,9,5,4})
	fmt.Printf("heap result:%v", res)
}