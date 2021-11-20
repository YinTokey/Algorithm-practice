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

func findLHS(nums []int) int {
	mp := make(map[int]int, 0)
	for _, val := range nums {
		mp[val]++
	}
	var ans int
	for key, val := range mp {
		c1, ok := mp[key+1]
		if ok && c1 > 0 && val+c1 > ans {
			ans = val + c1
		}
	}
	return ans
}
