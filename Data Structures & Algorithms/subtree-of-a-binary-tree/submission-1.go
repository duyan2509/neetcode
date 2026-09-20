func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }

    if sameTree(root, subRoot) {
        return true
    }

    return isSubtree(root.Left, subRoot) ||
           isSubtree(root.Right, subRoot)
}

func sameTree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil && subRoot == nil {
        return true
    }

    if root == nil || subRoot == nil {
        return false
    }

    if root.Val != subRoot.Val {
        return false
    }

    return sameTree(root.Left, subRoot.Left) &&
           sameTree(root.Right, subRoot.Right)
}