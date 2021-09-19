func countSubstrings(s string) int {
	length := len(s)
	// 二维切片初始化
	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, length)
	}
	ans := 0
	for i := 0; i < length; i++ {
		dp[i][i] = true
		ans++
	}

	for j := 1; j < length; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] && (j-i < 3 || (dp[i+1][j-1] == true)) {
				dp[i][j] = true
			}
			if dp[i][j] == true {
				ans++
			}
		}
	}

	return ans

}