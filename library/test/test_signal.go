package socket

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	signalChan := make(chan os.Signal)

	// 监听指定信号
	signal.Notify(
		signalChan,
		syscall.SIGHUP,
		syscall.SIGUSR2,
	)

	// 输出当前进程的pid
	fmt.Println("pid is: ", os.Getpid())

	// 处理信号
	for {
		sig := <-signalChan
		fmt.Println("get signal: ", sig)
	}
}