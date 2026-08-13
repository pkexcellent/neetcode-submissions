/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
	dummy := &TreeNode{Left: root}
    parent := dummy
	cur := root
	isLeft := true
	for cur != nil && cur.Val != key {
		parent = cur
		if cur.Val > key {
			cur = cur.Left
			isLeft = true
		} else {
			cur = cur.Right
			isLeft = false
		}
	}
	if cur == nil { // not found
		return root
	}
	toDelete := cur

	// find a node to swap,
	if toDelete.Left == nil && toDelete.Right == nil { // delete leaf node
		if isLeft {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		if parent == dummy {
			return nil
		} else {
			return root
		}	
	}
	swap := toDelete
	swapParent := parent
	if toDelete.Left != nil {
		swapParent = swap
		swap = toDelete.Left
		for swap.Right != nil {
			swapParent = swap
			swap = swap.Right
		}
		if swapParent == toDelete {
			swapParent.Left = swap.Left
			swap.Left = nil
		} else {
			swapParent.Right = swap.Left // put swap's left on its parent's right
			swap.Left = nil
		}
	} else {
		swapParent = swap
		swap = toDelete.Right
		for swap.Left != nil {
			swapParent = swap
			swap = swap.Left
		}
		if swapParent == toDelete {
			swapParent.Right = swap.Right
			swap.Right = nil
		} else {
			swapParent.Left = swap.Right // put swap's right on its parent's left
			swap.Right = nil
		}
	}
	if isLeft {
		parent.Left = swap
		swap.Right = toDelete.Right
		swap.Left = toDelete.Left
	} else {
		parent.Right = swap
		swap.Right = toDelete.Right
		swap.Left = toDelete.Left
	}
	if parent == dummy {
		return parent.Left
	}
	return root
}
