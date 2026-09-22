/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root==nil {
		return [][]int{}
	}
	j:=0
	rs:=make([][]int,0)
	q:=[]*TreeNode{root}
	for len(q)!=0{
		size:=len(q)
		rs=append(rs, []int{}) 
		for i:=0;i<size;i++ {
			current:=q[0]
			rs[j]=append(rs[j],current.Val)
			if current.Left!=nil {
				q=append(q,current.Left)
			}
			if current.Right!=nil {
				q=append(q,current.Right)
			}
			q=q[1:]
		}
		size=2*size
		j++
	}
	return rs
}
