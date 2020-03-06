package cl

import (
	"fmt"
)

/**
url:[抓取列表]
 */
func GrabImageHtmlList(url string) []string {
	reader, err := GetClHtmlData(url)
	if err != nil {
		fmt.Println("GrabImageHtmlList==GetImageListHtml==error:", err.Error())
		return nil
	}
	links := ParseImageListHtml(reader)
	return links
}

/**
抓取详情
 */
func GrabImageHtmlDetails(url string) (string, []string) {
	reader, err := GetClHtmlData(url)
	if err != nil {
		fmt.Println("GrabImageHtmlList==GetImageListHtml==error:", err.Error())
		return "", nil
	}
	title, links := ParseImageDetailUrls(reader)
	return title, links
}

