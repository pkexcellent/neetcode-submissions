/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
   var cur *ListNode = &ListNode{0, nil}
   var head = cur
   for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
		//fmt.Println(list1.Val, list2.Val, cur.Val)
   } 
   if list1 != nil {
		cur.Next = list1
		
   	}
	if list2 != nil {
		cur.Next = list2
   	}
	return head.Next
}
