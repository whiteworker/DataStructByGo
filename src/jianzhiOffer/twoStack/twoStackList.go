package main
import (
	"fmt"
	"container/list"
)

type CQueue struct {
    stack1, stack2 *list.List
}

func Constructor() CQueue {
    return CQueue{
        stack1: list.New(),
        stack2: list.New(),
    }
}

func (this *CQueue) AppendTail(value int)  {
    this.stack1.PushBack(value)
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
	queue := Constructor()
	queue.AppendTail(1)
	queue.AppendTail(2)
	queue.AppendTail(3)
	queue.DeleteHead()
	fmt.Printf("stack1:%v stack2:%v",queue.stack1,queue.stack2)
}

