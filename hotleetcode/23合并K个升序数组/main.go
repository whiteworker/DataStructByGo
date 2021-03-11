package main

import (
	"fmt"
)

func merge(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	res := make([]int, (m + n))
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums2[p2] > nums1[p1] {
			res[p] = nums2[p2]
			p2--
		} else {
			res[p] = nums1[p1]
			p1--
		}
		p--
	}
	if p2 < 0 {
		copy(res[:p1+1], nums1[:p1+1])
	} else {
		copy(res[:p2+1], nums2[:p2+1])
	}

	return res
}

// Time: O(k*n), Space: O(1) n是所有链表的节点总数，k是链表个数
func mergeKLists(lists [][]int) []int {
	if lists == nil || len(lists) == 0 {
		return nil // 链表为空或长度为0，直接返回空指针
	}
	var result []int             // 否则定义合并后的结果链表
	for _, list := range lists { // 定义链表数组
		// 和数组中的链表一个个合并
		result = merge(result, list)
	}
	return result // 最后返回合并后的链表2
}
func main() {
	fmt.Println(mergeKLists([][]int{[]int{1, 2, 3}, []int{4, 10}, []int{1, 5, 7, 8, 9}}))
	fmt.Println(mergeKLists([][]int{[]int{1, 2, 3}, []int{4, 10}, []int{1, 5, 7, 8, 9}}))
}

// // Time: O(n*log(k)), Space: O(k) n是总节点数量，k是链表个数
// func mergeKLists2(lists [][]int) []int {
// 	if lists == nil || len(lists) == 0 {
// 		return nil // 链表为空或长度为0，直接返回空指针
// 	}
// 	var h IntHeap // 最小堆用于维护当前k个节点
// 	heap.Init(&h) // 用于节点间的比较

// 	for _, list := range lists {
// 		// 数组中非空的链表加入到最小堆
// 		if list != nil {
// 			heap.Push(&h, list)
// 		}
// 	}
// 	// 定义dummy节点用于统一处理
// 	res := []int{}
// 	index := 0

// 	// 当最小堆不为空时，不断执行以下操作
// 	for h.Len() > 0 { // 取出堆顶元素，即取出最小值节点
// 		min := heap.Pop(&h).([]int)
// 		index++
// 		res[index] = min // 游标p指向最小值节点
// 		index++          // p向后移动一个位置
// 		// 这样就确定一个节点在合并链表中的位置
// 		if min.Next != nil { // 如果最小值节点后面的节点非空
// 			heap.Push(&h, min.Next) // 则把最小值节点后面的节点加入到最小堆中
// 		}
// 	}
// 	return res // 最后只要返回dummy.Next即可
// }

// type IntHeap [][]int

// func (h IntHeap) Len() int            { return len(h) }
// //func (h IntHeap) Less(i, j int) bool  { return h.([]int)[i] < h.([]int)[j] }
// func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
// func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
// func (h *IntHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }
