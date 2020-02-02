package main

import (
	"net/http"
	"io/ioutil"
	"log"
)

func echo(wr http.ResponseWriter, request *http.Request)  {
	msg, err := ioutil.ReadAll(request.Body)
	if err != nil {
		wr.Write([]byte("failed"))
		return
	}
	writeLen, err := wr.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len:", writeLen)
	}
}

func main() {
	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}