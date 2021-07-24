type cell struct {
	x int
	y int
	d int
}

var (
	dx = []int{1, 0, -1, 0}
	dy = []int{0, 1, 0, -1}
)

func orangesRotting(grid [][]int) int {
	q := list.New()
	n, m, ans := len(grid), len(grid[0]), 0
	mark := make([][]bool, n)
	for i := 0; i < n; i++ {
		mark[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			if grid[i][j] == 2 {
				c := cell{i, j, 0}
				q.PushBack(c)
			}
		}
	}

	for q.Len() > 0 {
		e := q.Front()
		v := e.Value.(cell)
		mark[v.x][v.y] = true

		for i := 0; i < 4; i++ {
			x, y, d := v.x+dx[i], v.y+dy[i], v.d+1
			if x < 0 || x >= n || y < 0 || y >= m {
				continue
			}
			c := cell{x, y, d}
			if !mark[x][y] && grid[x][y] == 1 {
				q.PushBack(c)
				mark[x][y] = true
				if d > ans {
					ans = d
				}
			}
		}
		q.Remove(e)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && !mark[i][j] {
				return -1
			}
		}
	}

	return ans
}