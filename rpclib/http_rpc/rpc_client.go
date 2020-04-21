package main

import (
	"net/rpc"
	"fmt"
	"log"
)

func main() {
	//连接服务
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1215")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply int
	var num = make(map[string]int)
	num["num1"] = 3
	num["num2"] = 2
	//调用远程MyMath的Add方法,也只能是三个参数
	err = client.Call("MyMath.Add",num,&reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	//输出结果
	fmt.Println(reply)
	client.Close()
}