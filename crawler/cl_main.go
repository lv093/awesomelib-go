package main

import (
	"fmt"
	"awesomelib-go/crawler/cl"
	"awesomelib-go/crawler/crawler_utils"
	"os"
	"strings"
)

func main() {

	args := os.Args
	fmt.Println(args)
	os.Exit(0)
	initUrl := args[1]
	initPath := args[2]
	result := cl.GrabImageHtmlList(initUrl)
	for _, detailUrl := range result {
		title, imgUrls := cl.GrabImageHtmlDetails(crawler_utils.CL_DOMAIN + detailUrl)
		fmt.Println("title==>", title)
		//创建title文件夹
		path := initPath + title
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}

		for _, img := range imgUrls {
			names := strings.Split(img, "/")
			name := img
			if len(names) > 0 {
				name = names[len(names) - 1]
			}
			_, err := cl.DownloadImg(img, path + "/" + name)
			if err != nil {
				fmt.Println("DownloadImg=error:", img, err)
			}
		}

	}
}
