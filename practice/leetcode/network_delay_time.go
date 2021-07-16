// https://leetcode.com/problems/network-delay-time
func networkDelayTime(times [][]int, n int, k int) int {
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0

	for i := 0; i < n; i++ {
		for _, t := range times {
			u, v, w := t[0], t[1], t[2]
			if dist[v] > dist[u]+w {
				dist[v] = dist[u] + w
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ { //dist[i] contains shortest path k-to-i
		if ans < dist[i] {
			ans = dist[i]
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}