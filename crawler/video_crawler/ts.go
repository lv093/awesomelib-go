package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)
/*
    var (
        //url = "http://flv5.bn.netease.com/live163/store/208588/serverpush_18783_1521024752481_208588_0-1.ts"
        url  string = "http://flv5.bn.netease.com/live163/store/208588/serverpush_18783_1521024752481_208588_0-"
        url_suffix string = ".ts"
    )
*/
func main() {

	fmt.Println(strings.Contains("", ""))
	os.Exit(0)
	fmt.Println("download begin.....")

	//url := "http://flv5.bn.netease.com/live163/store/208588/serverpush_18783_1521024752481_208588_0-"
	//url_suffix := ".ts"
	url_suffix := ".jpg"
	//url := "https://www.yuoimg.com/u/20200221/17220181.jpg"
	url := "https://www.privacypic.com/images/2020/03/05/1641907c5cd6ed4dd.jpg"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("17220181.jpg")
	if err != nil {
		panic(err)
	}
	fileSize, err := io.Copy(f, res.Body)
	fmt.Println(fileSize, err)
	os.Exit(1)

	for i := 180; i <= 180; i++ {

		str_index := strconv.Itoa(i)        // 通过Itoa方法转换
		//str2 := fmt.Sprintf("%d", i)      // 通过Sprintf方法转换
		//fmt.Println(str_index)   // 打印str1

		strFinal := url+str_index+ url_suffix

		res, err := http.Get(strFinal)
		if err != nil {
			panic(err)
		}
		f, err := os.Create(str_index+url_suffix)
		if err != nil {
			panic(err)
		}
		fileSize,writeErr := io.Copy(f, res.Body)

		fmt.Println(strFinal + " download done,", "file size(byte)=", fileSize)
		if writeErr != nil {
			fmt.Println(strFinal + " download failed ",  "errorInfo=", writeErr.Error())
			panic(err)
		}
	}


	fmt.Println("download finish.")

}