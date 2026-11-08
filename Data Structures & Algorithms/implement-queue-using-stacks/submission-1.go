type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue {
		stack1: []int{},
		stack2: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stack2 = append(this.stack2, x)
	
	this.stack2 = append(this.stack2, this.stack1...)
		
	this.stack1 = []int{}
	this.stack1, this.stack2 = this.stack2, this.stack1
}

func (this *MyQueue) Pop() int {
	rs := this.stack1[len(this.stack1) - 1]
	this.stack1 = this.stack1[:len(this.stack1)-1]
	return rs
}

func (this *MyQueue) Peek() int {
	return this.stack1[len(this.stack1)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
