package fetcher

import (
	"golang.org/x/net/html/charset"
	"bufio"
	"io"
	"io/ioutil"
	"golang.org/x/text/transform"
	"fmt"
	"net/http"
	"golang.org/x/text/encoding"
)

type Fetcher struct {

}

func Fetch(url string)([]byte, error){
	//试探网页能否正常打开 若能则将网页内容以结构体指针方式返回
	resp , err :=http.Get(url)
	if err != nil{
		return nil, err
	}
	//程序结束时实现结构体指针关闭
	defer resp.Body.Close()
	//判断头部内容是否正确
	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("Error Statuscode: %d", resp.StatusCode)
	}
	//自动判断字符格式函数 详细讲解在函数体内
	e := DetermineEncoding(resp.Body)
	//将从网页中获取的结构体放入函数并告诉函数结构体内的字符格式 返回utf8格式
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	//将结构体内容返回
	out,er := ioutil.ReadAll(utf8Reader)
	//fmt.Println("lvfetch content:",string(out))
	return  out,er
}
//判断字符格式并返回
func DetermineEncoding(r io.Reader)  encoding.Encoding {
	//提取结构体内的前1024个字符
	byte, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	//比较提取出来的字符进行判断
	e, _, _ := charset.DetermineEncoding(byte, "")
	//返回判断值
	return e

}