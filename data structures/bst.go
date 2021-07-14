package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTree struct {
	root *TreeNode
}

func insertNode(root *TreeNode, node *TreeNode) {
	if node.Val < root.Val {
		if root.Left == nil {
			root.Left = node
		} else {
			insertNode(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
		} else {
			insertNode(root.Right, node)
		}
	}
}

func (this *BSTree) insert(node *TreeNode) {
	if this.root == nil {
		this.root = node
	} else {
		insertNode(this.root, node)
	}
}

func (this *BSTree) inOrderPrint(node *TreeNode) {
	if node == nil {
		return
	}
	this.inOrderPrint(node.Left)
	fmt.Println(node.Val)
	this.inOrderPrint(node.Right)
}

func (this *BSTree) preOrderPrint(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	this.preOrderPrint(node.Left)
	this.preOrderPrint(node.Right)
}

// func main() {
// 	tree := &BSTree{}
// 	tree.insert(&TreeNode{Val: 5})
// 	tree.insert(&TreeNode{Val: 2})
// 	tree.insert(&TreeNode{Val: 8})
// 	tree.insert(&TreeNode{Val: 1})
// 	tree.insert(&TreeNode{Val: 7})
// 	tree.insert(&TreeNode{Val: 3})
// 	tree.insert(&TreeNode{Val: 15})
// 	tree.insert(&TreeNode{Val: 13})
// 	tree.insert(&TreeNode{Val: 12})
// 	tree.insert(&TreeNode{Val: 14})

// 	tree.inOrderPrint(tree.root)
// }
