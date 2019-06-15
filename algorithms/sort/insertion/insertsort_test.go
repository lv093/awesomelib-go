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

func benchmarkInsertSort(n int, b *testing.B) {
	list := library.GenRandArray(5,1,11)
	fmt.Println("--------")

	for i := 0; i < b.N; i ++ {
		InsertSort(list)
	}
}

func Benchmark100_InsertSort(b *testing.B) {
	benchmarkInsertSort(100, b)
}

func Benchmark1000_InsertSort(b *testing.B) {
	benchmarkInsertSort(1000, b)
}

func Benchmark10000_InsertSort(b *testing.B) {
	benchmarkInsertSort(10000, b)
}

func Benchmark100000_InsertSort(b *testing.B) {
	benchmarkInsertSort(100000, b)
}

func Benchmark_Alloc(b *testing.B) {

	for i := 0; i < b.N; i++ {
		benchmarkInsertSort(1, b)
	}
}