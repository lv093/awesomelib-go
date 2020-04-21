package main

import (
	"net/rpc"
	"net"
	"log"
	"net/http"
)

//自己的数据类
type MyMath struct{
}

//加法--只能两个参数--方法名第一个字母必须大写
func (mm *MyMath) Add(num map[string]int,reply *int) error {
	*reply = num["num1"] + num["num2"]
	return nil
}

func main() {
	//注册MyMath类，以代客户端调用
	rpc.Register(new(MyMath))
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1215")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}