package main
import (
	"fmt"
	"log"
	"net/http"
	"strings"
)
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join((v), ""))
	}
	fmt.Println(w, "hello world")
	w.Write([]byte("end"))
}

func sayhelloName2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join((v), ""))
	}
	fmt.Println(w, "hello world")
	w.Write([]byte("hello2"))
}

func main() {
	http.HandleFunc("/hello/", sayhelloName2)
	http.HandleFunc("/hello", sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}