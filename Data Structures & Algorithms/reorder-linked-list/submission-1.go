/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	// this needs to reverse the last half list
	// use fast and slow pointers to find the mid
	// to ensure fast's length is twice the length of slow
	// fast starts from head.next
	var slow, fast = head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var second = slow.Next
	slow.Next = nil
	var pre *ListNode
	for second != nil {
		next := second.Next
		second.Next = pre
		pre = second
		second = next
	}
	second = pre
	for pre != nil {fmt.Println(pre.Val); pre = pre.Next}
	// merge the two list
	var first = head
	var firstNext, secondNext = first, second
	for second != nil {
		firstNext = first.Next
		first.Next = second

		secondNext = second.Next
		second.Next = firstNext

		first = firstNext
		second = secondNext
	}
}
