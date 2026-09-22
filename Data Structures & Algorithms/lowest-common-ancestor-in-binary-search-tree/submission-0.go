/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    minV:=p.Val
	maxV:=q.Val
	if minV>maxV {
		minV,maxV=maxV,minV
	}
	for root!=nil {
		if root.Val<minV {
			root=root.Right
		} else if root.Val>maxV {
			root=root.Left
		} else {
			return root
		}
	}
	return nil
}
