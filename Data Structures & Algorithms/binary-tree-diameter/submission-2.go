func diameterOfBinaryTree(root *TreeNode) int {
	rs := 0
	maxDepth(root, &rs)
	return rs
}

func maxDepth(root *TreeNode, rs *int) int {
	if root == nil {
		return 0
	}
	l:=maxDepth(root.Left, rs)
	r:=maxDepth(root.Right, rs)
	if l+r>*rs {
		*rs=l+r
	}
	if l>r {
		return l+1
	}
	return r+1
}

