/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return valid(root, math.MinInt64, math.MaxInt64)
	
	// fmt.Println(rs)
	/*
	for i := 1; i<len(rs); i++ {
		if rs[i] <= rs[i-1] {
			return false
		}
	}
	*/
	return true
}
//method1
func valid(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val > min && node.Val < max {
		return valid(node.Left, min, node.Val) && valid(node.Right, node.Val, max)
	} else {
		return false
	}
}

// method2
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
