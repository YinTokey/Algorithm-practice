func missingNumber(nums []int) int {
	re := 0
	for i := 0; i <= len(nums); i++ {
		re ^= i
	}
	for _, value := range nums {
		re ^= value
	}
	return re
}