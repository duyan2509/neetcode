/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    rs:=0
	traversal(root, &k, &rs)
	return rs
}

func traversal(root *TreeNode, k *int, rs *int)  {
	if root==nil {
		return 
	}
	traversal(root.Left, k, rs)
	*k--
	if *k==0 {
		*rs=root.Val
	}
	traversal(root.Right, k, rs)
}