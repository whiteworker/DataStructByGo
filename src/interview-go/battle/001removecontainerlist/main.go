package main

import (
	"container/list"
	"fmt"
)

func main() {
	//初始化一个list
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	fmt.Println("Before Removing...")
	//遍历list，删除元素
	// for e := l.Front(); e != nil; e = e.Next() {
	// 	fmt.Println("removing", e.Value)
	// 	l.Remove(e)
	// }
	var n *list.Element
	for e := l.Front(); e != nil; e = n {
		fmt.Println("removing ", e.Value)
		n = e.Next()
		l.Remove(e)
	}
}

// remove removes e from its list, decrements l.len, and returns e.
// func (l *List) remove(e *Element) *Element {
// 	e.prev.next = e.next
// 	e.next.prev = e.prev
// 	e.next = nil // avoid memory leaks
// 	e.prev = nil // avoid memory leaks
// 	e.list = nil
// 	l.len--
// 	return e
// }
