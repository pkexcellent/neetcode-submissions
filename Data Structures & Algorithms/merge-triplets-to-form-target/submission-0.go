func mergeTriplets(triplets [][]int, target []int) bool {
    // use col to find which row(s) contain the target number
    // then we can find out the mustSelected(can be multiple) rows, and must exclude rows(having num > target)
    // after getting exclude rows by column, use exclude to extract mustSelect, 
    // if any set is empty, false
    row, _ := len(triplets), len(triplets[0])
    excludeRows := make(map[int]struct{})
    selectedRows := []map[int]struct{}{}
    for c := 0; c < len(target); c++ {
        selectedRow := make(map[int]struct{})
        for r := 0; r < row; r++ {
            if triplets[r][c] > target[c] {
                excludeRows[r] = struct{}{}
            }
            if triplets[r][c] == target[c] {
                selectedRow[r] = struct{}{}
            }
        }
        selectedRows = append(selectedRows, selectedRow)
    }
    // 
    for _, rows := range selectedRows {
        for execludeRow, _ := range excludeRows {
            if _, exist := rows[execludeRow]; exist {
                delete(rows, execludeRow)
            }
        }
        if len(rows) == 0 {
            return false
        }
    }
    return true
}
