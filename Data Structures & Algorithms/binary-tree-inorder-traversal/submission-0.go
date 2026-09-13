/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	rs := make([]int,0)
	inorder(root, &rs)
	return rs
}

func inorder(root *TreeNode, rs *[]int) {
	if root==nil {
		return
	}
	inorder(root.Left, rs)
	*rs=append(*rs,root.Val)
	inorder(root.Right, rs)
}