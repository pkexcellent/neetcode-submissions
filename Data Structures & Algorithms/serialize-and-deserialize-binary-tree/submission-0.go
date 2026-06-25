/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// use level traversal to serilize the tree
	var rs = []string{}
	var stack = []*TreeNode{root}
	for len(stack) > 0 {
		front := stack[0]
		stack = stack[1:]
		if front != nil {
			stack = append(stack, []*TreeNode{front.Left, front.Right}...)
			rs = append(rs, strconv.Itoa(front.Val))
		} else {
			rs = append(rs, "N")
		}
	}
	return strings.Join(rs, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	if vals[0] == "N" {return nil}
	rootVal, _ := strconv.Atoi(vals[0])
	var root = &TreeNode{Val: rootVal}
	var queue = []*TreeNode{root}
	var idx = 1
	for len(queue) > 0 && idx < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if vals[idx] != "N" {
			leftVal, _ := strconv.Atoi(vals[idx])
			node.Left = &TreeNode{Val: leftVal}
			queue = append(queue, node.Left)		
		}
		idx++ // if the current idx is N, ignore it 
		if vals[idx] != "N" && idx < len(vals) {
			rightVal, _ := strconv.Atoi(vals[idx])
			node.Right = &TreeNode{Val: rightVal}
			queue = append(queue, node.Right)
		}
		idx++
	}
	return root
}
