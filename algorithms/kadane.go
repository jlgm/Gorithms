func maxSubArray(nums []int) int {
	prev := nums[0]
	ans := prev
	for i := 1; i < len(nums); i++ {
		cand := nums[i]
		if prev+cand > cand {
			prev = cand + prev
		} else {
			prev = cand
		}
		if prev > ans {
			ans = prev
		}
	}

	return ans
}