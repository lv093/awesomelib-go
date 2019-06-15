package parser

import (
	"regexp"
	"awesomelib-go/crawler/engine"
	"fmt"
)

const citylistRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

//提取所需内容 如网页地址 城市名称
func PrintCityList(contents []byte)engine.ParseResult {
	//设置被提取者所需要的条件
	re := regexp.MustCompile(citylistRe)
	fmt.Println("lvfix-regexp content:%s",string(contents))
	//从网页文件中提取所需文件
	matches := re.FindAllSubmatch(contents, -1)
	//提取出来的内容需要一个接受者 result定义类型
	result := engine.ParseResult{}
	for _, m := range matches {
		fmt.Println("matches for each ,m0:", string(m[0]))
		fmt.Println("matches for each ,m1:", string(m[1]))
		fmt.Println("matches for each ,m2:", string(m[2]))
		//将提取出来的信息按照顺序放入变量result中
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			string(m[1]),
			engine.Nilparser,
		})
	}
	//返回result result内部是数组
	return result
}