package main
import (
	"fmt"
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
    if this.stack2.Length == 0 {
        for this.stack1.Length > 0 {
            this.stack2.Push(this.stack1.Pop())
        }
    }
    if this.stack2.Length != 0 {
        e := this.stack2.Pop()
        return e
    }
    return -1
}
func main(){
	queue := Constructor()
	queue.AppendTail(1)
	queue.AppendTail(2)
	queue.AppendTail(3)
    queue.DeleteHead()
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
func (s *Stack)Push(num int ){
    s.List=append(s.List,num)
    s.Length++
}
func (stack *Stack)Pop() int{
    if stack.Length>0{
        res := stack.List[stack.Length-1]
        stack.List=stack.List[:stack.Length-1]
        stack.Length--
        return res
    } else{
        return -1
    }
}