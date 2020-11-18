package main
import (
	"fmt"
	"container/list"
)

type CQueue struct {
    stack1, stack2 *Stack
}

func Constructor() CQueue {
    return CQueue{
        stack1: NewStack(),
        stack2: NewStack(),
    }
}

func (this *CQueue) AppendTail(value int)  {
    this.stack1.Push(value)
}

func (this *CQueue) DeleteHead() int {
    // 如果第二个栈为空
    if this.stack2.Len() == 0 {
        for this.stack1.Len() > 0 {
            this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
        }
    }
    if this.stack2.Len() != 0 {
        e := this.stack2.Back()
        this.stack2.Remove(e)
        return e.Value.(int)
    }
    return -1
}
func main(){
    test :=Queue{}
	queue := Constructor()
	queue.AppendTail(1)
	queue.AppendTail(2)
	queue.AppendTail(3)
	queue.DeleteHead()
	fmt.Printf("stack1:%v stack2:%v",queue.stack1,queue.stack2)
}

type Stack struct {
    List []int
    Length int
}
func NewStack()*Stack{
    return &Stack{
        List: make([]int,0),
        Length: 0,
    }
}
func Push(s *Stack)(num int ){
    s.List=append(s.List,num)
}
func Pop(stack *Stack)int{
    length := len(stack.List)
    if length>0{
        lastIndex := length-1;
        res := stack.List[lastIndex]
        stack.List=stack.List[:lastIndex-1]
    } else{
        return -1
    }
    
}

