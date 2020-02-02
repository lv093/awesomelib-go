package main

import (
	"fmt"
	"reflect"
)

func Sprint(target interface{}) string {
	type Stringx interface {
		String() string
	}
	switch tar := target.(type) {
	case Stringx:
		return "Stringx"
	case string:
		return "string"
	case int:
		return "int"
	case bool:
		if tar {
			return "true"
		} else {
			return "false"
		}
	default:
		return "??"
	}
}

func main() {
	inter := new(string)
	fmt.Println("inter type :%v", *inter)
	fmt.Println(reflect.TypeOf(*inter))
}