/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func recoverTree(root *TreeNode) {
	var lrr []*TreeNode 
	dfs(root, &lrr)
    var first,second *TreeNode

    for i:=1;i<len(lrr);i++ {
        if lrr[i].Val<lrr[i-1].Val {
            if first==nil {
                first=lrr[i-1]
            }
            second=lrr[i]
        }
    }
    if first!=nil&&second!=nil {
        first.Val,second.Val=second.Val,first.Val
    }
}

func dfs(root *TreeNode, lrr *[]*TreeNode) {
	if root==nil {
		return 
	}
	dfs(root.Left, lrr)
	*lrr=append(*lrr,root)
	dfs(root.Right, lrr)
}

