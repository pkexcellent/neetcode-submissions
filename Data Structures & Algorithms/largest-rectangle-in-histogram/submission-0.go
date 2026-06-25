func largestRectangleArea(heights []int) int {
  // I will use brute force method first
  var ans = 0
  for idx, h := range heights {
	var l, r = idx, idx
	for l >=0 && heights[l] >= h {
		l--
	}
	for r < len(heights) && heights[r] >= h {
		r++
	}
	l++
	r--
	fmt.Println(idx, l, r)
	size := h * (r-l+1)
	ans = max(ans, size)
  }
  return ans

}
