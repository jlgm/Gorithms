// https://leetcode.com/problems/connecting-cities-with-minimum-cost
var (
	p []int
	s []int
)

func createSet(i int) {
	p[i] = i
}

func getRank(i int) int {
	if p[i] == i {
		return s[i]
	}
	return s[findSet(i)]
}

func findSet(i int) int {
	if p[i] == i {
		return i
	}
	p[i] = findSet(p[i])
	return p[i]
}

func union(i, j int) {
	x, y := findSet(i), findSet(j)
	if x != y {
		rx, ry := getRank(x), getRank(y) //union by rank
		if rx > ry {
			s[p[x]] += ry
			p[y] = p[x]
		} else {
			s[p[y]] += rx
			p[x] = p[y]
		}
	}
}

func minimumCost(n int, connections [][]int) int {
	p = make([]int, n+1)
	s = make([]int, n+1)
	for i := 1; i <= n; i++ {
		createSet(i)
	}
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})
	ans := 0
	for _, e := range connections {
		u, v, w := e[0], e[1], e[2]
		if findSet(u) != findSet(v) {
			union(u, v)
			ans += w
		}
	}

	for i := 2; i <= n; i++ {
		if findSet(i) != findSet(i-1) {
			return -1
		}
	}
	return ans
}