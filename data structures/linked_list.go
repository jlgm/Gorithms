package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
}

func (this *List) PushBack(node *ListNode) {
	if this.Head == nil {
		this.Head = node
	} else {
		it := this.Head
		for it.Next != nil {
			it = it.Next
		}
		it.Next = node
	}
}

func (this *List) PushFront(node *ListNode) {
	if this.Head == nil {
		this.Head = node
	} else {
		node.Next = this.Head
		this.Head = node
	}
}

func (this *List) PopFront() *ListNode {
	if this.Head == nil {
		return nil
	}
	it := this.Head
	this.Head = this.Head.Next
	return it
}

func (this *List) PopBack() *ListNode {
	if this.Head == nil {
		return nil
	}
	it := this.Head
	prev := it
	for it.Next != nil {
		prev = it
		it = it.Next
	}
	prev.Next = nil
	return it
}

func (this *List) PrintList() {
	it := this.Head
	for it != nil {
		fmt.Printf("%+v ", it.Val)
		it = it.Next
	}
	fmt.Println()
}

func (this *List) PrintReverse(node *ListNode) {
	if node == nil {
		return
	}
	this.PrintReverse(node.Next)
	fmt.Printf("%+v ", node.Val)
	if node == this.Head {
		fmt.Println()
	}
}

func (this *List) ReverseList(node *ListNode) {
	if node.Next == nil {
		this.Head = node
		return
	}
	this.ReverseList(node.Next)
	q := node.Next
	q.Next = node
	node.Next = nil
}

// func main() {
// 	list := &List{}
// 	list.PushBack(&ListNode{Val: 1})
// 	list.PushBack(&ListNode{Val: 2})
// 	list.PushBack(&ListNode{Val: 4})
// 	list.PushBack(&ListNode{Val: 8})
// 	list.PushFront(&ListNode{Val: -12})
// 	list.PushBack(&ListNode{Val: 15})

// 	list.PrintList()

// 	fmt.Println("Removing from front:")
// 	fmt.Println(list.PopFront().Val)
// 	fmt.Println(list.PopFront().Val)

// 	fmt.Println("Removing from back:")
// 	fmt.Println(list.PopBack().Val)
// 	fmt.Println(list.PopBack().Val)

// 	fmt.Println("Adding 9 & 13 on the back")
// 	list.PushBack(&ListNode{Val: 9})
// 	list.PushBack(&ListNode{Val: 13})

// 	fmt.Println("Print normal:")
// 	list.PrintList()
// 	fmt.Println("Print reverse:")
// 	list.PrintReverse(list.Head)

// 	fmt.Println("Reverse the list and print normal:")
// 	list.ReverseList(list.Head)
// 	list.PrintList()

// }
