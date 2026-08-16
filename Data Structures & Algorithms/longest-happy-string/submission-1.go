type LetterFreq struct {
	val rune
	freq int
}
type MaxHeap []LetterFreq
func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Less(i, j int) bool {
	return h[i].freq > h[j].freq
}
func (h MaxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MaxHeap) Push(x any) {*h = append(*h, x.(LetterFreq))}
func (h *MaxHeap) Pop() any {
	n := len(*h)
	last := (*h)[n-1]
	*h = (*h)[:n-1]
	return last
}
func (h *MaxHeap) Peak() any {
	return (*h)[len(*h)-1]
}
func longestDiverseString(a int, b int, c int) string {
	// use a heap to put each letter's frequency
	// pick the letter with largest fre
	// but if picked twice, pick another
	if a == 0 && b == 0 && c == 0 {
		return ""
	}
	letters := [3]LetterFreq {
		LetterFreq {
			val: 'a',
			freq: a,
		},
		LetterFreq {
			val: 'b',
			freq: b,
		},
		LetterFreq {
			val: 'c',
			freq: c,
		},
	}
	h := &MaxHeap{}
	heap.Init(h)
	for _, l := range letters {
		if l.freq > 0 {
			heap.Push(h, l)
		}	
	}
	//fmt.Println(h)
	// 
	rs := []rune{}
	pick := heap.Pop(h).(LetterFreq)
	for i := 0; i < a+b+c; i++ {
		rs = append(rs, pick.val)
		//fmt.Println(h, rs)	
		if i >= 1 && rs[i] == rs[i-1] {
			if h.Len() == 0 {
				break
			}
			nextPick := heap.Pop(h).(LetterFreq)
			pick.freq--
			if pick.freq > 0 {
				heap.Push(h, pick)
			}
			pick = nextPick
		} else {
			pick.freq--
			if pick.freq > 0 {
				heap.Push(h, pick)
			}
			if h.Len() == 0 {break}
			pick = heap.Pop(h).(LetterFreq)
		}
	}
	return string(rs)

}
