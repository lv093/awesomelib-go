package main

import "fmt"

type Cat struct {
	Name string
}

func (this Cat) Call(name string) string {
	fmt.Println("Cat", this.Name)
	return name
}

type Dog struct {
	Name string
}

func (this Dog) Call(name string) string {
	fmt.Println("Dog", this.Name)
	return name
}


type Animal interface {
	Call(name string) string
}

func NewAnimal(name string) Animal {
	if name == "dog" {
		return &Dog{name}
	} else {
		return &Cat{name}
	}
}


func main() {
	animal := NewAnimal("dog")
	animal.Call("world")
}

