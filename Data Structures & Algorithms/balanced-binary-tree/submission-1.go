/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	rs:=true
	maxDepth(root,&rs)
	return rs
}

func maxDepth(root *TreeNode, rs *bool) int {
	if root==nil {
		return 0
	}

	l:=0
	r:=0
	if root.Left!=nil {
		l=1+maxDepth(root.Left, rs)
	}
	if root.Right!=nil {
		r=1+maxDepth(root.Right, rs)
	}
	if l-r>1 || r-l>1 {
		*rs=false
	}
	if l>r {
		return l
	}
	return r
}
