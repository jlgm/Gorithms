// https://leetcode.com/problems/count-of-smaller-numbers-after-self/submissions/

const (
	offs = 10000
)

func update(index, val int, tree *[]int) {
	index++
	for index < len(*tree) {
		(*tree)[index] += val
		index += index & -index
	}
}

func query(index int, tree *[]int) int {
	res := 0
	for index >= 1 {
		res += (*tree)[index]
		index -= index & -index
	}
	return res
}

func countSmaller(nums []int) []int {

	n := len(nums)
	tree := make([]int, 2*offs+2)
	res := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		idx := nums[i] + offs
		freq := query(idx, &tree)
		res[i] = freq
		update(idx, 1, &tree)
	}

	return res

}