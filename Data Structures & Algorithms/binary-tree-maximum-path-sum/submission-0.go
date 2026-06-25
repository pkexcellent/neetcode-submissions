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
	dfs(root, &maxSoFar)
	return maxSoFar
}

func dfs(node *TreeNode, maxSoFar *int) {
	if node == nil {
		return
	}
	includePath := getDownMaxPath(node.Left) + getDownMaxPath(node.Right) + node.Val
	*maxSoFar = max(*maxSoFar, includePath)
	dfs(node.Left, maxSoFar)
	dfs(node.Right, maxSoFar)
}

func getDownMaxPath(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(0, node.Val + max(getDownMaxPath(node.Left), getDownMaxPath(node.Right)))
}
