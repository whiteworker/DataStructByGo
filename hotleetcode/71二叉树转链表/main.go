func flatten(root *TreeNode)  {
    list := []*TreeNode{}
    stack := []*TreeNode{}
    node := root
    for node != nil || len(stack) > 0 {
        for node != nil {
            list = append(list, node)
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1]
        node = node.Right
        stack = stack[:len(stack)-1]
    }

    for i := 1; i < len(list); i++ {
        prev, curr := list[i-1], list[i]
        prev.Left, prev.Right = nil, curr
    }
}

