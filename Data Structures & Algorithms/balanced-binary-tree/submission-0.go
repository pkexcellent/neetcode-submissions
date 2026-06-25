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
	
	left := height(root.Left)
	right := height(root.Right)
	if left - right <= 1 && right - left <= 1 {
		if isBalanced(root.Left) && isBalanced(root.Right) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}

}
func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(height(node.Left), height(node.Right))
}
