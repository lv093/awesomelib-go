package main

import (
	"net/rpc"
	"net"
)

type HelloService struct {
}

func (this *HelloService) SayHello(req string, reply *string) error {
	*reply = "hello:" + req
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	listerner, err := net.Listen("tcp", ":9900")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listerner.Accept()
		if err != nil {
			panic(err)
		}
		rpc.ServeConn(conn)
	}
}

