func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    var letterMap = make(map[rune]int)
    for _, char := range s {
        letterMap[char] += 1 
    }
    for _, char := range t {
        if _, exist := letterMap[char]; exist {
            letterMap[char] -= 1
            if letterMap[char] == 0 {
                delete(letterMap, char)
            }
        } else {
            return false
        }
    }
    if len(letterMap) > 0 {
        return false
    }
    return true
}
