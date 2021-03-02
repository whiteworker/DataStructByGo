package main

import (
	"fmt"
)

func main() {

	list := &SingleList{val: 1, next: &SingleList{val: 2, next: &SingleList{val: 3, next: nil}}}
	list.ShowList()
	Reverse(list).ShowList()
}

type SingleList struct {
	val  int
	next *SingleList
}

func (s *SingleList) ShowList() {
	if s == nil {
		return
	}
	curr := s
	for curr != nil {
		fmt.Printf("%v", curr.val)
		curr = curr.next
	}
	fmt.Printf("\n")
}

func Reverse(head *SingleList) *SingleList {
	if head == nil {
		return nil
	}
	var curr *SingleList
	prev := head
	for prev != nil {
		temp := prev.next
		prev.next = curr
		curr = prev
		prev = temp
	}
	return curr
}
