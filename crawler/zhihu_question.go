package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ZhiHuModel struct {
	QuestionTitle string
	QuestionContent string
	HeadUrl string
	Author string
	Voters string
	Content string
	CreateTime string
	CommentCount string
}

func Crawling() {
	var zeus []*ZhiHuModel
	for i := 321450693; i > 321450680; i-- {
		url := "https://www.zhihu.com/question/" + strconv.Itoa(i)
		res, err := http.Get(url)
		if err != nil || res.StatusCode != 200 {
			continue
		}
		fmt.Println("valid== %s", url)

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var zhihu ZhiHuModel
		doc.Find(".QuestionHeader .QuestionHeader-content .QuestionHeader-main").Each(func(i int, s *goquery.Selection) {
			questionTitle := s.Find(".QuestionHeader-title").Text()
			questionContent := s.Find(".QuestionHeader-detail").Text()
			questionContent = questionContent[0 : len(questionContent)-12]

			zhihu.QuestionTitle = questionTitle
			zhihu.QuestionContent = questionContent
		})

		if zhihu.QuestionTitle == ""{
			continue
		}

		doc.Find(".ContentItem-actions").Each(func(i int, s *goquery.Selection) {

		})
		doc.Find(".ListShortcut .List .List-item ").Each(func(i int, s *goquery.Selection) {
			head_url, _ := s.Find("a img").Attr("src")
			author := s.Find(".AuthorInfo-head").Text()
			zhihu.HeadUrl = head_url
			zhihu.Author = author

			voters := s.Find(".Voters").Text()
			voters = strings.Split(voters, " ")[0]
			content := s.Find(".RichContent-inner").Text() //带标签的可以用Html()
			createTime := s.Find(".ContentItem-time").Text()
			createTime = strings.Split(createTime, " ")[1]

			commentCount := s.Find(".ContentItem-actions span").Text()
			zhihu.Voters = voters
			zhihu.Content = content
			zhihu.CreateTime = createTime
			zhihu.CommentCount = commentCount
		})
		zeus = append(zeus, &zhihu)
	}
	fmt.Println(JsonString(zeus))
}

func JsonString(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(b)
}

func main() {
	Crawling()
}