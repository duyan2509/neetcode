/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    rs:=0
	traversal(root, -101, &rs)
	return rs
}

func traversal(root *TreeNode, oldMax int, rs *int)  {
	if root==nil {
		return 
	}
	if root.Val>=oldMax {
		oldMax=root.Val
		*rs++
	}
	traversal(root.Left, oldMax, rs)
	traversal(root.Right, oldMax, rs)
}
