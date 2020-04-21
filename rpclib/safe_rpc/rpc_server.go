package main

import (
	"net/rpc"
	"net"
	"awesomelib-go/rpclib/safe_rpc/hello"
)

type HelloServiceInterface interface {
	SaidHello(request string, reply *string) error
}

func RegisterHelloService(service HelloServiceInterface) error {
	return rpc.RegisterName("safe_rpc/hello/HelloAliceService", service)
}


func main() {
	aliceSvc := new(hello.HelloAliceService)
	RegisterHelloService(aliceSvc)
	listener, err := net.Listen("tcp", ":9991")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		println(conn)
		if err != nil {
			println(err.Error())
			continue
		}
		rpc.ServeConn(conn)
	}
}