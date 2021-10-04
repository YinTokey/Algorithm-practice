package main

import (
	"fmt"
	"sort"
)

func main()  {
	nums := []int{3,2,2,3}
	val := 3
	removeElement(nums,val)
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	l, r := len(nums)-1, len(nums)-1

	for l >= 0 {
		if nums[l] == val {
			l--
		} else {
			nums[r] = nums[l]
			l--
			r--
		}
	}
	re := nums[r+1:]
	nums = *nums[r+1:]
	fmt.Println(re)
	return len(nums)
}

func combinationSum(candidates []int, target int) [][]int {

	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findcombinationSum(candidates, target, 0, c, &res)
	return res

}

func findcombinationSum(nums []int, target, index int, c []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		b := make([]int,len(c))
		copy(b,c)
		*res = append(*res,b)
		return
	}
	for i := index; i < len(nums); i++ {
		// 剪枝优化
		if nums[i] > target {
			break
		}
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		// 用于下一层递归的c
		nextC := append(c,nums[i])
		findcombinationSum(nums,target-nums[i],i+1,nextC,res)
		// c = c[:len(c)-1] 如果不用 NextC，可以这么做，复用c，减掉append 的元素，可能会比较不好理解
	}
}


