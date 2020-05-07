package main

import (
	"fmt"
	"sync"
)

func main() {
	//maps := map[int]int{4: 111, 3:112, 5:113, 1:114}
	//for k, v := range maps {
	//	fmt.Println(k, v)
	//}
	init := []int{1, 3, 5, 4,1, 7}
	fmt.Println(init, 2)
	after := moveToEnd(init, 2)
	fmt.Print(after)
}

type CacheItem struct {
	Key 	int
	Val 	int
	Index 	int
}

type LRUCache struct {
	Lock 		sync.RWMutex
	Capacity 	int
	Keys		[]int
	Values 		map[int]CacheItem
}


func Constructor(capacity int) LRUCache {
	cache := new(LRUCache)
	cache.Capacity = capacity
	cache.Keys = make([]int, 0)
	cache.Values = make(map[int]CacheItem, 0)
	return *cache
}


func (this *LRUCache) Get(key int) int {
	this.Lock.Lock()
	defer this.Lock.Unlock()
	if val, ok := this.Values[key]; ok {
		idx := val.Index
		this.Keys = moveToEnd(this.Keys, idx)

		val.Index = len(this.Keys) - 1
		this.Values[key] = val
		return val.Val
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	if val, ok := this.Values[key]; ok {
		idx := val.Index
		this.Keys = moveToEnd(this.Keys, idx)

		val.Index = len(this.Keys) - 1
		val.Val = value
		this.Values[key] = val
	} else {
		if len(this.Keys) >= this.Capacity {
			this.Keys = append(this.Keys[1:], key)
		} else {
			this.Keys = append(this.Keys, key)
		}
		val := CacheItem{key, value, len(this.Keys) - 1}
		this.Values[key] = val
	}
}

func moveToEnd(value []int, pos int) []int {
	if pos >= len(value)-1 {
		return value
	}
	if pos == 0 {
		value = append(value[1:], value[0])
	} else {
		target := value[pos]
		value = append(value[0:pos], value[pos+1:]...)
		value = append(value, target)
	}
	return value
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */