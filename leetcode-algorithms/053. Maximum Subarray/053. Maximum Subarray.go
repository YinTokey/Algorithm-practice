func maxSubArray(nums []int) int {

    result := nums[0]
    for i := 1; i < len(nums); i++ {
        nums[i] = max(nums[i-1]+nums[i],nums[i])
        if nums[i] > result {
            result = nums[i]
        }
    }
    return result
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}