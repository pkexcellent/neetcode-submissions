func mergeTriplets(triplets [][]int, target []int) bool {
    good := make(map[int]bool)

    for _, t := range triplets {
        if t[0] > target[0] || t[1] > target[1] || t[2] > target[2] {
            continue
        }
        for i, v := range t {
            if v == target[i] {
                good[i] = true
            }
        }
    }
    return len(good) == 3
}