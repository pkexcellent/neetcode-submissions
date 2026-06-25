/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	// use a hashmap to store the New Node's reference
	var oldNewMap = make(map[*Node]*Node)
	var cur = head
	for cur != nil {
		copied := &Node{Val: cur.Val}
		oldNewMap[cur] = copied
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		copied := oldNewMap[cur]
		copied.Next = oldNewMap[cur.Next]
		copied.Random = oldNewMap[cur.Random]
		cur = cur.Next
	}
	return oldNewMap[head]
}
