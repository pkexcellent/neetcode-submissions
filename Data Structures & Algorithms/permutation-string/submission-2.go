func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {return false}
    var m1 = make(map[byte]int)
    var m2 = make(map[byte]int)
    for _, char := range []byte(s1) {
        m1[char]++
    }
    for i := 0; i < len(s1); i++ {
        m2[s2[i]]++
    }
    var l = 0
    for r := len(s1); r < len(s2); r++ {
        
        if compareMap(m1, m2) {
            return true
        }
        m2[s2[r]]++
        m2[s2[l]]--
        if m2[s2[l]] == 0 {
            delete(m2, s2[l])
        }
        l++
    }
    return compareMap(m1, m2)
}

func compareMap(m1, m2 map[byte]int) bool {
    //fmt.Println(m1, m2)
    if len(m1) != len(m2) {
        return false
    }
    for k1, v1 := range m1 {
        if v2, exist := m2[k1] ; exist {
            if v2 != v1 {
                //fmt.Println("false")
                return false
            }
        } else {
            //fmt.Println("false")
            return false
        }
    }
    return true
}
