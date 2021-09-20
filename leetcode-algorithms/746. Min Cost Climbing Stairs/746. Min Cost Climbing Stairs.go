func minCostClimbingStairs(cost []int) int {

	// 遍历时，对于最后一个元素，分两种情况。 对比哪一种耗费更少
	// 1. 踩着最后一个跳过去
	// 2. 踩倒数第二个，越过最后一个
	a := 0 // 这个是最关键的
	b := min(cost[0], cost[1])
	c := b

	for i := 2; i < len(cost); i++ {
		c = min(b+cost[i], a+cost[i-1])
		a = b
		b = c
	}
	return c
}