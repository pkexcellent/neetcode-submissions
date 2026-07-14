func partitionLabels(s string) []int {
    // get most left and right pos for each letter
    // extend the left and right with all letter converred
    m := make(map[rune][2]int)
    for i, char := range s {
        if pos, exist := m[char]; !exist {
            m[char] = [2]int{i, i}
        } else {
            pos[1] = i
            m[char] = pos
        }
    }
    rss := []int{}
    start := 0
    for start < len(s) {
        c := rune(s[start])
        right := m[c][1]
        for j := 0; j <= right; j++ {
            cc := rune(s[j])
            right = max(right, m[cc][1])
        }
        rss = append(rss, right - start + 1)
        start = right+1
    }
    return rss
    
}
