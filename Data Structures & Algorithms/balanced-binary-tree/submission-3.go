/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}
	var isBalance = true
	var height func(*TreeNode) int

	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := height(node.Left)
		right := height(node.Right)
		if left - right > 1 || right - left > 1 {
			isBalance = false 
		}
		return 1 + max(height(node.Left), height(node.Right))
	}
	
	left := height(root.Left)
	right := height(root.Right)
	return isBalance && left-right <=1 && right-left <=1
}
