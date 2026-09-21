/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	var buildNode func(l, r, t, b int) *Node
    buildNode = func(l, r, t, b int) *Node {
		rs := grid[t][l]
		for i := t; i <= b; i++ {
			for j := l; j <= r; j++ {
				if rs^grid[i][j] == 1 {
					return &Node {
						Val: true,
						IsLeaf: false,
						TopLeft: buildNode(l, (r+l)/2, t, (t+b)/2),
						TopRight: buildNode((r+l)/2+1, r, t, (t+b)/2),
						BottomLeft: buildNode(l, (r+l)/2, (t+b)/2+1, b),
						BottomRight: buildNode((r+l)/2+1, r, (t+b)/2+1, b),
					}
				}
			}
		}
		return &Node {
						Val: rs == 1,
						IsLeaf: true,
				}
	}
	return buildNode(0, len(grid[0])-1, 0, len(grid)-1)
}




