func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b, c := 1, 1, 2

	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}