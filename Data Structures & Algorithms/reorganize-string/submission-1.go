// todo, change this pq to a interface
type freq struct {
	val rune
	f int
}
type MaxHeap []freq
func (mh MaxHeap) Len() int {return len(mh)}
func (mh MaxHeap) Less(i, j int) bool {return mh[i].f > mh[j].f}
func (mh MaxHeap) Swap(i, j int) {mh[i], mh[j] = mh[j], mh[i]}
func (mh *MaxHeap) Push(x any) {
	*mh = append(*mh, x.(freq))
}
func (mh *MaxHeap) Pop() any {
	n := len(*mh)
	last := (*mh)[n-1]
	*mh = (*mh)[:n-1]
	return last
}

func reorganizeString(s string) string {
	h := &MaxHeap{}
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
