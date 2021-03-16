package main

import "fmt"

func main() {
	testmap()
}
func testmap() {
	m := make(map[string]map[string]int)
	m["a"] = make(map[string]int)
	m["a"]["a"] = 1
	fmt.Printf("%v", m)
	m["a"]["a"] = 2
	fmt.Printf("%v", m)
}

type A struct {
	a int
	b string
	c *A
	B
}
type B struct {
	a int
}

func teststruct() {
	m := make(map[string]A)
	b := A{a: 1, b: "AA", c: &A{a: 2, b: "BB"}}
	m["a"] = b
	b.c.a = 3
	b.B = B{a: 6}
	fmt.Println(b, *b.c)
	fmt.Println(m["a"], *m["a"].c)
}
func testarray() {
	m := make(map[string][3]int)
	b := [3]int{0, 0, 0}
	m["a"] = b
	b[0] = 6
	fmt.Println(b)
	fmt.Println(m["a"])
}
func testslice() {
	m := make(map[string][]int)
	m["a"] = make([]int, 3)
	m["a"][0] = 7
	fmt.Println(m["a"])
}
func testStr() {
	m := make(map[string]*string)
	str := "123"
	m["a"] = &str
	b := m["a"]
	*b = "223"
	fmt.Println(b, *b)
	fmt.Println(m["a"], *m["a"])
}
