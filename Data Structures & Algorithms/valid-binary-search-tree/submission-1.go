/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    // use mid-order trverse
	rs := midOrderTvs(root, []int{})
	// fmt.Println(rs)
	for i := 1; i<len(rs); i++ {
		if rs[i] <= rs[i-1] {
			return false
		}
	}
	return true
}

func midOrderTvs(root *TreeNode, rs []int) []int{
	if root == nil {
		return rs
	}
	if root.Left != nil {
		rs = midOrderTvs(root.Left, rs)
	}
	rs = append(rs, root.Val)
	if root.Right != nil {	
		rs = midOrderTvs(root.Right, rs)
	}
	return rs
}
