//https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var parent map[*TreeNode]*TreeNode

 func gen(node *TreeNode, prev *TreeNode) {
	 if node == nil {
		 return
	 }
	 if prev != nil {
		 parent[node] = prev
	 }
	 gen(node.Left, node)
	 gen(node.Right, node)
 }
 
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	 parent = make(map[*TreeNode]*TreeNode)
	 var prev *TreeNode
	 gen(root, prev)
	 P, Q := make([]*TreeNode, 0), make([]*TreeNode, 0)
	 
	 P = append(P, p)
	 for p != root {
		 P = append(P, parent[p])
		 p = parent[p]
	 }
	 Q = append(Q, q)
	 for q != root {
		 Q = append(Q, parent[q])
		 q = parent[q]
	 }
	 
	 for _, elemP := range P {
		 for _, elemQ := range Q {
			 if elemP == elemQ {
				 return elemP
			 }
		 }
	 }
	 
	 return root
 }