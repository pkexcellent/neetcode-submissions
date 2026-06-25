func lengthOfLongestSubstring(s string) int {
    // a map, key is the char, value is the most right index
    var runeM = make(map[rune]int, len(s))
	var l, r = 0, 0
	var maxL = 0
	for idx, char := range []rune(s) {
		if mostRightIdx, exist := runeM[char]; exist {
			if mostRightIdx >= l {
				l = mostRightIdx + 1
			}
			runeM[char] = idx
            r = idx
			maxL = max(maxL, r-l+1)
		} else {
			runeM[char] = idx
			r = idx
			maxL = max(maxL, r-l+1)
		}
	}
    return maxL
}
/*
"thequickbro wnfoxjumps overthelaz  ydogthequi ckbrownfoxjumpsovert"
 0         10         20         30          40
*/
