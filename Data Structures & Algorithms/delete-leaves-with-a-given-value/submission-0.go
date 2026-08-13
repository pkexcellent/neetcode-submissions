/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    // use recursive
	
	var dfs func(node *TreeNode) *TreeNode 
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Left != nil {
			node.Left = dfs(node.Left)
		}
		if node.Right != nil {
			node.Right = dfs(node.Right)
		}
		if node.Left == nil && node.Right == nil {
			// a leaf
			if node.Val == target {
				return nil
			} else {
				return node
			}
		}
		
		return node
	}
	return dfs(root)
}
