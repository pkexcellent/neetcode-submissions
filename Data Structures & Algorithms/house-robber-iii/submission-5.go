/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func rob(root *TreeNode) int {
    // use recursive method
	// add memo 
	memo := make(map[*TreeNode]int)
	var robb func(node *TreeNode) int
	robb = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if v, exist := memo[node]; exist {
			return v
		}
		// rob this one
		cur := node.Val
		sumRob := cur
		if node.Left != nil {
			sumRob += robb(node.Left.Left) + robb(node.Left.Right)
		}
		if node.Right != nil {
			sumRob += robb(node.Right.Left) + robb(node.Right.Right)
		}
		// not rob this one
		sumNotRob := robb(node.Left) + robb(node.Right)
		rs := max(sumNotRob, sumRob)
		memo[node] = rs
		return rs
	}
	return robb(root)
}
