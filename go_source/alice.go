package main

import (
	"awesomelib-go/go_source/lib"
)

func init() {
	println("test.init")
}
func test() {
	println(lib.Sum(1, 2, 3))
}
