/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	rs:=false
	sum(root, &rs, targetSum, 0)
	return rs
}

func sum(root *TreeNode, rs *bool, targetSum int, current int)  {
	if root!=nil && root.Right==nil && root.Left==nil {
		current+=root.Val
		if current==targetSum {
			*rs=true
		}
	} else if root!=nil {
		sum(root.Left, rs, targetSum, current+root.Val)
		sum(root.Right, rs, targetSum, current+root.Val)
	}
}
