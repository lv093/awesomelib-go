
package main

import (
	"awesomelib-go/crawler/engine"
	"awesomelib-go/crawler/parser"
)

func main() {
	//运行爬虫的起始条件 当内部循环一次结束后 里面信息属于无效
	engine.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parser.PrintCityList,
	})
}