package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//注意字段必须是导出
type Params struct {
	Width, Height int
}
type Rect struct{}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}
func main() {
	rect := new(Rect)
	//注册rpc服务
	rpc.Register(rect)
	//监听端口
	tcplisten, _ := net.Listen("tcp", ":8080")
	for {
		conn, err3 := tcplisten.Accept()
		if err3 != nil {
			continue
		}
		//这里使用jsonrpc进行处理
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		//go jsonrpc.ServeConn(conn)
	}
}