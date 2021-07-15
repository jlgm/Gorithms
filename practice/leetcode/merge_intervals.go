// https://leetcode.com/problems/merge-intervals
type interval struct {
	i    int
	j    int
	show bool
}

func canMerge(a, b interval) bool {
	return b.i <= a.j
}

func merge(intervals [][]int) [][]int {
	m := make([]interval, 0)
	for _, in := range intervals {
		t := interval{
			i:    in[0],
			j:    in[1],
			show: true,
		}
		m = append(m, t)
	}
	sort.Slice(m, func(i, j int) bool {
		if m[i].i == m[j].i {
			return m[i].j < m[j].j
		}
		return m[i].i < m[j].i
	})

	curLead := 0
	for i := 1; i < len(m); i++ {
		if canMerge(m[curLead], m[i]) {
			m[i].show = false
			if m[i].j > m[curLead].j {
				m[curLead].j = m[i].j
			}
		} else {
			curLead = i
		}
	}

	ret := make([][]int, 0)
	for _, in := range m {
		if in.show {
			ar := []int{in.i, in.j}
			ret = append(ret, ar)
		}
	}

	return ret
}