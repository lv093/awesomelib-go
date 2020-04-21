package main

import (
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:9991")
	if err != nil {
		panic(err)
	}
	reply := ""
	err = client.Call("safe_rpc/hello/HelloAliceService.NotSaidHello", "from safe_rpc/rpc_client.", &reply)
	if err != nil {
		panic(err)
	}
	println(reply)
}