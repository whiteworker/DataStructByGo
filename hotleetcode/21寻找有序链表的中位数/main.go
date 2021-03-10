package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}
	fmt.Println(findmiddle(list1))
}
func findmiddle(head *ListNode) float64 {
	if head == nil {
		return -1
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast.Next == nil {
		return float64(slow.Val)
	}
	return float64((slow.Val + slow.Next.Val)) / 2

}
