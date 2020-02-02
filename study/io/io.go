package main

import (
	"fmt"
	"encoding/json"
	"bytes"
)

func main() {
	type Book struct {
		Price int
		Name string
	}
	var person = struct {
		Name string	`json:"name"`
		Price int	`json:"price"`
	}{Name: "alice", Price: 123}
	js, _ := json.Marshal(person)
	fmt.Println("json==book==>", string(js))

	j := `{"name":"alice","price":123}`
	var res = make(map[string]interface{})
	json.Unmarshal([]byte(j), &res)
	price, _ := res["price"].(int)
	fmt.Printf("unmarshal==>%f==%T", price, price)

	var p = make(map[string]interface{})
	decoder := json.NewDecoder(bytes.NewReader([]byte(j)))
	decoder.UseNumber()
	decoder.Decode(&p)
	pc := p["price"]
	fmt.Printf("unmarshal==>%f==%T \n", pc, pc)
}