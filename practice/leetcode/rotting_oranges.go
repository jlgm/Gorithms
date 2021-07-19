var (
	dx = []int{0, 1, 0, -1}
	dy = []int{1, 0, -1, 0}
)

type pair struct {
	x int
	y int
	d int
}

func orangesRotting(grid [][]int) int {

	m := make([][]bool, len(grid))
	for i := range m {
		m[i] = make([]bool, len(grid[0]))
	}
	q := list.New()
	ans := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 2 {
				continue
			}
			q.PushBack(pair{i, j, 0})
		}
	}
	for q.Len() > 0 {
		e := q.Front()
		v := e.Value.(pair)
		if !m[v.x][v.y] && v.d > ans {
			ans = v.d
		}
		m[v.x][v.y] = true

		for k := 0; k < len(dx); k++ {
			x, y, d := v.x+dx[k], v.y+dy[k], v.d+1
			if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
				continue
			}
			if !m[x][y] && grid[x][y] == 1 {
				q.PushBack(pair{x, y, d})
			}
		}
		q.Remove(e)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 && !m[i][j] {
				return -1
			}
		}
	}

	return ans
}