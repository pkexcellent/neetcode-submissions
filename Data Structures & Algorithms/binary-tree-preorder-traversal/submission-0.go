/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    // recursive
	preorderRs := []int{}

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		preorderRs = append(preorderRs, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return preorderRs
}
