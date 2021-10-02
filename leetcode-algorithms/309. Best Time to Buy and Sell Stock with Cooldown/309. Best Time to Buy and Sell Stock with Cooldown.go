func maxProfit(prices []int) int {
	profit0 := 0
	preProfit0 := 0
	profit1 := -prices[0]

	for i := 1; i < len(prices); i++ {
		nextProfit0 := max(profit0, profit1+prices[i])
		nextProfit1 := max(profit1, preProfit0-prices[i])
		preProfit0 = profit0
		profit0 = nextProfit0
		profit1 = nextProfit1
	}
	return profit0
}