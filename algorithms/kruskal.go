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

type edge struct { //(u-v, w)
	u int
	v int
	w int
}

func kruskal(n int, edges []edge) []edge {
	p = make([]int, n+1)
	s = make([]int, n+1)
	for i := 1; i <= n; i++ {
		createSet(i) //create sets for each vertex
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	}) //sort edges by weight (increasing order)

	ret := make([]edge, 0)
	for _, e := range edges {
		u, v, w := e.u, e.v, e.w
		if findSet(u) != findSet(v) { //if edge is not connected, connect
			union(u, v)
			ret = append(ret, e)
		}
	}

	// ret will have the MST for the connected graph
	return ret
}