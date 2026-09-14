/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    rs:=make([]int,0)
	postorder(root,&rs)
	return rs
}

func postorder(root *TreeNode, rs *[]int) {
	if root==nil {
		return
	}
	postorder(root.Left,rs)
	postorder(root.Right,rs)
	*rs=append(*rs,root.Val)
}
