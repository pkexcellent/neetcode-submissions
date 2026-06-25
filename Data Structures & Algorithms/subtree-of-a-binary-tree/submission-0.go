/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil && subRoot == nil {
		return true
	}

	if root != nil && subRoot != nil && root.Val == subRoot.Val {
		if sameTree(root.Left, subRoot.Left) && sameTree(root.Right, subRoot.Right) {
			return true
		}
	} 
	if root == nil && subRoot != nil {
		return false
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func sameTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a != nil && b != nil {
		if a.Val == b.Val {
			return sameTree(a.Left, b.Left) && sameTree(a.Right, b.Right)
		} else {
			return false
		}
	} else {
		return false
	}
}
