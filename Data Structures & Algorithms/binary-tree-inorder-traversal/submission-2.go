/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	// use stack to traversal
	if root == nil {
		return []int{}
	}
	inorderRs := []int{}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		if len(stack) != 0 {
			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			inorderRs = append(inorderRs, top.Val)
			cur = top.Right
		}
	}
	return inorderRs
}
