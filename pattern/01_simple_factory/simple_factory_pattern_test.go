package simple_factory

import "testing"

func TestType1(t *testing.T) {
	animal := NewAnimal("fish")
	//if animal == nil {
	//	t.Fatal("error type")
	//}
	animal.Sing("world")
}