/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root==nil {
		return [][]int{}
	}
	rs:=make([][]int,0)
	q:=[]*TreeNode{root}
	flip:=true
	for len(q)!=0{
		tmp:=make([]*TreeNode,0)
		var currentRow []int
		if !flip{
			for i:=len(q)-1;i>=0;i--{
				currentRow=append(currentRow, q[i].Val)
				if q[i].Right!=nil {
					tmp=append(tmp,q[i].Right)
				}
				if q[i].Left!=nil{
					tmp=append(tmp,q[i].Left)
				}
			}
		} else {
			for i:=len(q)-1;i>=0;i--{
				currentRow=append(currentRow, q[i].Val)
				if q[i].Left!=nil{
					tmp=append(tmp,q[i].Left)
				}
				if q[i].Right!=nil {
					tmp=append(tmp,q[i].Right)
				}
			}
		}
		rs=append(rs,currentRow)
		q=tmp
		flip=!flip
	}
	return rs
}
