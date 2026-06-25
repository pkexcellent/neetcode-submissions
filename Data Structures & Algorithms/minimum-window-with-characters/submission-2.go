func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
    var mt = make(map[byte]int)
	for _, c := range []byte(t) {
		mt[c]++
	}
	var matched, target = 0, len(mt)
	var matchedPos = []int{-1,-1}
	
	var l, r, ansL = 0, 0, math.MaxInt
	var ms = make(map[byte]int)
	for r = 0; r < len(s); r++ {
		var curChar = s[r]
		ms[curChar]++
		if ms[curChar] == mt[curChar] {
			matched++
			for matched == target {
				if r-l+1 <= ansL {
					matchedPos = []int{l, r}
					ansL = r-l+1
				}
				if mt[s[l]] == ms[s[l]] {
					matched--
				}
				ms[s[l]]--
				l++
			}
		}
	}
	if matchedPos[0] == -1 {
		return ""
	}
	return string(s[matchedPos[0]: matchedPos[1]+1])
}
