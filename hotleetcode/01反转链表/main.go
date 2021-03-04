package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	if l == nil {
		return
	}
	curr := l
	for curr != nil {
		fmt.Printf("%v", curr.Val)
		curr = curr.Next
	}
}
func reverseList(head *ListNode) *ListNode {
	var curr *ListNode
	prev := head
	for prev != nil {
		temp := prev.Next
		prev.Next = curr
		curr = prev
		prev = temp
	}
	return curr
}
func main() {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	res := reverseList(list)
	res.Print()
}
