package main

import (
	"sort"
	"go/src/fmt"
)

func RadixIndexSort(targets []string, index int) []string {

	return targets
}

func main() {
	targets := []string{"hell", "world"}
	sort.Sort(targets)
	RadixSort(targets, 1)
}
/**
基数排序
@params tailSupple
如果未0，不填充，1-尾部填充，2-头部填充
 */
func RadixSort(targets []string, tailSupple int) []string {
	if len(targets) <= 1 {
		return targets
	}
	//填充
	maxSize := 0
	for _, target := range targets {
		if maxSize < len(target) {
			maxSize = len(target)
		}
	}
	if tailSupple == 1 {
		for i, target := range targets {
			if len(target) < maxSize {
				j := 1
				for j <= 2 {
					targets[i] = targets[i] + "0"
					j++
				}
			}
		}
	} else if tailSupple == 2 {
		for i, target := range targets {
			if len(target) < maxSize {
				j := 1
				for j <= 2 {
					targets[i] = "0" + targets[i]
					j++
				}
			}
		}
	}
	//排序
	index := maxSize - 1
	for index >= 0 {
		wrapper := SortWrapper{Targets:targets, Index:index}
		sort.Sort(wrapper)
		fmt.Println(wrapper)
		index--
	}
	return targets
}

/**
对string数组的index位置进行排序
 */
type SortWrapper struct {
	Targets []string
	Index 	int
}
func (this SortWrapper) Len() int {
	return len(this.Targets)
}
func (this SortWrapper) Swap(i, j int) {
	this.Targets[i], this.Targets[j] = this.Targets[j], this.Targets[i]
}
func (this SortWrapper) Less(i, j int) bool {
	if this.Targets[i][this.Index] < this.Targets[j][this.Index] {
		return true
	}
	return false
}

