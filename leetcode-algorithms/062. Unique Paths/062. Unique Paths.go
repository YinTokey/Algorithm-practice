func uniquePaths(m int, n int) int {

	opt := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				opt[j] = 1
			} else {
				opt[j] = opt[j-1] + opt[j]
			}
		}
	}
	return opt[n-1]
}