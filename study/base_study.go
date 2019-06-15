package main

import (
	"reflect"
	"fmt"
	"strings"
	"os"
)

func main() {
	PrintArgs1(os.Args)
	fmt.Println("++++++++++++++++")
	printOs(os.Args)
}

//变参函数的定义方式
func PrintArgs1(args ...interface{}) {
	fmt.Println(args[0].([]string))

	fmt.Println(reflect.TypeOf(args[0]))
	for k, v := range args[0].([]string) {
		fmt.Println(k, " =", v, reflect.TypeOf(v))
	}
	fmt.Println("========================================")
	var myStr = (args[0].([]string))[1]	 //直接点击调试运行,此行代码数组就越界了
	fmt.Println("myStr =", myStr, "myStr Type =", reflect.TypeOf(myStr))
	ss := strings.Split(myStr, ",")
	for k, v := range ss {
		fmt.Println(k, " =", v, reflect.TypeOf(v))
	}
}

func printOs(args... interface{})  {
	for key,value := range args[0].([]string) {
		fmt.Println("arg:",key,value)
	}
}
