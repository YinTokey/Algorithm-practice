func maxProfit(prices []int) int {
	re := 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			re += prices[i+1] - prices[i]
		}
	}
	return re
}