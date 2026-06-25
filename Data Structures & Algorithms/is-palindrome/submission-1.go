func isPalindrome(s string) bool {
    var r = []rune(s)
    for i, j := 0, len(s)-1; i<=j; i, j = i+1, j-1 {
        for i < len(s) {
            if !( (r[i] >= 'A' && r[i] <= 'Z') || (r[i] >= 'a' && r[i] <= 'z') || (r[i] >= '0' && r[i] <= '9')) {
                i++
            } else {
                break
            }
        }
        for j >= 0 {
            if !( (r[j] >= 'A' && r[j] <= 'Z') || (r[j] >= 'a' && r[j] <= 'z') || (r[j] >= '0' && r[j] <= '9')) {
                j--
            } else {
                break
            }
        }
        //fmt.Println(i, j)
        if i > j {
            return true
        }
        if unicode.ToLower(r[i]) != unicode.ToLower(r[j]) {
            return false
        }
    }
    return true
}
