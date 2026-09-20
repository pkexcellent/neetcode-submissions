type MyCircularQueue struct {
    q []int
	head int
	rear int
	capacity int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue {
		q: make([]int, k),
		head: 0,
		rear: -1,
		capacity: k,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.rear - this.head + 1 >= this.capacity {
		return false
	}
	this.rear++
	this.q[this.rear%this.capacity] = value
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.rear - this.head + 1 > 0 {
		this.head++
		return true
	}
	return false
}


func (this *MyCircularQueue) Front() int {
    if this.rear - this.head + 1 > 0 {
		return this.q[this.head%this.capacity]
	} else {
		return -1
	}
}


func (this *MyCircularQueue) Rear() int {
    if this.rear - this.head + 1 > 0 {
		return this.q[this.rear%this.capacity]
	} else {
		return -1
	}
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.rear - this.head + 1 == 0
}


func (this *MyCircularQueue) IsFull() bool {
    return this.rear - this.head + 1 == this.capacity 
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param1 := obj.EnQueue(value);
 * param2 := obj.DeQueue();
 * param3 := obj.Front();
 * param4 := obj.Rear();
 * param5 := obj.IsEmpty();
 * param6 := obj.IsFull();
 */
 