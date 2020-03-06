package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	params := map[string]string{"hello": "1"}
	value, ok := strconv.Atoi(params["alc"])
	fmt.Println(value, ok)
	fmt.Println(strconv.FormatInt(int64(65), 10))
	fmt.Println(strings.ToLower("GOalkeeper"))
	//num1 := []int{1, 2, 5}
	//num2 := []int{3,6,7}
	//fmt.Println(findMedianSortedArrays(num1, num2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	//half := (m + n) / 2
	mid1, mid2 := m/2, n/2
	if m > n {
		for  {
			if nums1[mid1 + 1] < nums2[mid2 - 1] {
				mid1++
				mid2--
			}
		}
	}
	if m + n % 2 == 0 {

	} else {

	}
	return 0
}