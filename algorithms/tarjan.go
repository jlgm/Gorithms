var (
	adjList map[int][]int
	low     map[int]int
	disc    map[int]int
	mark    map[int]bool
	cur     int
	ans     [][]int
)

func dfs(v, p int) {
	mark[v] = true
	disc[v] = cur
	low[v] = cur
	cur++
	for _, w := range adjList[v] {
		if w == p {
			continue
		}
		if mark[w] {
			if disc[w] < low[v] { //back-edge
				low[v] = disc[w]
			}
		} else {
			dfs(w, v)
			if low[w] < low[v] {
				low[v] = low[w]
			}
		}
		if low[w] > disc[v] {
			ans = append(ans, []int{v, w})
		}
	}
}

func findBridges(n int, connections [][]int) [][]int {
	low, disc = make(map[int]int), make(map[int]int)
	mark, cur = make(map[int]bool), 0
	adjList = make(map[int][]int)
	for i := 0; i < n; i++ {
		adjList[i] = make([]int, 0)
	}

	for _, c := range connections {
		v, w := c[0], c[1]
		adjList[v] = append(adjList[v], w)
		adjList[w] = append(adjList[w], v)
	}
	ans = make([][]int, 0)
	dfs(0, -1)

	return ans
}