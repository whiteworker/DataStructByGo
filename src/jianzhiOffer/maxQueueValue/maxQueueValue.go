package main

import (
	"fmt"
)
type MaxQueue struct {
	queue *Queue
	maxQueue *Queue
}


func Constructor() MaxQueue {
	return MaxQueue{
		queue:NewQueue(),
		maxQueue :NewQueue(),
	}
}


func (this *MaxQueue) Max_value() int {
	return this.maxQueue.Peek()
}


func (this *MaxQueue) Push_back(value int)  {
	this.queue.EnQueue(value)
	for(this.maxQueue.Length>0&&this.maxQueue.PeekLast()<value){
		this.maxQueue.DeQueueLast()
	}
	this.maxQueue.EnQueue(value)
}


func (this *MaxQueue) Pop_front() int {
	res := this.queue.DeQueue()
	if(res==this.maxQueue.Peek()){
		this.maxQueue.DeQueue()
	}
	return res
}

type Queue struct{
	List []int
	Length int
}
func (queue *Queue)EnQueue(value int ){
	queue.List=append(queue.List,value)
	queue.Length++
}
func (queue *Queue)EnQueueLast(value int ){
	queue.List=append(queue.List,value)
	queue.Length++
}

func (queue *Queue)DeQueue() int{
	if(queue.Length>0){
		res := queue.Peek()
		queue.List =queue.List[1:]
		queue.Length--
		return res
	}
	return -1
}
func (queue *Queue)DeQueueLast() int{
	if(queue.Length>0){
		res := queue.PeekLast()
		queue.List =queue.List[:len(queue.List)-1]
		queue.Length--
		return res
	}
	return -1
}
func (queue *Queue)Peek() int{
	if(queue.Length>0){
		return queue.List[0]
	}
	return -1
}
func (queue *Queue)PeekLast() int{
	if(queue.Length>0){
		return queue.List[len(queue.List)-1]
	}
	return -1
}
func NewQueue()*Queue{
	return &Queue{
		 List: []int{},
		 Length :0,
	}
}
/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
 func main(){
	 maxQueue := Constructor()
	 maxQueue.Push_back(94)
	 maxQueue.Push_back(16)
	 maxQueue.Push_back(89)
	 fmt.Println(maxQueue.Pop_front())
	 maxQueue.Push_back(22)
	 maxQueue.Push_back(33)
	 maxQueue.Push_back(44)
	 maxQueue.Push_back(111)
	 maxQueue.Pop_front()
	 maxQueue.Pop_front()
	 maxQueue.Pop_front()
	 fmt.Println(maxQueue.Max_value())
 }