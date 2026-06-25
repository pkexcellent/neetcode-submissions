type MinStack struct {
	elems []int
	minElems []int
}

func Constructor() MinStack {
	return MinStack{
		elems: []int{},
		minElems: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.minElems) == 0 {
		this.minElems = append(this.minElems, val)
	} else {
		minVal := min(this.minElems[len(this.minElems)-1], val)
		this.minElems = append(this.minElems, minVal)
	}
	this.elems = append(this.elems, val)
}

func (this *MinStack) Pop() {
	this.minElems = this.minElems[:len(this.minElems)-1]
	this.elems = this.elems[:len(this.elems)-1]
}

func (this *MinStack) Top() int {
	return this.elems[len(this.elems)-1]
}

func (this *MinStack) GetMin() int {
	return this.minElems[len(this.minElems)-1]
}
