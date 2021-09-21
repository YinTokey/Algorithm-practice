n := len(coins)

dp := make([][]int, n+1)
for i := range dp {
	dp[i] = make([]int, amount+1)
}
//  第一列都为0，默认初始化时就是了，不用操作

// 设置一个无效值，因为求最小，所以无效值尽可能取非常大的
invalidNum := amount + 1
for j := 1; j <= amount; j++ {
	//当 i =0 , j != 0
	//	dp[0][x] = 无效值 因为计算过程中涉及到了min(,)操作，怎么避免无效值的干扰？我们设定无效值等于一个很大的不可能的值
	dp[0][j] = invalidNum
}
//其他为0的情况 我们就不用初始化了，数组默认0
for i := 1; i <= n; i++ {
	for j := 1; j <= amount; j++ {
		if j >= coins[i-1] {
			dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
		} else {
			dp[i][j] = dp[i-1][j]
		}
	}
}

if dp[n][amount] == invalidNum {
	return -1
} else {
	return dp[n][amount]
}