/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var ans = &ListNode{Next: nil}
	var bit = 0
	var cur = ans
	for l1 != nil && l2 != nil {
		v := (l1.Val + l2.Val + bit)%10
		bit = (l1.Val + l2.Val + bit)/10
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		v := (l1.Val + bit)%10
		bit = (l1.Val + bit)/10
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		v := (l2.Val + bit)%10
		bit = (l2.Val + bit)/10
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
		l2 = l2.Next
	}
	if bit == 1 {
		cur.Next = &ListNode{Val: 1}
	}
	return ans.Next
}
