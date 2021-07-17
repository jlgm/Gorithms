// https://leetcode.com/problems/longest-palindromic-substring
func longestPalindrome(s string) string {

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	max := -1
	ans := ""
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			dp[i][j] = (s[i] == s[j])
			if j-i > 2 {
				dp[i][j] = (dp[i][j] && dp[i+1][j-1])
			}
			if j-i > max && dp[i][j] {
				max = j - i
				ans = s[i : j+1]
			}
			// fmt.Println(i, j, s[i:j+1], dp[i][j])
		}
	}
	// fmt.Println(dp)
	return ans

}   