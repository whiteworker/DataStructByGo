package main

import "fmt"

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 {
		for _, value := range stack1 {
			stack2 = append(stack2, value)
		}
		stack1 = []int{}
	}

	if len(stack2) > 0 {
		res := stack2[0]
		stack2 = stack2[1:len(stack2)]
		return res
	}
	return -1
}
func main() {
	Push(1)
	Push(2)
	Push(3)
	fmt.Println(Pop())
}
