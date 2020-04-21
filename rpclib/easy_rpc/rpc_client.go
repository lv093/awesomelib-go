package main

import (
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:9900")
	if err != nil {
		panic(err)
	}
	reply := "init"
	//err = client.Call("HelloService.SayHello", "alice", &reply)

	call := client.Go("HelloService.SayHello", "alice", &reply, nil)
	 <-call.Done
	//println(res.Reply)

	//time.Sleep(time.Second)
	if err != nil {
		panic(err)
	}
	println(reply)
}