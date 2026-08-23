/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    cur := head
	if cur == nil {
		return nil
	}
	next := cur.Next
	for next != nil {
		insertNode := &ListNode{
			Val: mostCommonDivisor(cur.Val, next.Val),
		}
		cur.Next = insertNode
		insertNode.Next = next
		cur = next
		next = next.Next
	}
	return head
}

func mostCommonDivisor(a, b int) int {
    for b > 0 {
        a, b = b, a%b
    }
    return a
}
