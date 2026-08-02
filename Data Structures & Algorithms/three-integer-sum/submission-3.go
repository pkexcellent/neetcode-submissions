func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    m := make(map[int][]int)
    for i, num := range nums {
        m[num] = append(m[num], i)
    }
    rs := [][]int{}
    for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
        for j := i+1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
            sum := nums[i] + nums[j]
            if indices, exist := m[-1*sum]; exist {
				for _, idx := range indices {
					if idx > j {
						rs = append(rs, []int{nums[i], nums[j], nums[idx]})
						break
					}
				}
            }
        }
    }
    return rs
}
