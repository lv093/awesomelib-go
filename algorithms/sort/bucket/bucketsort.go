package main

import (
	"awesomelib-go/algorithms/sort/quick"
	"fmt"
	"math/rand"
)

/**
先对数据进行桶分类，然后分别对各个桶进行快排，然后组装到一起
 */
func BucketSort(targets []int, bucketSize int) []int {
	splitBuckets := make(map[int][]int, 0)
	for _, target := range targets {
		key := target/bucketSize
		splitBuckets[key] = append(splitBuckets[key], target)
	}
	fmt.Println(splitBuckets)
	fmt.Println("==========")
	sortedBuckets := make(map[int][]int, len(splitBuckets))
	for key, bucket := range splitBuckets {
		sortedBuckets[key] = quick.QuickSort(bucket, 0, len(bucket) - 1)
	}

	index := 0
	for index < bucketSize {
		if bucket, ok := sortedBuckets[index]; ok {
			fmt.Println(bucket)
		}
		index++
	}
	return nil
}

func main() {
	i := 0
	targets := make([]int, 100)
	for i <= 100 {
		targets = append(targets, rand.Intn(100))
		i++
	}
	fmt.Println(targets)
	fmt.Println("===========")
	BucketSort(targets, 10)
}