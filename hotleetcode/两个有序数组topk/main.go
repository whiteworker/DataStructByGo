package main

import "fmt"

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
func topk(arr []int, k int) int {
	if k > len(arr) {
		return -1
	}
	return arr[len(arr)-k]
}
func main() {
	res := merge([]int{1, 2, 3}, []int{4, 5, 6})
	fmt.Println(topk(res, 1))

}
