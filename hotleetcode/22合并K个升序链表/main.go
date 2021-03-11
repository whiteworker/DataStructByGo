package main

import (
	"container/heap"
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

// 辅助函数用于合并两个有序链表
func mergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	// 连上l1剩余的链,连上l2剩余的链
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dummy.Next
}

// Time: O(k*n), Space: O(1) n是所有链表的节点总数，k是链表个数
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil // 链表为空或长度为0，直接返回空指针
	}
	var result *ListNode         // 否则定义合并后的结果链表
	for _, list := range lists { // 定义链表数组
		// 和数组中的链表一个个合并
		result = mergeTwoSortedLists(result, list)
	}
	return result // 最后返回合并后的链表2
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	list2 := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}
	list3 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}
	mergeKLists([]*ListNode{list1, list2, list3}).Print()
	list4 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	list5 := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}
	list6 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}
	mergeKLists2([]*ListNode{list4, list5, list6}).Print()
}

// Time: O(n*log(k)), Space: O(k) n是总节点数量，k是链表个数
func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil // 链表为空或长度为0，直接返回空指针
	}
	var h IntHeap // 最小堆用于维护当前k个节点
	heap.Init(&h) // 用于节点间的比较

	for _, list := range lists {
		// 数组中非空的链表加入到最小堆
		if list != nil {
			heap.Push(&h, list)
		}
	}
	// 定义dummy节点用于统一处理
	dummy := &ListNode{}
	p := dummy // p初始指向dummy节点

	// 当最小堆不为空时，不断执行以下操作
	for h.Len() > 0 { // 取出堆顶元素，即取出最小值节点
		min := heap.Pop(&h).(*ListNode)
		p.Next = min // 游标p指向最小值节点
		p = p.Next   // p向后移动一个位置
		// 这样就确定一个节点在合并链表中的位置
		if min.Next != nil { // 如果最小值节点后面的节点非空
			heap.Push(&h, min.Next) // 则把最小值节点后面的节点加入到最小堆中
		}
	}
	return dummy.Next // 最后只要返回dummy.Next即可
}

type IntHeap []*ListNode

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
