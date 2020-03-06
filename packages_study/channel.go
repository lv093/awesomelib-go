package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	go func() {
		i := 0
		for i < 5 {
			ch<- i
			i++
			time.Sleep(time.Second)
		}
	}()
	j := 0
	for {
		fmt.Println("for:", j, time.Now().Second())
		select {
		case c, ok := <-ch:
			fmt.Println(c, ok)
		case d, ok := <-ch:
			fmt.Println("d",d, ok)
		case <- time.After(time.Second):
			fmt.Println("after second")
		}
		j++
	}
}
