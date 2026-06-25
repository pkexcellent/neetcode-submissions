type minHeap []int
func (h minHeap) Len() int {return len(h)}
func (h minHeap) Less(i, j int) bool {return h[i] < h[j]}
func (h minHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *minHeap) Push(x any) {*h = append(*h, x.(int))}
func (h *minHeap) Pop() any {
	size := len(*h)
	top := (*h)[size-1]
	(*h) = (*h)[0:size-1]
	return top
}

type maxHeap []int
func (h maxHeap) Len() int {return len(h)}
func (h maxHeap) Less(i, j int) bool {return h[i] > h[j]}
func (h maxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *maxHeap) Push(x any) {*h = append(*h, x.(int))}
func (h *maxHeap) Pop() any {
	size := len(*h)
	top := (*h)[size-1]
	(*h) = (*h)[0:size-1]
	return top
}


type MedianFinder struct {
    small *maxHeap // maxHeap
    large *minHeap // minHeap
}

func Constructor() MedianFinder {
    h1 := &maxHeap{}
    heap.Init(h1)
    h2 := &minHeap{}
    heap.Init(h2)
    return MedianFinder{h1, h2}
}

func (this *MedianFinder) AddNum(num int) {
    // judge whether the num is larger than the max group, or the other
    if (*this.small).Len() > 0  && num < (*this.small)[0] {
        heap.Push(this.small, num)
    } else if (*this.large).Len() > 0  && num > (*this.large)[0] {
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
        return float64((*this.small)[0])
    } else if (*this.small).Len() < (*this.large).Len() {
        return float64((*this.large)[0]) 
    } else {
        return float64((*this.small)[0] + (*this.large)[0]) / 2.0
    }

}

