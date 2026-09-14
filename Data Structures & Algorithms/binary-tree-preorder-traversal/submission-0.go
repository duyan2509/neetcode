/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    rs:=make([]int,0)
	preorder(root,&rs)
	return rs
}

func preorder(root *TreeNode, rs *[]int) {
	if root==nil {
		return
	}
	*rs=append(*rs,root.Val)
	preorder(root.Left,rs)
	preorder(root.Right,rs)
}