type stringWithIdx struct {
    str string
    idx int
}
func groupAnagrams(strs []string) [][]string {
    var sortedStrs = make([]stringWithIdx, len(strs))
    for idx, str := range strs {
        chars := []rune(str)
        sort.Slice(chars, func(i, j int) bool {
            return chars[i] < chars[j]
        })
        sortedStrs[idx] = stringWithIdx{string(chars), idx}
    }
    // order by ordered strings
    sort.Slice(sortedStrs, func(i, j int) bool {return sortedStrs[i].str < sortedStrs[j].str})

    var ans = make([][]string, 0, len(strs))
    if len(sortedStrs) > 0 {
        current := []string{strs[sortedStrs[0].idx]}
        for i := 1; i < len(sortedStrs); i++ {
            if sortedStrs[i].str == sortedStrs[i-1].str {
                current = append(current, strs[sortedStrs[i].idx])
            } else {
                ans = append(ans, current)
                current = []string{strs[sortedStrs[i].idx]}
            }
        }
        ans = append(ans, current)
    }
    
    return ans

}
