/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	rs := []int{}
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node != nil {
			traversal(node.Left)
			rs = append(rs, node.Val)
			traversal(node.Right)
		}
	}
	traversal(root)
	return rs
}
