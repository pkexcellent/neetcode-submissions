func letterCombinations(digits string) []string {
	digitMap := map[byte]string {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
	}
	if len(digits) == 0 {return []string{}}

	rs := []string{""}
	for _, digit := range digits {
		tmp := []string{}
		for _, char := range digitMap[byte(digit)] {
			for _, one := range rs {
				one = one + string(char)
				tmp = append(tmp, one)
			}
		}
		rs = tmp
	}
	return rs
}
