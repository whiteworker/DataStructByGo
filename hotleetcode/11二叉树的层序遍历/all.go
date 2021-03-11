package main

import (
	"container/list"
)

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
	res := make([]int, 0)
	curr := root
	stack := list.New()
	for curr != nil || stack.Len() != 0 {
		for curr != nil {
			res = append(res, curr.Val) //visit
			stack.PushBack(curr)
			curr = curr.Left
		}
		for stack.Len() != 0 {
			r := stack.Back()
			curr := r.Value.(*TreeNode)
			curr = curr.Right
			stack.Remove(r)
		}
	}
	return res
}
func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

func postorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, res)
	postorder(root.Right, res)
	*res = append(*res, root.Val)
}
