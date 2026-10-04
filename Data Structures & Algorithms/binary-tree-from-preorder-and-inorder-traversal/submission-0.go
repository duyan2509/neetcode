/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	n:=len(inorder)
	inorderHash:=make(map[int]int)
	for i:=0;i<n;i++ {
		inorderHash[inorder[i]] = i
	}
	var BuildTree func(
		preorderStart int,
		preorderEnd int,
		inorderStart int,
		inorderEnd int,
	) *TreeNode
	BuildTree = func(
		preorderStart int,
		preorderEnd int,
		inorderStart int,
		inorderEnd int,
	) *TreeNode {
		if preorderStart>preorderEnd || inorderStart>inorderEnd {
			return nil
		}
		root:=&TreeNode{
			Val:preorder[preorderStart],
		}
		rootInd:=inorderHash[root.Val]
		leftInorderSize:=rootInd-inorderStart
		root.Left=BuildTree(
			preorderStart+1,
			preorderStart+leftInorderSize,
			inorderStart,
			rootInd-1,
		)
		root.Right=BuildTree(
			preorderStart+leftInorderSize+1,
			preorderEnd,
			rootInd+1,
			inorderEnd,
		)
		return root
	}
	return BuildTree(0, n-1, 0, n-1)
}