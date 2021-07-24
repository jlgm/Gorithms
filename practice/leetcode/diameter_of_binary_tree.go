/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
	ans int
)

func gen(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	a, b := gen(root.Left), gen(root.Right)
	res := a + b
	if res > ans {
		ans = res
	}
	return 1 + max(a, b)
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans = 0
	_ = gen(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}