/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
    var rootNodeVal = preorder[0]
	rootNode := TreeNode{Val: rootNodeVal}
	var index = pos(rootNodeVal, inorder)
	var leftPart = inorder[0:index]
	var rightPart = inorder[index+1:]
	rootNode.Left = buildTree(preorder[1:index+1], leftPart)
	rootNode.Right = buildTree(preorder[index+1:], rightPart)
	return &rootNode
}

func pos(target int, arr []int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
