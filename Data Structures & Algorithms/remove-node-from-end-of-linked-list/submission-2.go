/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// use slow fast pointer
	var slow, fast = head, head
	for n>0 && fast != nil {
		fast = fast.Next
		n--
	}
	if (n>0) {
		return nil
	}
	var pre *ListNode
	for fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}
	if pre == nil { // pre is never moved, indicate the fist is to be removed
		return slow.Next
	} else {
		pre.Next = slow.Next
	}
	return head
}
