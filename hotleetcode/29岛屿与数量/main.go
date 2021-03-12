package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	count := 0
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			if grid[x][y] == '1' {
				count++
				dfs(grid, x, y)
			}
		}
	}
	return count
}
func dfs(grid [][]byte, x, y int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	dfs(grid, x-1, y)
	dfs(grid, x+1, y)
	dfs(grid, x, y-1)
	dfs(grid, x, y+1)
}
