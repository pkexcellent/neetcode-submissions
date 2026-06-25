/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {return nil}
    visited := make(map[*Node]*Node)
	stack := []*Node{node}
	visited[node] = &Node{Val: node.Val}
	for len(stack) > 0 {
		cur := stack[0]
		copied := visited[cur]

		for _, neighbor := range cur.Neighbors {
			var copiedChild *Node
			if n, exist := visited[neighbor]; exist {
				copiedChild = n
			} else {
				copiedChild = &Node{Val: neighbor.Val}
				visited[neighbor] = copiedChild
				stack = append(stack, neighbor)
			}
			copied.Neighbors = append(copied.Neighbors, copiedChild)
		}
		stack = stack[1:]
	}
	return visited[node]
}
