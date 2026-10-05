/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isCompleteTree(root *TreeNode) bool {
	q:=[]*TreeNode{}
	q=append(q,root)
	hasNil:=false
	for len(q)!=0 {
		children:=[]*TreeNode{}
		for i:=0;i<len(q);i++{
			current:=q[i]
			if hasNil && (current.Left!=nil ||current.Right!=nil) {
				return false
			} 
			if current.Left==nil && current.Right!=nil{
				return false
			} 
			
			if  current.Left==nil || current.Right==nil {
				hasNil=true
			}
			if current.Left!=nil {
				children=append(children, current.Left)
			}
			if current.Right!=nil {
				children=append(children, current.Right)
			}
		}
		q=children
	}
	return true
}
