package main

import (
	"fmt"
)

type ListNode struct{
	Value int
	Next *ListNode 
}
func AddTwoNumber(l1 *ListNode,l2 *ListNode)(*ListNode){
	var head *ListNode
	var tail *ListNode
	carry :=0
	for(l1 !=nil||l2!=nil){
		n1,n2 := 0,0
		if(l1!=nil){
			n1=l1.Value
		}
		if(l2!=nil){
			n2=l2.Value
		}
		sum := n1+n2+carry
		if(head==nil){
			head=&ListNode{Value :sum%10 }
			tail=head
		} else{
			tail.Next=&ListNode{Value :sum%10}
			tail=tail.Next
		}
		carry = sum/10
		if(l1!=nil){
			l1=l1.Next
		}
		if(l2!=nil){
			l2=l2.Next
		}
	}
	if(carry>0){
		tail.Next=&ListNode{Value:carry}
		tail=tail.Next
	}
	return head
}
func Print(listNode *ListNode){
	current := listNode
	result := 0
	for current!=nil{
		result = current.Value +result*10
		current=current.Next
	}
	fmt.Println(result)
}
func main(){
	l1 := &ListNode{Value :2 ,Next :&ListNode{Value:4,Next :&ListNode{Value:3}}}
	l2 := &ListNode{Value :5 ,Next :&ListNode{Value:6 ,Next :&ListNode{Value:4}}}
	result :=AddTwoNumber(l1,l2)
	Print(result)
}