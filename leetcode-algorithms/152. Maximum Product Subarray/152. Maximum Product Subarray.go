func maxProduct(nums []int) int {
	length := len(nums)
	dp_max := make([]int, length)
	dp_min := make([]int, length)

	dp_max[0] = nums[0]
	dp_min[0] = nums[0]
	res := dp_max[0]

	for i := 1; i < length; i++ {
		if nums[i] < 0 {
			dp_max[i-1], dp_min[i-1] = dp_min[i-1], dp_max[i-1]
		}

		dp_max[i] = max(dp_max[i-1]*nums[i], nums[i])
		dp_min[i] = min(dp_min[i-1]*nums[i], nums[i])

		res = max(res, dp_max[i])
	}
	return res
}