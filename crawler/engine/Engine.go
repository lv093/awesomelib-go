package engine

import (
	"log"
	"fmt"
	"awesomelib-go/crawler/fetcher"
)

//信息存放位置 每一个信息具有单独不连续的内存
type Request struct {
	Url string   //存放地址
	ParserFunc func([]byte) ParseResult  // 存放函数类型,参数[]byte就是下载器根据这个URL返回的内容
}
//信息大的集合 注意结构体中的类型都为数组 这是一个很关键的设置
type ParseResult struct{
	Requests []Request   //存放一个或多个Request供程序使用
	Items []interface{} //存放多个参数 基本打印就靠他
}
//暂时设置为空让程序跑起来 开始收集用户信息时  它会被替换掉
func Nilparser([] byte) ParseResult {
	return ParseResult{}
}


//引擎 控制整个程序的流程
func Run(seeds ...Request) {
	var requests []Request
	//接收main函数传过来的值
	for _, r := range seeds {
		requests = append(requests, r)
	}
	//利用传过来的值进行 解析 及 提取
	for len(requests) > 0 {
		//获取第一个值
		r := requests[0]
		//进行切片 把已经提取的内容筛选出去
		requests = requests[1:]
		//第一次打印为main函数传入地址 然后每次打印是从r.ParserFunc函数中提取出的城市地址
		//将不同URL传输进去 返回不同的页面源代码
		body, err := fetcher.Fetch(r.Url)
		//判断URL是否正确 如果不正确 跳过此次循环
		if err != nil {
			log.Printf("Fetcher: error fetching url %s:%v", r.Url, err)
			continue
		}
		//注意关键的地方 r为Request结构体变量 在main函数中 我们设置Request结构体变量中的ParserFunc值为parser.PrintCityList
		//所以当第一次循环时r.ParserFunc(body)相当于parser.PrintCityList(body) 再一次体会到go语言的函数式编程魅力
		//r.ParserFunc(body)得到的结构体组分为城市名称及要执行的函数 仔细揣摩结构体Request的ParserFunc值
		//第一次循环成功后 parser.PrintCityList(body)被我们的engine.Nilparser代替 当然现在engine.Nilparser为空没有任何返回值
		ParseResult := r.ParserFunc(body)
		//requests被填满 requests又得到新的URL和运算函数 被抓取信息只要足够就可以一直运行下去
		requests = append(requests, ParseResult.Requests...)
		//打印所有在PrintCityList函数返回的Item值 Item值是任何类型可以使城市名也可以是用户信息
		for _, item := range ParseResult.Items {
			fmt.Printf("Got item %v\n", item)
		}
		fmt.Printf("engine run over\n")
	}
}