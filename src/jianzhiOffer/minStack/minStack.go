package main

import (
	"fmt"
)

type MinStack struct {
	stack *Stack
	minStack *Stack
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{ 
		stack : NewStack(),
		minStack : NewStack(),
	 }
}


func (this *MinStack) Push(x int)  {
	this.stack.Push(x)
	if(this.minStack.Length==0||this.minStack.Peek()>=x){
		this.minStack.Push(x)
	} 
}


func (this *MinStack) Pop()  {
	if(this.stack.Pop()==this.minStack.Peek()){
		this.minStack.Pop()
	}
	
}


func (this *MinStack) Top() int {
	return this.stack.Peek()
	
}


func (this *MinStack) Min() int {
	return this.minStack.Peek()
}


type Stack struct{
	List []int 
	Length int   
}
func NewStack()*Stack{
	return &Stack{
		List : []int{},
		Length : 0 ,
	}
}
func (stack *Stack)Push(value int){
	stack.List=append(stack.List,value)
	stack.Length++
}
func (stack *Stack)Pop()int{
	if(stack.Length>0){
		res := stack.List[stack.Length-1]
		stack.List=stack.List[:stack.Length-1]
		stack.Length--
		return res
	} else{
		return -1
	}
}
func (stack *Stack)Peek()int{
	if(stack.Length>0){
		res := stack.List[stack.Length-1]
		return res
	} else{
		return -1
	}
}

func main(){
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
    stack.Push(-1)
    
	fmt.Printf("stack:%v,minStack:%d",stack.stack,stack.minStack)
	stack.Pop()
	fmt.Printf("stack:%v,minStack:%d",stack.stack,stack.minStack)
}