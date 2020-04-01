package main

import (
	"sync"
	"fmt"
	"runtime"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go ii(i)
	}
	wg.Wait()
}

func ii(i int) {
	defer wg.Done()
	fmt.Println(i)
}