package main

import (
	"container/list"
	"fmt"
)

type graph struct {
	adjList map[int][]int
	visited map[int]bool
}

func newGraph() *graph {
	return &graph{
		adjList: map[int][]int{},
		visited: map[int]bool{},
	}
}

func (g *graph) createNode(v int) {
	if g.adjList[v] == nil {
		g.adjList[v] = make([]int, 0)
	}
}

func (g *graph) addEdge(v, w int) {
	g.createNode(v)
	g.createNode(w)
	g.adjList[v] = append(g.adjList[v], w)
	g.adjList[w] = append(g.adjList[w], v)
}

func (g *graph) reset() {
	g.visited = map[int]bool{}
}

func (g *graph) dfs(v int) {
	if g.visited[v] {
		return
	}
	fmt.Println("(dfs) visiting node", v)
	g.visited[v] = true
	for _, w := range g.adjList[v] {
		g.dfs(w)
	}
}

func (g *graph) bfs() {
	q := list.New()
	q.PushBack(1) 
	for q.Len() > 0 {
		e := q.Front()
		v := e.Value.(int)
		fmt.Println("(bfs) visiting node", v)
		g.visited[v] = true
		for _, w := range g.adjList[v] {
			if !g.visited[w] {
				q.PushBack(w)
			}
		}
		q.Remove(e)
	}
}

// func main() {
// 	g := newGraph()
// 	g.addEdge(1,2)
// 	g.addEdge(1,3)
// 	g.addEdge(5,4)
// 	g.addEdge(4,7)
// 	g.addEdge(4,2)
// 	g.dfs(1)
// 	fmt.Println()
// 	g.reset()
// 	g.bfs()
// }