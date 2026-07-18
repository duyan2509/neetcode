/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    current:=root
	if current!=nil {
		tmp:=current.Left
		current.Left=current.Right
		current.Right=tmp
		invertTree(current.Left)
		invertTree(current.Right)
	}
	return root
}
