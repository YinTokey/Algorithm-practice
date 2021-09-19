func findLength(nums1 []int, nums2 []int) int {
	len1 := len(nums1)
	len2 := len(nums2)
	// 1. dp 初始化 (1.1 构造  1.2 设置初始值)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	// 默认 dp[0][0] = 0
	res := 0

	// 3. 开始遍历
	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			// dp[i][j] 与 dp[i-1][j-1] 关系
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			res = max(res, dp[i][j])
		}
	}

	return res
}