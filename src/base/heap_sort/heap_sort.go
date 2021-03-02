package main

import (
	"fmt"
)

func HeapSort(array []int) {
	m := len(array)
	s := m/2
	for i := s; i > -1; i-- {
		//从第一个非叶子结点从下至上，从右至左调整结构
		heap(array, i, m-1)
	}
	for i := m-1; i > 0; i-- {
		array[i], array[0] = array[0], array[i]//将堆顶元素与末尾元素进行交换
		heap(array, 0, i-1)//重新对堆进行调整
	}
}

func heap(array []int, i, end int){
	l := 2*i+1 
	if l > end {
		return
	}
	n := l
	r := 2*i+2
	if r <= end && array[r]>array[l]{
		n = r
	}
	if array[i] > array[n]{
		return
	}
	array[n], array[i] = array[i], array[n]
	heap(array, n, end)
}

func main()  {
	array := []int{
		55, 94,87,1,4,32,11,77,39,42,64,53,70,12,9,
	}
	HeapSort(array)
	fmt.Println(array)
}