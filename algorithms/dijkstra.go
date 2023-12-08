const INF = math.MaxInt32

type Edge struct {
    v int 
    w int
}

func networkDelayTime(times [][]int, n int, k int) int {
    // initialization variables
    d, ans := make([]int, n), 0; 
    k-=1 //force 0-index
    
    // build adjList
    adjList := make([][]*Edge, n)
    for i := 0; i < len(times); i++ {
        u, v, w := times[i][0]-1, times[i][1]-1, times[i][2]
        adjList[u] = append(adjList[u], &Edge{
            v: v,       
            w: w,
        })
    }

    // initialize distance vector
    for i := 0; i < n; i++ {
        d[i] = INF
    }
    d[k] = 0

    // initialize priority queue
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    heap.Push(&pq, &Edge{k,0})
    
    // dijkstra's algorithm:
    for pq.Len() > 0 {
        top := heap.Pop(&pq).(*Edge)
        u := top.v
        for _, adj := range adjList[u] {
            // relaxation step
            if d[u] + adj.w < d[adj.v] {
                d[adj.v] = d[u] + adj.w
                heap.Push(&pq, &Edge{adj.v, d[u] + adj.w})
            }
        }
    }

    // calculate ans
    for i := 0; i < n; i++ {
        if d[i] > ans {
            ans = d[i]
        }
    }

    if ans == INF {
        return -1
    }
    return ans
}

/* Data Structures */
type PriorityQueue []*Edge
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].w < pq[j].w
}
func (pq PriorityQueue) Swap(i, j int) { 
    pq[i], pq[j] = pq[j], pq[i] 
}
func (pq *PriorityQueue) Push(x any) {
    item := x.(*Edge)
    *pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
