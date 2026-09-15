/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	rs:=0
	q:=make([]*TreeNode,0)
	q=append(q,root)
	for len(q)!=0 {
		current:=q[0]
		l:=maxDepth(current.Left)
		r:=maxDepth(current.Right)
		if r+l>rs {
			rs=r+l
		}
		if current.Left!=nil {
			q=append(q,current.Left)
		}
		if current.Right!=nil {
			q=append(q,current.Right)
		}
		q=q[1:len(q)]
	}
	return rs
}

func maxDepth(root *TreeNode) int {
	if root==nil {
		return 0
	}
	l:=maxDepth(root.Left)
	r:=maxDepth(root.Right)
	if l>r {
		return 1+l
	}
	return 1+r
}
