package main

import "fmt"

func main() {
	data := []int{1, 2, 3}
	printAny((data))
}
func printAny(any []interface{}) { //删除[]
	fmt.Println(any)
}
