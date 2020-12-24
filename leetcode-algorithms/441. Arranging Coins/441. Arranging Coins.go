func arrangeCoins(n int) int {
	for i := 1; n >= 0; i++ {
		if n >= i {
			n -= i
		} else {
			return i - 1
		}
	}
	return 0
}
