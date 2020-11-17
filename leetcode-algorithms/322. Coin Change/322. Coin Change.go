func coinChange(coins []int, amount int) int {
	// 二维切片初始化
	dp := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		dp[i] = make([]int, amount+1)

		for j := 0; j <= amount; j++ {
			if j == 0 {
				dp[i][j] = 0
				continue
			}
			if i == 0 {
				// if j % coins[i] != 0 {
				//     return -1
				// } else {
				dp[0][j] = j / coins[i]
				// }
			} else {
				if j >= coins[i] {
					dp[i][j] = min(dp[i-1][j], 1+dp[i][j-coins[i]])
				} else {
					dp[i][j] = dp[i-1][j]
				}

			}

		}
	}
	if dp[len(coins)-1][amount] > amount {
		return -1
	} else {
		return dp[len(coins)-1][amount]

	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}
