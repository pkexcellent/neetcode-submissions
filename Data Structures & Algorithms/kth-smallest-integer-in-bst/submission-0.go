/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    ar := midOrdTraversal(root)
	return ar[k-1]
}

// magic
func midOrdTraversal(root *TreeNode) []int {
	var rs = []int{}
	if root == nil {return rs}
	if root.Left != nil {
		rs = append(rs, midOrdTraversal(root.Left)...)
	}
	rs = append(rs, root.Val)
	if root.Right != nil {
		rs = append(rs, midOrdTraversal(root.Right)...)
	}
	return rs
}
