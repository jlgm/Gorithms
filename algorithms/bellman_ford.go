type edge struct { //(u-v, w)
	u int
	v int
	w int
}

func bellmanFord(edges []edge, s int, n int) []int {
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[s] = 0
	for i := 0; i < n; i++ {
		for _, e := range edges {
			u, v, w := e.u, e.v, e.w
			if dist[v] > dist[u]+w {
				dist[v] = dist[u] + w
				//parent[v] = u (if path is needed)
			}
		}
	}

	return dist //dist[i] has the shortest path from s to i (or inf if they arent connected)
}