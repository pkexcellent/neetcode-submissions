type doubleListNode struct {
	val int
	key int
	pre *doubleListNode
	next *doubleListNode
}

type LRUCache struct {
	cap int
	data map[int]*doubleListNode
	newest *doubleListNode
	oldest *doubleListNode
}

func Constructor(capacity int) LRUCache {
    var cache = LRUCache{
		cap: capacity,
		data: make(map[int]*doubleListNode),
		newest: &doubleListNode{},
		oldest: &doubleListNode{},
	}
	cache.newest.pre = cache.oldest
	cache.oldest.next = cache.newest
	return cache
}

func (this *LRUCache) _remove(node *doubleListNode) {
	pre := node.pre
	next := node.next
	pre.next = next
	next.pre = pre
}
func (this *LRUCache) _insert(node *doubleListNode) {
	pre := this.newest.pre
	pre.next = node
	node.pre = pre
	node.next = this.newest
	this.newest.pre = node
}

func (this *LRUCache) Get(key int) int {
    if node, exist := this.data[key]; exist {
		this._remove(node)
		this._insert(node)
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
    if node, exist := this.data[key]; exist {
		this._remove(node)
	} 

	newNode := &doubleListNode{val: value, key: key}
	this.data[key] = newNode
	this._insert(newNode)

	if len(this.data) > this.cap {
		oldestNode := this.oldest.next
		this._remove(oldestNode)
		delete(this.data, oldestNode.key)
	}
}
