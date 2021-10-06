func permuteUnique(nums []int) [][]int {

	path, res, used := make([]int, 0), make([][]int, 0), make([]bool, len(nums))
	sort.Ints(nums)
	recursion(nums, 0, path, &res, &used)
	return res
}

func recursion(nums []int, index int, path []int, res *[][]int, used *[]bool) {

	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if (*used)[i] == false {
			// 同级隔壁去重
			if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] { // 这里是去重的关键逻辑
				continue
			}
			(*used)[i] = true
			path = append(path, nums[i])
			recursion(nums, index+1, path, res, used)
			path = path[:len(path)-1]
			(*used)[i] = false
		}
	}
}