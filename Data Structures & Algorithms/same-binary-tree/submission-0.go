/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q==nil && p==nil {
		return true
	} else if q==nil || p==nil {
		return false
	}
    if p.Val!=q.Val {
		return false
	}
	return isSameTree(q.Left,p.Left) && isSameTree(q.Right,p.Right)
}
