/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	nodeFound:=0
	var rs *TreeNode
	find(root,p,q,&rs,&nodeFound)
	return rs
}

func find(root, p, q *TreeNode,  rs **TreeNode, nodeFound *int)  {
	if root==nil {
		return 
	}
	oldFound:=*nodeFound
	if root==p||root==q {
		*nodeFound++
	}
	if *nodeFound<2 {
		find(root.Left, p, q, rs, nodeFound)
		find(root.Right, p, q, rs, nodeFound)
		if *nodeFound==2 && oldFound==0{
			*rs=root
			*nodeFound=0
		} 
	}
}