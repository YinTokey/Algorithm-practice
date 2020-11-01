func findLHS(nums []int) int {
	res := 0
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
		if count[num-1] != 0 && res < count[num-1]+count[num] {
			res = count[num-1] + count[num]
		}
		if count[num+1] != 0 && res < count[num+1]+count[num] {
			res = count[num+1] + count[num]
		}
	}
	return res
}
