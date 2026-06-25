/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	var cur = root
    var p2root = pathOfNode(cur, p)
	cur = root
	var q2root = pathOfNode(cur, q)
	//fmt.Println(p2root)
	//fmt.Println(q2root)
	var i = 0
	for i < len(p2root) && i < len(q2root) {
		if p2root[i].Val == q2root[i].Val {
			i++
		} else {
			return p2root[i-1]
		}
	}
	if i >= len(p2root) {
		return q2root[i-1]
	} else {
		return p2root[i-1]
	}
	
}

func pathOfNode(root, node *TreeNode) []*TreeNode {
	var path = []*TreeNode{}
	for root.Val != node.Val {
		if node.Val > root.Val {
			path = append(path, root)
			root = root.Right
		} else {
			path = append(path, root)
			root = root.Left
		}
	}
	path = append(path, root)
	return path
}
