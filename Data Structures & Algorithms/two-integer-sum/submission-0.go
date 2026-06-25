func twoSum(nums []int, target int) []int {
    var peerMap = make(map[int]int) // key is the peer number, value is the idx
    for idx, num := range nums {
        peerNum := target - num
        if peerIndex, exist := peerMap[peerNum]; exist {
            if peerIndex < idx {
                return []int{peerIndex, idx}
            }
            return []int{idx, peerIndex}
        } else {
            peerMap[num] = idx
        }
    }
    return nil
}
