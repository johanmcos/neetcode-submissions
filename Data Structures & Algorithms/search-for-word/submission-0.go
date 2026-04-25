func exist(board [][]byte, word string) bool {
    for row := range board {
        for col := range board[row] {
            if dfs(board, word, row, col, 0) {
                return true
            }
        }
    }
    return false
}

func dfs(board [][]byte, word string, row, col, idx int) bool {
    if row < 0 || col < 0 || row >= len(board) || col >= len(board[row]) || board[row][col] != word[idx] {
        return false
    }
    
    if idx == len(word) - 1 {
        return true
    }

    temp := board[row][col]
    board[row][col] = 0

    result := dfs(board, word, row-1, col, idx+1) || dfs(board, word, row+1, col, idx+1) || dfs(board, word, row, col-1, idx+1) || dfs(board, word, row, col+1, idx+1)

    board[row][col] = temp

    return result
}
