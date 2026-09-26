/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumNumbers(root *TreeNode) int {
	return sum(root, 0)
}
func sum(root *TreeNode, prev int) int {
	if root==nil {
		return 0
	}
	current:=prev*10+root.Val
	if root.Right==nil && root.Left==nil {
		return current
	}
	return sum(root.Left, current)+sum(root.Right, current)
}
