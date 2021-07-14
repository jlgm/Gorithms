package main

import "fmt"

const (
	N = 10
)

var (
	p []int // parent (or representative) of set
	s []int // size (or rank) of set
)

func createSet(i int) {
	p[i] = i
	s[i] = 1
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
	p[i] = findSet(p[i]) //path compression
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

func printSet() {
	for i := 1; i < N; i++ {
		fmt.Printf("%d->%d ", i, findSet(i))
	}
	fmt.Println()
}

func main() {

	p = make([]int, N)
	s = make([]int, N)

	for i := 1; i < N; i++ { //create the sets
		createSet(i)
	}

	printSet() // all disjoint

	union(1, 2)
	union(2, 3)

	printSet() // {1,2,3}, others disjoint

	union(5, 6)
	union(7, 8)
	union(3, 7)

	printSet() // {1,2,3,7,8}, {5,6}, {4}, {9}

}
