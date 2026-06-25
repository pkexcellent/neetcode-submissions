/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	var ans = [][]int{}
	if root == nil {return ans}
    var queue = []*TreeNode{root}
	
	for len(queue) != 0 {
		// by level a loop
		l := len(queue)
		var row = []int{}
		for i:= 0; i< l; i++ {
			node := queue[0]
			row = append(row, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, row)
	}
	return ans
}
