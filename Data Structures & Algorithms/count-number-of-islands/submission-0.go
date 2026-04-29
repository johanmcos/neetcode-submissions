func numIslands(grid [][]byte) int {
    var islands int
	for i, row := range grid {
		for j, node := range row {
			if node == '1' {
				islands++
				dfs(grid, i, j)
			}
		}
	}

	return islands
}

func dfs(grid [][]byte, i, j int) {
	if i > len(grid) - 1 || i < 0 || j < 0 || j > len(grid[i]) - 1 {
		return
	}
	if grid[i][j] == '1' {
		grid[i][j] = '0'
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}
}
