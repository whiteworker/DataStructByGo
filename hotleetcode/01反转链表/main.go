package main

import (
	"encoding/json"
	"fmt"
)

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
	fmt.Printf("%v\n", res)
	fmt.Printf("%+v\n", res)
	fmt.Printf("%#v\n", res)
	s, _ := json.Marshal(res)
	fmt.Printf("%s\n", string(s))
	s, _ = json.MarshalIndent(res, "", "\t")
	fmt.Printf("%s\n", string(s))
	res.Print()
}
func reverseBetween(head *ListNode, left, right int) *ListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
