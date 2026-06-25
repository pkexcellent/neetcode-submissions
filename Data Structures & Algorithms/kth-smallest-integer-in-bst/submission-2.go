/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    //ar := midOrdTraversal(root)
	//return ar[k-1]
	var rs, cnt = 0, 0 
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil || cnt == k {
			return
		}
		if node.Left != nil {traverse(node.Left)}
		cnt++
		if cnt == k {rs = node.Val; return}
		if node.Right != nil {traverse(node.Right)}
		return
	}
	traverse(root)
	return rs
}

/*
// write separately
// var rs = 0
// midOrderTraverse(root, &cnt)
func midOrderTraverse(node *TreeNode, cnt *int) {
	if node == nil || *cnt == 0{
		return
	}
	if node.Left != nil && *cnt > 0 {
		midOrderTraverse(node.Left, cnt)
	}
	*cnt--
	if *cnt == 0 {rs = node.Val; return}
	if node.Right != nil && *cnt > 0 {
		midOrderTraverse(node.Right, cnt)
	}
	return 
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
*/
