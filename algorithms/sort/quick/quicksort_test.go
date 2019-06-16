package quick

import (
	"testing"
	"awesomelib-go/library"
	"fmt"
)

func TestQuickSort(t *testing.T) {
	arr := library.GenRandArray(10,1,20)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	QuickSort(arr, 0, len(arr) - 1)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}