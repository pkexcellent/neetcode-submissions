/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	// binary search
	n := mountainArr.length()
	if n == 0 {return -1}
	l, r := 1, n-2
	mid := 0
	for l <= r {
		mid = l + (r-l)/2
		midn := mountainArr.get(mid)
		midnL := mountainArr.get(mid-1)
		midnR := mountainArr.get(mid+1)
		if midn > midnL && midn > midnR {
			l = mid + 1
			r = mid - 1
		} else if midnL < midn && midn < midnR {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	peak := mid
	//fmt.Println(peak)
	l, r = 0, peak 
	for l <= r {
		mid = l + (r-l)/2
		midv := mountainArr.get(mid)
		//fmt.Println(mid, midv)
		if midv < target {
			l = mid+1
		} else if midv > target {
			r = mid-1
		} else {
			return mid
		}
	}
	l, r = peak, n-1
	for l <= r {
		mid = l + (r-l)/2
		midv := mountainArr.get(mid)
		if midv > target {
			l = mid+1
		} else if midv < target {
			r = mid-1
		} else {
			return mid
		}
	}
	return -1
}
