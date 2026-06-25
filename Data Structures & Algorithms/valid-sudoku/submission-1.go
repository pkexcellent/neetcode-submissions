func isValidSudoku(board [][]byte) bool {
    var rowCnt = len(board) // actually they are 9
    var columnCnt = len(board[0])
    for row := 0; row < rowCnt; row++ {
        seenRow := make(map[byte]bool, 9)
        for column := 0; column < columnCnt; column ++ {
            if board[row][column] != '.' {
                if seenRow[board[row][column]] == true {
                    return false
                } 
                seenRow[board[row][column]] = true
            }
        }
        fmt.Println(seenRow)
    }
    
    for column := 0; column < columnCnt; column ++ {
        seenColumn := make(map[byte]bool, 9)
        for row := 0; row < rowCnt; row ++ {
            if board[row][column] != '.' {
                if seenColumn[board[row][column]] == true {
                    return false
                } 
                seenColumn[board[row][column]] = true
            }
        }
        fmt.Println(seenColumn)
    }
    for n := 0; n < 3; n ++ {
        for m := 0; m < 3; m ++ {
            seenBox := make(map[byte]bool, 9)
            for j := 0; j < 3; j ++ {
                for i := 0; i < 3; i++ {
                    if board[n*3 + i][m*3 + j] != '.' {
                        if seenBox[board[n*3 + i][m*3 + j]] == true {
                        return false
                    } 
                    seenBox[board[n*3 + i][m*3 + j]] = true
                    }
                } 
            }
            fmt.Println(seenBox)
        }
    }
    return true
}


