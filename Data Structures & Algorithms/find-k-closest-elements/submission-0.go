func findClosestElements(arr []int, k int, x int) []int {
	// use binary search to find the target position idx
	// then l = idx-1, r=idx
	// pick the closest nums until reach k
	if len(arr) < k {
		return arr
	}
	pos := 0
	l, r := 0, len(arr)
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] > x {
			r = mid
		} else if arr[mid] < x {
			l = mid+1
		} else {
			pos = mid
			break
		}
	}
	if pos == 0 {
		pos = l
	}
	//fmt.Println("pos:", pos)
	l, r = pos-1, pos
	for k > 0 {
		if l < 0 {
			k--
			r++
			continue
		}
		if r >= len(arr) {
			k--
			l--
			continue
		}
		leftn, rightn := arr[l], arr[r]
		if abs(leftn - x) < abs(rightn - x) || 
			(abs(leftn - x) == abs(rightn - x) && leftn < rightn) {
				k--
				// pick leftn
				l--
		} else {
			k--
			//pick rightn
			r++
		}
	}
	//fmt.Println(l, r)
	return arr[l+1:r]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
