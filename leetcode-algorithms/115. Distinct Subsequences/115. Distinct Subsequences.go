func numDistinct(s string, t string) int {
	len1, len2 := len(s), len(t)
	dp := make([][]int, len2+1)
	for i, _ := range dp {
		dp[i] = make([]int, len1+1)
	}
	// 初始化dp值
	for j := 0; j <= len1; j++ {
		// 首行全部为1
		dp[0][j] = 1
	}

	for j := 1; j <= len1; j++ {
		for i := 1; i <= len2; i++ {
			if s[j-1] == t[i-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len2][len1]
}