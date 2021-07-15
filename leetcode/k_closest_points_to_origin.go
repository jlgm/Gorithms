// https://leetcode.com/problems/k-closest-points-to-origin
type point struct {
	x    int
	y    int
	dist float64
}

type PQ []*point //priority queue of points

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].dist > pq[j].dist //priority defined by dist to (0,0)
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	item := x.(*point)
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func kClosest(points [][]int, k int) [][]int {
	if len(points) <= k {
		return points
	}

	pq := make(PQ, k)
	for i := 0; i < k; i++ {
		x, y := points[i][0], points[i][1]
		pq[i] = &point{
			x:    x,
			y:    y,
			dist: math.Sqrt(float64(x*x + y*y)),
		}
	}

	heap.Init(&pq)
	for i := k; i < len(points); i++ {
		x, y := points[i][0], points[i][1]
		p := &point{
			x:    x,
			y:    y,
			dist: math.Sqrt(float64(x*x + y*y)),
		}
		heap.Push(&pq, p)
		heap.Pop(&pq)
	}

	ret := make([][]int, 0)

	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*point)
		ret = append(ret, []int{p.x, p.y})
	}
	return ret
}