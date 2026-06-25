/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// use slow fast pointer
	var dummy = &ListNode{Next:head}
	var slow, fast = dummy, head
	for n>0 && fast != nil {
		fast = fast.Next
		n--
	}
	if (n>0) {
		return nil
	}
	//var pre *ListNode
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}
