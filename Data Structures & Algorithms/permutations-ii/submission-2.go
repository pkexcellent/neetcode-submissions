func permuteUnique(nums []int) [][]int {
	rs := [][]int{[]int{}}
	for _, num := range nums {
		roundRs := [][]int{}
		for _, baseCase := range rs {
			for k := 0; k <= len(baseCase); k++ {
				tmp := append([]int{}, baseCase...)
				tmp = append(tmp[0:k], append([]int{num}, tmp[k:]...)...)
				roundRs = append(roundRs, tmp)
				if k < len(baseCase) && baseCase[k] == num {
					break
				}
			}
		}
		rs = roundRs
	}
	return rs
}
