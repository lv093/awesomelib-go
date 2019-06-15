package insertion

import (
	"testing"
	"fmt"
	"awesomelib-go/library"
)

func TestInsertSort(t *testing.T) {
	list := library.GenRandArray(5,1,11)
	for i := 0; i < len(list); i ++ {
		fmt.Println(list[i])
	}

	fmt.Println("--------")
	InsertSort(list)

	for i := 0; i < len(list); i ++ {
		fmt.Println(list[i])
	}

}