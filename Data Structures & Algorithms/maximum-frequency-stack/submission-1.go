type element struct {
	val int
	freq int
	index int
}

type mheap []element

func (m mheap) Len() int {return len(m)}
func (m mheap) Less(i, j int) bool {
	if m[i].freq == m[j].freq {
		return m[i].index > m[j].index
	} 
	return m[i].freq > m[j].freq
}
func (m mheap) Swap(i, j int) {m[i], m[j] = m[j], m[i]}
func (m *mheap) Push(x interface{}) {
	*m = append(*m, x.(element))
}
func (m *mheap) Pop() interface{} {
	last := (*m)[len(*m)-1]
	*m = (*m)[0:len(*m) - 1]
	return last
}

type FreqStack struct {
	stack *mheap
	freqM map[int]int
	curCnt int
}

func Constructor() FreqStack {
	mh := &mheap{}
	heap.Init(mh)
	return FreqStack {
		stack: mh, 
		freqM: make(map[int]int),
		curCnt: -1,
	}
}

func (this *FreqStack) Push(val int) {
	this.freqM[val]++
	this.curCnt++

	newEle := element {
		val: val,
		freq: this.freqM[val],
		index: this.curCnt,
	}
	heap.Push(this.stack, newEle)
}

func (this *FreqStack) Pop() int {
	ele := heap.Pop(this.stack).(element)
	this.freqM[ele.val]--
	return ele.val
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor()
 * obj.Push(val)
 * param2 := obj.Pop()
 */
 