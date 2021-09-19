func longestPalindrome(s string) string {
	length := len(s)
	// 二维切片初始化
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}

	for i := 0; i < length; i++ {
		dp[i][i] = 1
	}
	ans := string(s[0])

	for j := 1; j < length; j++ {
		for i := 0; i <= j; i++ {
			if s[i] == s[j] && (j-i < 3 || dp[i+1][j-1] > 0) {
				dp[i][j] = j - i + 1
			}
			if dp[i][j] > 0 && dp[i][j] > len(ans) {
				ans = string(s[i : j+1])
			}
		}
	}

	return ans
}