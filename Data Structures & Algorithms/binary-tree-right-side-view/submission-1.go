/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root==nil {
		return nil
	}
    q:=[]*TreeNode{root}
	rs:=[]int{}
	for len(q)!=0{
		size:=len(q)
		rs=append(rs,q[size-1].Val)
		for i:=0;i<size;i++{
			current:=q[i]
			if current.Left!=nil {
				q=append(q,current.Left)
			}
			if current.Right!=nil {
				q=append(q,current.Right)
			}
		}
		q=q[size:len(q)]
	}
	return rs
}
