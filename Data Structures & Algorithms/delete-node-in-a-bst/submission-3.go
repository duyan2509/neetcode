/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
    current:=root
	parent:=root
	for current!=nil && current.Val!=key{
		parent=current
		if current.Val>key {
			current=current.Left
		} else if current.Val<key {
			current=current.Right
		}
	}
	if current==root && root.Left==nil && root.Right==nil {
		return nil
	}
	if current==nil {
		return root
	}
	if current.Left!=nil {
		current.Val=getAndDeleteMax(current, current.Left, current)
	} else if current.Right!=nil {
		current.Val=getAndDeleteMin(current, current.Right, current)
	} else {
		if parent.Left==current {
			parent.Left=nil
		} else if parent.Right==current {
			parent.Right=nil
		}
	}
	return root
}

func getAndDeleteMax(deleted *TreeNode, root *TreeNode, parent *TreeNode) int {
	if root.Right==nil && root.Left==nil {
		if deleted==parent{
			parent.Left=nil
		} else {
			parent.Right=nil
		}
		return root.Val
	}
	if root.Right==nil {
		if deleted==parent {
			parent.Left=root.Left
		} else {
			parent.Right=nil
		}
		return root.Val
	}
	return getAndDeleteMax(deleted, root.Right, root)
}

func getAndDeleteMin(deleted *TreeNode, root *TreeNode, parent *TreeNode) int {
	if root.Right==nil && root.Left==nil {
		if deleted==parent{
			parent.Right=nil
		} else {
			parent.Left=nil
		}
		return root.Val
	}
	if root.Left==nil {
		if deleted==parent {
			parent.Right=root.Right
		} else {
			parent.Left=nil
		}
		return root.Val
	}
	return getAndDeleteMin(deleted, root.Right, root)
}
