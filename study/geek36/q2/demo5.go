package q2

import (
	"flag"
	lib5 "awesomelib-go/study/geek36/q2/lib"
	"awesomelib-go/study/geek36/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.HelloWorld()
	lib5.Hello(name)
}