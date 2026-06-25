func largestRectangleArea(heights []int) int {
  // use a smarter way
  // use monotonic stack, in stack, store the index
  var stack = []int{}
  heights = append(heights, 0)
  var largest = 0
  for i:= 0; i < len(heights); i++ {
	if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
		stack = append(stack, i)
		//fmt.Println(i, stack)
	} else {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := 0
			if len(stack) == 0 {
				w = i
			} else {
				w = i - (stack[len(stack)-1]+1)// mind to use the stack top to get the left bound
			}
			largest = max(largest, h*w)
			
			//fmt.Println(i, stack, h, w)
		}
		stack = append(stack, i)
	}
  }
  return largest
  
  /*
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
  */

}
