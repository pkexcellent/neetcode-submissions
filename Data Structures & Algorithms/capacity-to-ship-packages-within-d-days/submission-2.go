func shipWithinDays(weights []int, days int) int {
	sum, maxw := 0, 0
	for _, w := range weights {
		sum += w
		maxw = max(w, maxw)
	}
	l, r, mcap := maxw, sum, -1
	bigcan := r
	for l <= r {
		mcap = l + (r-l)/2
		day := 1
		tmpsum := 0
		for _, w := range weights {
			tmpsum += w
			if tmpsum > mcap {
				tmpsum = w
				day++
			}
		}
		fmt.Println(l, r, mcap, day)
		if day <= days {
			bigcan = mcap
			r = mcap - 1
		} else { // day > days {
			l = mcap + 1
		}
	}
	return bigcan
}
