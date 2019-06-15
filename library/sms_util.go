package library

import (
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
)
//
//func main() {
//	account := "17802929676"
//	pwd := ""
//	mobile := ""
//	contentYzm := "hello"
//
//	urlSend := "www.seagullfly.cn"
//	dataSend := url.Values{"account":{account}, "pwd":{pwd}, "mobile":{mobile},"msg":{contentYzm},"needstatus":{"true"}}
//
//	httpsPostForm(urlSend, dataSend)
//}

func httpsPostForm(url string, data url.Values)  {
	resp, err := http.PostForm(url, data)
	if err != nil {
		fmt.Printf("postform error %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp failed:%s", err.Error())
	}
	fmt.Printf(string(body))
}
//
//func httpsPostForm(url string,data url.Values) {
//	    resp, err := http.PostForm(url,data)
//	    if err != nil {
//	        // handle error
//	    }
//	    defer resp.Body.Close()
//	    body, err := ioutil.ReadAll(resp.Body)
//	    if err != nil {
//	        // handle error
//	    }
//	    fmt.Println(string(body))
//}