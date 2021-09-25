func maxProfit(prices []int) int {
	profitOne0 := 0
	profitOne1 := -prices[0]
	profitTwo0 := 0
	profitTwo1 := -prices[0]

	for i := 1; i < len(prices); i++ {
		profitTwo0 = max(profitTwo0, profitTwo1+prices[i])
		profitTwo1 = max(profitTwo1, profitOne0-prices[i])
		profitOne0 = max(profitOne0, profitOne1+prices[i])
		profitOne1 = max(profitOne1, -prices[i])
	}
	return profitTwo0
}