func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	opt := make([]int, m)

	if obstacleGrid[0][0] == 0 {
		opt[0] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				opt[j] = 0
				continue
			}

			if j >= 1 && obstacleGrid[i][j-1] == 0 {
				opt[j] += opt[j-1]
			}
		}
	}

	return opt[m-1]
}