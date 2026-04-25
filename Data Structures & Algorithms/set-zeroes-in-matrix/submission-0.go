func setZeroes(matrix [][]int) {
    // create arrays to represent the rows and columns
    rowZero := make([]bool, len(matrix)) // true -> has at least one zero
    colZero := make([]bool, len(matrix[0]))

    // iterate through the matrix to find zeros
    for rowNo, row := range matrix {
        for colNo, el := range row {
            if el == 0 {
                rowZero[rowNo] = true
                colZero[colNo] = true
            }
        }
    }

    // now set other elements to zero
    for rowNo, row := range matrix {
        for colNo := range row {
            if rowZero[rowNo] || colZero[colNo] {
                matrix[rowNo][colNo] = 0
            }
        }
    }
    
}
