package main

import (
	"container/list"
)

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组 https://blog.csdn.net/wade3015/article/details/106897792
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func threeOrders(root *TreeNode) [][]int {
	// write code here
	pre := preorder(root)
	in := inorder(root)
	post := []int{}
	postorder(root, &post)
	return [][]int{pre, in, post}
}
func preorder(root *TreeNode) []int {
	res := []int{}
	curr := root
	stack := list.New()
	for curr != nil || stack.Len() != 0 {
		for curr != nil {
			res = append(res, curr.Val) //visit
			stack.PushBack(curr)
			curr = curr.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			curr = v.Value.(*TreeNode)
			curr = curr.Right
			stack.Remove(v)
		}
	}
	return res
}
func inorder(root *TreeNode) []int {
	curr := root
	stack := list.New()
	res := make([]int, 0)
	for curr != nil || stack.Len() != 0 {
		for curr != nil {
			stack.PushBack(curr)
			curr = curr.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			curr = v.Value.(*TreeNode)
			res = append(res, curr.Val) //visit
			curr = curr.Right
			stack.Remove(v)
		}
	}
	return res

}
func postorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, res)
	postorder(root.Right, res)
	*res = append(*res, root.Val)
}
