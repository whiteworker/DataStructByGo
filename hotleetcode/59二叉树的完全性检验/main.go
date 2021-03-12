func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	lastIsNil := false
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil { // 层次遍历出现nil（完全二叉树只能是最后一个节点后面全是nil才行）
			lastIsNil = true
		} else {
			if lastIsNil { // 上次为nil,本次为节点则说明不是完全二叉树
				return false
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return true
}
