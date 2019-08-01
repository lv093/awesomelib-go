package main

import (
	"net/http"
	"fmt"
	"log"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count ++
	mu.Unlock()
	fmt.Fprintf(w, "hello %s", r.URL.Path)
	fmt.Println( "hello %s", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "hello %d", count)
	fmt.Println( "hello %d", count)
	mu.Unlock()
}