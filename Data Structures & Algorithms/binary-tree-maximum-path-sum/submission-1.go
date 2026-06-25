/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    var maxSoFar = math.MinInt64
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftPathDownward := max(0, dfs(node.Left))
		rightPthDownward := max(0, dfs(node.Right))
		// use a global val to calculate the left+right+node
		maxSoFar = max(maxSoFar, node.Val + leftPathDownward + rightPthDownward)
		// return the current node + the max path downward
		return node.Val + max(leftPathDownward, rightPthDownward) 
	}
	dfs(root)
	return maxSoFar
}
