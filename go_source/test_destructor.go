package main
import (
	"runtime"
	"time"
)
func main() {
	x := 123
	runtime.SetFinalizer(&x, func(x *int) {
		println(x, *x, "finalizer.")
	})
	runtime.GC()
	time.Sleep(time.Second * 10)
	//runtime.GC()
	println("end")
}