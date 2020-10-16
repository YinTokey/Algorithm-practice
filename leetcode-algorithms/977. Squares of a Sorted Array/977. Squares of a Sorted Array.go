func sortedSquares(A []int) []int {
	lastNegIdx := -1
	ans := make([]int, 0)

	for i := 0; i < len(A); i++ {
		if A[i] < 0 {
			lastNegIdx++
		}
	}

	for i, j := lastNegIdx, lastNegIdx+1; i >= 0 || j < len(A); {
		if i < 0 {
			ans = append(ans, A[j]*A[j])
			j++
		} else if j == len(A) {
			ans = append(ans, A[i]*A[i])
			i--
		} else if A[i]*A[i] < A[j]*A[j] {
			ans = append(ans, A[i]*A[i])
			i--
		} else {
			ans = append(ans, A[j]*A[j])
			j++
		}

	}

	return ans
}