/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
		return &TreeNode {Val: val}
	}
	cur := root
	for {
		if val > cur.Val {
			if cur.Right == nil {
				cur.Right = &TreeNode {Val: val}
				return root
			} else {
				cur = cur.Right
			}
		} else {
			if cur.Left == nil {
				cur.Left = &TreeNode {Val: val}
				return root
			} else {
				cur = cur.Left
			}
		}
	}
	return root
}
