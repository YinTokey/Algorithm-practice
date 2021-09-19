func findLengthOfLCIS(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	dp[0] = 1

	result := dp[0]

	for i := 1; i < length; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		result = max(result, dp[i])
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}