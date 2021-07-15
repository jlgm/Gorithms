// https://leetcode.com/problems/top-k-frequent-elements
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if count[nums[i]] == count[nums[j]] {
			return nums[i] < nums[j]
		}
		return count[nums[i]] > count[nums[j]]
	})

	ret := make([]int, 0)
	ret = append(ret, nums[0])
	last := nums[0]
	for i := 1; i < len(nums) && len(ret) < k; i++ {
		if nums[i] != last {
			ret = append(ret, nums[i])
			last = nums[i]
		}
	}

	return ret
}