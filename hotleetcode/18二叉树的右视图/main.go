package main

import "fmt"

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		res = append(res, queue[length-1].Val)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(rightSideView(tree))
}
