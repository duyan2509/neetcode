/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	if (root.Left!=nil && root.Val<=root.Left.Val) || (root.Right!=nil && root.Val>=root.Right.Val) {
		return false
	}

	return check(root.Left, root.Val, -1000000001) &&
	check(root.Right, 1000000001, root.Val)
}

func check(root *TreeNode, maxVal int, minVal int) bool {
	if root==nil {
		return true
	}

	if root.Left!=nil && (root.Val<=root.Left.Val||root.Left.Val<=minVal) {
		return false
	}

	if root.Right!=nil && (root.Val>=root.Right.Val||root.Right.Val>=maxVal) {
		return false
	}


	return check(root.Left, root.Val, minVal) &&
	check(root.Right, maxVal, root.Val)
}
