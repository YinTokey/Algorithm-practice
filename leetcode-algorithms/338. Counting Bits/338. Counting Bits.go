func countBits(num int) []int {
	resutl := make([]int, num+1)
	//第一个为0，从1开始遍历
	// f(i) = f(i/2)+ i%2
	for i := 1; i <= num; i++ {
		resutl[i] = resutl[i>>1] + (i & 1)
	}
	return resutl
}