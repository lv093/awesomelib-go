package main

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"log"
	"net/http"
	"time"
)

func Qrcode(w http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(500)
			return
		}
	}()
	q, err := qrcode.New(fmt.Sprintf("https://www.wafunny.com?t=%d", time.Now().Unix()), qrcode.Medium)
	if err != nil {
		return
	}
	png, err := q.PNG(256)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(png)))
	w.Write(png)
}

func main() {
	http.HandleFunc("/qrcode", Qrcode)
	log.Fatal(http.ListenAndServe(":8008", nil))
}
