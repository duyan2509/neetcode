/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder)==0 || len(postorder)==0 {
		return nil
	}
	rootVal:=postorder[len(postorder)-1]
	root:=&TreeNode{Val:rootVal}
	inorderRootIndex:=0
	for i:=0;i<len(inorder);i++{
		if inorder[i]==rootVal {
			inorderRootIndex=i
			break
		}
	}
	root.Left=buildTree(inorder[:inorderRootIndex],postorder[:inorderRootIndex])
	root.Right=buildTree(inorder[inorderRootIndex+1:len(inorder)],postorder[inorderRootIndex:len(inorder)-1])
	return root
}

