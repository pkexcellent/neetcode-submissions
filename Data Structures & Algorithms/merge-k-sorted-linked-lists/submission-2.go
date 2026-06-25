/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
	var dummy = &ListNode{}
	var cur = dummy
	
	for {
		var minListIdx = -1
		for i, list := range lists {
			if list == nil {
				continue
			}
			if minListIdx == -1 || list.Val < lists[minListIdx].Val {
				minListIdx = i
			}
		}
		if minListIdx == -1 {
			break
		}
		cur.Next = lists[minListIdx]
		cur = cur.Next
		lists[minListIdx] = lists[minListIdx].Next
	}
	return dummy.Next

}