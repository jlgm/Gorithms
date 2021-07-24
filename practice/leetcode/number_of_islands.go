var (
	dx   = []int{0, 1, 0, -1}
	dy   = []int{1, 0, -1, 0}
	mark [][]bool
)

func dfs(grid *[][]byte, x, y int) {
	if mark[x][y] {
		return
	}
	cpy := *grid
	mark[x][y] = true
	for i := 0; i < 4; i++ {
		x1, y1 := x+dx[i], y+dy[i]
		if x1 < 0 || x1 >= len(cpy) || y1 < 0 || y1 >= len(cpy[0]) {
			continue
		}
		if cpy[x1][y1] == '1' {
			dfs(grid, x1, y1)
		}
	}
}

func numIslands(grid [][]byte) int {
	mark = make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		mark[i] = make([]bool, len(grid[i]))
	}

	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !mark[i][j] && grid[i][j] == '1' {
				dfs(&grid, i, j)
				ans = ans + 1
			}
		}
	}

	return ans

}