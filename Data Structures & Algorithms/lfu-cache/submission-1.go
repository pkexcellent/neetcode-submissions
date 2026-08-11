type ListNode struct {
	key int
	val int
	freq int
	prev *ListNode
	next *ListNode
}
type DoubleLinkedList struct {
	size int
	left *ListNode
	right *ListNode
}
func (dll *DoubleLinkedList) PushRight(node *ListNode) {
	rprev := dll.right.prev
	rprev.next = node
	node.prev = rprev
	node.next = dll.right
	dll.right.prev = node
	dll.size++
}
func (dll *DoubleLinkedList) Pop(node *ListNode) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
	node.prev, node.next = nil, nil
	dll.size--
}
func (dll *DoubleLinkedList) PopLeft() *ListNode {
	nodeToPop := dll.left.next
	dll.left.next = nodeToPop.next
	nodeToPop.next.prev = dll.left
	nodeToPop.next, nodeToPop.prev = nil, nil
	dll.size--
	return nodeToPop
}

type LFUCache struct {
    minFeq int
	cap int
	data map[int]*ListNode
	llWithFreq map[int]*DoubleLinkedList
}

func Constructor(capacity int) LFUCache {
	lfu := LFUCache {
		minFeq: 0,
		cap: capacity,
		data: make(map[int]*ListNode, capacity),
		llWithFreq: make(map[int]*DoubleLinkedList),
	}
	return lfu
}


func (this *LFUCache) Get(key int) int {
    if v, exist := this.data[key]; exist {
		// adjust freq
		feq := v.freq
		this.llWithFreq[feq].Pop(v)
		if this.llWithFreq[feq].size == 0 {
			if feq == this.minFeq {
				this.minFeq++
			}
		}
		if _, exist := this.llWithFreq[feq+1]; !exist {
			this.llWithFreq[feq+1] = &DoubleLinkedList {
				left: &ListNode{},
				right: &ListNode{},
			}
			this.llWithFreq[feq+1].left.next = this.llWithFreq[feq+1].right
			this.llWithFreq[feq+1].right.prev = this.llWithFreq[feq+1].left
		}
		v.freq++
		this.llWithFreq[feq+1].PushRight(v)
		return v.val
	}
	return -1
}


func (this *LFUCache) Put(key int, value int)  {
	newNode := &ListNode {
		val: value,
		key: key,
		freq: 1,
	}
    if v, exist := this.data[key]; exist {
		newNode.freq = v.freq + 1
		this.llWithFreq[v.freq].Pop(v)
		if this.llWithFreq[v.freq].size == 0 {
			if v.freq == this.minFeq {
				this.minFeq++
			}
		}
	} else {
		// when insert a new one, mind the capacity thing
		if len(this.data) == this.cap {
			// the min freq, the left most one need to be pop
			toRemove := this.llWithFreq[this.minFeq].PopLeft() 
			delete(this.data, toRemove.key)
		}
		this.minFeq = 1 // we'll insert a new with f=1, no matter if cap is full or not in this case
	}
	this.data[key] = newNode
	if _, exist := this.llWithFreq[newNode.freq]; !exist {
		this.llWithFreq[newNode.freq] = &DoubleLinkedList {
			left: &ListNode{},
			right: &ListNode{},
		}
		this.llWithFreq[newNode.freq].left.next = this.llWithFreq[newNode.freq].right
		this.llWithFreq[newNode.freq].right.prev = this.llWithFreq[newNode.freq].left
	}
	this.llWithFreq[newNode.freq].PushRight(newNode)

}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param1 := obj.Get(key);
 * obj.Put(key,value);
 */
