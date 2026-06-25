/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	if root == nil {return 0}
    return isGoodNode(root, root.Val)
}

func isGoodNode(node *TreeNode, maxNodeSoFar int) int {
	cnt := 0
	if node.Val < maxNodeSoFar {
		cnt = 0
	} else {
		cnt = 1
		maxNodeSoFar = node.Val
	}
	if node.Left != nil {
		cnt += isGoodNode(node.Left, maxNodeSoFar)
	}
	if node.Right != nil {
		cnt += isGoodNode(node.Right, maxNodeSoFar)
	}
	return cnt
}
