/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root==nil {
		return &TreeNode{Val:val}
	}
    current:=root
	for current!=nil {
		if val>current.Val{
			if current.Right==nil {	
				newNode:=&TreeNode{Val:val}
				current.Right=newNode
				break
			} else {
				current=current.Right
			}
		} else {
			if current.Left==nil {	
				newNode:=&TreeNode{Val:val}
				current.Left=newNode
				break
			} else {
				current=current.Left
			}
		}
	}
	return root
}
