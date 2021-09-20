func longestPalindromeSubseq(s string) int {
	length := len(s)
	// 二维切片初始化
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}
	for i := 0; i < length; i++ {
		dp[i][i] = 1
	}

	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- { // 这里j 要反向遍历，理解表格填写顺序
			if s[j] == s[i] {
				dp[i][j] = dp[i-1][j+1] + 2
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j+1])
			}
		}
	}

	return dp[length-1][0]
}