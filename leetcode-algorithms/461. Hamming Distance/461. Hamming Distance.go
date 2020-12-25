func hammingDistance(x int, y int) int {
	i := x ^ y
	count := 0
	for i != 0 {
		if (i & 1) == 1 {
			count++
		}
		i = i >> 1
	}
	return count

}
