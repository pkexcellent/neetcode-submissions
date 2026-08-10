/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var preL *ListNode
    cur := &ListNode{Next: head}
	i := 1
	for i < left {
		cur = cur.Next
		i++
	}
	preL = cur
	for i <= right {
		cur = cur.Next
		i++
	}
	tail := cur.Next
	//fmt.Println(tail.Val, preL, preL.Next)
	// revert from preL.Next to tail.Prev
	pre := tail
	cur = preL.Next
	for cur != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	preL.Next = pre
	if left != 1 {
		return head
	} else {
		return pre
	}
}
