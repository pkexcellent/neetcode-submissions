// todo, change this pq to a interface
type freq struct {
	val rune
	f int
}
type MaxHeap[T any] struct {
	data []T
	less func(x, y T) bool
}

func NewMaxHeap[T any](less func(x, y T) bool) *MaxHeap[T] {
	return &MaxHeap[T]{
		less: less,
	}
}
func (mh MaxHeap[T]) Len() int {return len(mh.data)}
func (mh MaxHeap[T]) Less(i, j int) bool {return mh.less(mh.data[i], mh.data[j])}
func (mh MaxHeap[T]) Swap(i, j int) {mh.data[i], mh.data[j] = mh.data[j], mh.data[i]}
func (mh *MaxHeap[T]) Push(x any) {
	mh.data = append(mh.data, x.(T))
}
func (mh *MaxHeap[T]) Pop() any {
	n := len(mh.data)
	last := mh.data[n-1]
	mh.data = mh.data[:n-1]
	return last
}

func reorganizeString(s string) string {
	h := NewMaxHeap[freq](func(x, y freq) bool {return x.f > y.f})
	heap.Init(h)

	m := make(map[rune]int)
	maxl := 0
	for _, c := range s {
		m[c]++
		maxl = max(maxl, m[c])
	}
	if maxl > (len(s)+1)/2 {return ""}
	for k, v := range m {
		heap.Push(h, freq{val: k, f: v})
	}
	//fmt.Println(h)
	// now pop
	ans := ""
	curMax := heap.Pop(h).(freq)
	for h.Len() > 0 {
		//fmt.Println("now", curMax)
		ans += string(curMax.val)
		nextMax := heap.Pop(h).(freq)
		curMax.f--
		if curMax.f > 0 {
			heap.Push(h, curMax)
		}
		curMax = nextMax
	}
	ans += string(curMax.val)
	return ans
}
