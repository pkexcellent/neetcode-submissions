/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    var dummy = &ListNode{Next: head}
	var groupPre = dummy
	for {
		var i = 0
		var nextHead = head
		for i < k && nextHead != nil {
			nextHead = nextHead.Next
			i++
		}
		if i < k {
			break
		}
		i = 0
		var pre = nextHead
		for i < k && head != nil {
			//fmt.Println("head: ", head.Val, "pre:", pre.Val)
			next := head.Next
			head.Next = pre
			pre = head
			head = next
			i++
		}
		rear := groupPre.Next
		//fmt.Println(head.Val, pre.Val, rear.Val)
		groupPre.Next = pre
		groupPre = rear
	}
	return dummy.Next

}
