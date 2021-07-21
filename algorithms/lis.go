func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	ans := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] > dp[i] && nums[j] < nums[i] {
				dp[i] = dp[j]
			}
		}
		dp[i] += 1
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}