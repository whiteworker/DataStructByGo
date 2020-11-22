package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func mirrorTree(root *TreeNode) *TreeNodebo {
    var  res *TreeNode= nil
    if(root!=nil){
        res := TreeNode{Val:root.Val}
        res.Left=mirrorTree(root.Right)
        res.Right=mirrorTree(root.Left)
    }
    return res
}
type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

func main(){

}