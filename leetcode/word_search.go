var (
	dx  = []int{1, 0, -1, 0}
	dy  = []int{0, 1, 0, -1}
	m   int
	n   int
	ans bool
)

func solve(mark [7][7]bool, board [][]byte, cur []byte, x, y int, target []byte) {
	if x < 0 || y < 0 || x >= n || y >= m {
		return
	}
	if mark[x][y] || ans {
		return
	}
	if len(cur) > len(target) {
		return
	}
	cur = append(cur, board[x][y])
	if bytes.Equal(cur, target) {
		ans = true
		return
	}
	mark[x][y] = true
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		solve(mark, board, cur, nx, ny, target)
	}
}

func exist(board [][]byte, word string) bool {
	ans = false
	n = len(board)
	m = len(board[0])
	wb := []byte(word)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var mark [7][7]bool
			solve(mark, board, []byte(""), i, j, wb)
			if ans {
				return true
			}
		}
	}
	return false
}