func minEatingSpeed(piles []int, h int) int {
	var maxTime = 0
	for _, pile := range piles {
		if pile > maxTime {
			maxTime = pile
		}
	}
	var l, r = 1, maxTime
	var minSpeed = maxTime
	for l <= r {
		mid := l + (r-l)/2
		
		totalT := 0
		for _, pile := range piles {
			totalT += int(math.Ceil(float64(pile)/float64(mid)))
		}
		fmt.Println(mid, totalT)
		if totalT <= h {
			minSpeed = mid	
			r = mid-1
		} else if totalT > h {
			l = mid+1
		}

	}
	return minSpeed
}
