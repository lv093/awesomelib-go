package cl

import (
	"golang.org/x/net/html"
	"io"
	"fmt"
	"strings"
	"github.com/axgle/mahonia"
)

func ParseHtmlTitle(reader io.Reader) string {
	n, err := html.Parse(reader)
	if err != nil {
		return err.Error()
	}
	//var enc mahonia.Decoder
	//enc = mahonia.NewDecoder("gbk")

	title := ""
	title = parseHtmlTitle(n)
	return title
}

func ParseImageListHtml(reader io.Reader) []string {
	n, err := html.Parse(reader)
	if err != nil {
		fmt.Println("ParseImageDetailHtml, html parse error:", err.Error())
		return nil
	}
	links := parseHtml(nil, n, "a", "href", "htm_data")
	return links
}

func ParseImageDetailUrls(reader io.Reader) (string, []string) {
	n, err := html.Parse(reader)
	if err != nil {
		fmt.Println("ParseImageDetailHtml, html parse error:", err.Error())
		return "", nil
	}
	title := parseHtmlTitle(n)
	//fmt.Println("ParseImageDetailUrls,title==>", title)
	links := parseHtml(nil, n, "img", "data-src", "")
	return title, links
}

func parseHtml(links []string, n *html.Node, label string, key string, filter string ) []string {
	if n.Type == html.ElementNode && n.Data == label {
		for _, a := range n.Attr {
			if a.Key == key &&  strings.Contains(a.Val, filter) {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = parseHtml(links, c, label, key, filter)
	}
	return links
}


func parseHtmlTitle(n *html.Node) string {
	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	if n.Data == "title" {
		title := enc.ConvertString(n.FirstChild.Data)
		if title != "" {
			return title
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title := parseHtmlTitle(c)
		if title != "" {
			return title
		}
	}
	return ""
}