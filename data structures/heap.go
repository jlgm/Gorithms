package main

import "fmt"

type heap struct {
	q    []int
	size int
}

func NewHeap() heap {
	return heap{
		q:    make([]int, 1),
		size: 0,
	}
}

func (h *heap) heapify(i int) {
	l, r := h.left(i), h.right(i)
	largest := 0
	if l <= h.size && h.q[l] > h.q[i] {
		largest = l
	} else {
		largest = i
	}
	if r <= h.size && h.q[r] > h.q[largest] {
		largest = r
	}
	if largest != i {
		h.q[i], h.q[largest] = h.q[largest], h.q[i]
		h.heapify(largest)
	}
}

func (h *heap) makeHeap() {
	for i := h.size / 2; i >= 1; i-- {
		h.heapify(i)
	}
}

func (h *heap) Add(val int) {
	h.q = append(h.q, val)
	h.size++

	h.makeHeap()
}

func (h *heap) Top() int {
	return h.q[1]
}

func (h *heap) RemoveTop() int {
	top := h.q[1]
	h.q[1] = 0
	h.q = h.q[1:]
	h.size -= 1
	h.makeHeap()

	return top
}

func (h *heap) left(i int) int {
	return i << 1
}

func (h *heap) right(i int) int {
	return (i << 1) + 1
}

func (h *heap) parent(i int) int {
	return i / 2
}

func (h *heap) print() {
	for i := 1; i <= h.size; i++ {
		fmt.Print(h.q[i], " ")
	}
	fmt.Println("")
}

func findRelativeRanks(score []int) []string {
	ans := make([]string, 0)

	heap := NewHeap()
	heap.Add(2)
	heap.Add(5)
	heap.Add(8)
	heap.Add(7)
	heap.Add(9)

	fmt.Println(heap.RemoveTop())
	fmt.Println(heap.RemoveTop())
	fmt.Println(heap.RemoveTop())
	fmt.Println(heap.RemoveTop())
	heap.print()

	return ans
}
