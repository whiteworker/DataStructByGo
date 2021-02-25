package main

import (
	"encoding/json"
	"fmt"
)

func test() {
	var jsonBlob = []byte(`{"name": "Platypus", "order": "Monotremata","id":1}`)
	fmt.Println("jsonBlob: ", string(jsonBlob))
	type Animal struct {
		id    int32  `json:"id,omitempty"`
		name  string `json:"name,omitempty"`
		order string `json:"order,omitempty"`
	}
	var animals Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("animals: %+v\n", animals)
}
func testTrue() {
	//结构体必须是大写字母开头的成员才会被JSON处理到，小写字母开头的成员不会有影响。
	var jsonBlob = []byte(`{"name": "Platypus", "order": "Monotremata","id":1}`)
	fmt.Println("jsonBlob: ", string(jsonBlob))
	type Animal struct {
		Id    int32  `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Order string `json:"order,omitempty"`
	}
	var animals Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("animals: %+v\n", animals)
}
func main() {
	test()
	testTrue()
}
