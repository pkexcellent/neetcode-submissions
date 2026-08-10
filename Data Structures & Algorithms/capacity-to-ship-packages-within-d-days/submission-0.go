func shipWithinDays(weights []int, days int) int {
	sum, avg, maxw := 0, 0, -1
	for _, w := range weights {
		sum += w
		maxw = max(maxw, w)
	}
	n := len(weights)
	avg = (sum+n-1)/n
	mincap := max(avg, maxw)
	for {
		day := 1
		tmpsum := 0
		//fmt.Println(mincap, avg, maxw)
		for i, w := range weights {
			tmpsum += w
			//fmt.Println(day, tmpsum)
			if tmpsum > mincap {
				tmpsum = w
				day++
			} 
			if day > days {
				mincap++
				break
			} else if i == n-1 {
				return mincap
			}
		}
	}
	return mincap
}
