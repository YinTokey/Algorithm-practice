func maxUncrossedLines(nums1 []int, nums2 []int) int {
	len1, len2 := len(nums1), len(nums2)
	dp := make([][]int, len1+1)
	for i, _ := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len1][len2]
}