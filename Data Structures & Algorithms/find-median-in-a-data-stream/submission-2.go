type mheap[T any] struct {
    data []T
    less func(i, j T) bool
}
func (h mheap[T]) Len() int {return len(h.data)}
func (h mheap[T]) Less(i, j int) bool {return h.less(h.data[i], h.data[j])}
func (h mheap[T]) Swap(i, j int) {h.data[i], h.data[j] = h.data[j], h.data[i]}
func (h *mheap[T]) Push(x any) {h.data = append(h.data, x.(T))}
func (h *mheap[T]) Pop() any {
	size := len(h.data)
	top := h.data[size-1]
	h.data = h.data[0:size-1]
	return top
}

type MedianFinder struct {
    small *mheap[int] // maxHeap
    large *mheap[int] // minHeap
}

func Constructor() MedianFinder {
    h1 := &mheap[int]{less: func(i, j int) bool {return i > j}}
    heap.Init(h1)
    h2 := &mheap[int]{less: func(i, j int) bool {return i < j}}
    heap.Init(h2)
    return MedianFinder{h1, h2}
}

func (this *MedianFinder) AddNum(num int) {
    // judge whether the num is larger than the max group, or the other
    if (*this.small).Len() > 0  && num < this.small.data[0] {
        heap.Push(this.small, num)
    } else if (*this.large).Len() > 0  && num > this.large.data[0] {
        heap.Push(this.large, num)
    } else {
        // if this num is in between
        heap.Push(this.large, num)
    }

    if (*this.small).Len() > (*this.large).Len() + 1 {
        x := heap.Pop(this.small)
        heap.Push(this.large, x)
    }
    if (*this.large).Len() > (*this.small).Len() + 1 {
        x := heap.Pop(this.large)
        heap.Push(this.small, x)
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if (*this.small).Len() > (*this.large).Len() {
        return float64(this.small.data[0])
    } else if (*this.small).Len() < (*this.large).Len() {
        return float64(this.large.data[0]) 
    } else {
        return float64(this.small.data[0] + this.large.data[0]) / 2.0
    }

}

