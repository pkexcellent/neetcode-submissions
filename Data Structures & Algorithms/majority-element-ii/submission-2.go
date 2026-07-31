func majorityElement(nums []int) []int {
	cand1, cand2 := -1000, -1000
	count1, count2 := 0, 0 
	for i := 0; i < len(nums); i++ {
		if cand1 == nums[i] {
			count1++
		} else if cand2 == nums[i] {
			count2++
		} else if count1 <= 0 {
			cand1 = nums[i]
			count1 = 1
		} else if count2 <= 0 {
			cand2 = nums[i]
			count2 = 1
		} else {
			count1--
			count2--
		}
	}
	// still need to verify if they are > n/3
	cnt1, cnt2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == cand1 {
			cnt1++
		} else if nums[i] == cand2 {
			cnt2++
		}
	}
	//fmt.Println(cand1, cand2)
	rs := []int{}
	if cnt1 > len(nums)/3 {
		rs = append(rs, cand1)
	}
	if cnt2 > len(nums)/3 {
		rs = append(rs, cand2)
	}
	return rs
}
