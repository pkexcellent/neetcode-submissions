// write again, same as solution1
// use double linked list, a map
type ListNode struct {
	val int
	key int
	next *ListNode
	prev *ListNode
}
type DoubleLinkedList struct {
	leftOldest *ListNode
	rightNewest *ListNode
}

func NewDoubleLinkedList() *DoubleLinkedList {
	dll := &DoubleLinkedList {
		leftOldest: &ListNode{},
		rightNewest: &ListNode{},
	}
	dll.leftOldest.next = dll.rightNewest
	dll.rightNewest.prev = dll.leftOldest
	return dll
}

func (dl *DoubleLinkedList) insert(node *ListNode) {
	curRightPrev := dl.rightNewest.prev
	curRightPrev.next = node
	node.prev = curRightPrev
	node.next = dl.rightNewest
	dl.rightNewest.prev = node
}

func (dl *DoubleLinkedList) remove(node *ListNode) {
	next, prev := node.next, node.prev
	prev.next = next
	next.prev = prev
	node.next, node.prev = nil, nil
}

type LRUCache struct {
	cap int
	elements map[int]*ListNode // store the elements
	eleList *DoubleLinkedList
}

func Constructor(capacity int) LRUCache {
    return LRUCache {
		cap: capacity,
		elements: make(map[int]*ListNode, capacity),
		eleList: NewDoubleLinkedList(),
	}
}

func (this *LRUCache) Get(key int) int {
    if v, exist := this.elements[key]; exist {
		this.eleList.remove(v)
		this.eleList.insert(v)
		return v.val
	} 
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}
    if v, exist := this.elements[key]; exist {
		this.eleList.remove(v)
		v.val = value
		this.eleList.insert(v)
	} else {
		newNode := &ListNode {val: value, key: key}
		this.eleList.insert(newNode)
		this.elements[key] = newNode
		if len(this.elements) > this.cap {
			oldest := this.eleList.leftOldest.next
			this.eleList.remove(oldest)
			delete(this.elements, oldest.key)
		}
	}
}
