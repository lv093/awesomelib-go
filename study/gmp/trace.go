package main

import (
	"fmt"
)

func main() {
	rounds := []int{1,2,3,4}
	last := make([]int, 0)
	for k := range rounds{
		fmt.Println("k:" , k, rounds[k])
		last = append(last, rounds[k])
	}
	fmt.Println("last==>", last)

}