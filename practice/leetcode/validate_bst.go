/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func mountSlice(root *TreeNode, mount *[]int) {
	if root == nil {
		return
	}
	mountSlice(root.Left, mount)
	*mount = append(*mount, root.Val)
	mountSlice(root.Right, mount)
}

func isValidBST(root *TreeNode) bool {
	mount := make([]int, 0)
	mountSlice(root, &mount)
	for i := 1; i < len(mount); i++ {
		if mount[i] <= mount[i-1] {
			return false
		}
	}
	return true
}