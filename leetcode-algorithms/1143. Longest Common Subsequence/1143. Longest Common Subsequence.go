func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	dp := make([]int, m+1)

	for i := 1; i <= n; i++ {
		prev := 0
		for j := 1; j <= m; j++ {
			tmp := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = prev + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = tmp
		}
	}
	return dp[m]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}