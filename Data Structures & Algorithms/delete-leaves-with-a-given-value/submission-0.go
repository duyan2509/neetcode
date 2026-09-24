/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    if root==nil  {
		return nil
	}
	delete(root, root, target)
	if root.Left==nil && root.Right==nil && root.Val==target {
		return nil
	}
	return root
}

func delete(current *TreeNode, parent *TreeNode, target int) {
	if current.Left==nil && current.Right==nil {
		if target==current.Val {
			if current==parent {
				parent=nil
			} else if current==parent.Left{
				parent.Left=nil
			} else if current==parent.Right{
				parent.Right=nil
			}
		}
	} else {
		if current.Left!=nil {
			delete(current.Left, current, target)
		}
		if current.Right!=nil {
			delete(current.Right, current, target)
		}
		if current.Right==nil && current.Left==nil && current.Val==target {
			if current==parent {
				parent=nil
			} else if current==parent.Left{
				parent.Left=nil
			} else if current==parent.Right {
				parent.Right=nil
			}
		}
	}
}