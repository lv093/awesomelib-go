package main

import (
	"net"
	"fmt"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err == nil {
		fmt.Printf("hello")
	}
	for  {
		conn, _ := listener.Accept()
		handleConn(conn)
	}

}

func handleConn(c net.Conn)  {
	fmt.Println("handle")
}
