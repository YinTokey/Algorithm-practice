package main

import (
	"fmt"
	"sort"

)

func main()  {
	re := multiply("99","9")
	fmt.Println(re)
}

func multiply(num1 string, num2 string) string {
	b1, b2 := []byte(num1),[]byte(num2)
	re := 0
	for i := len(b1)-1; i >= 0; i-- {
		tmp := 0
		for j := len(b2)-1; j >= 0; j-- {
			int1 := int(b1[i]-'0')
			int2 := int(b2[j]-'0')
			powT := len(b2)-1-j
			tmp += int1 * int2 * powTen(powT)
			// fmt.Println("loop1",int1,int2,powT,tmp)
		}
		// fmt.Println(tmp)
		tmp *= powTen(len(b2)-1-i)
		re += tmp
	}

	fmt.Println(re)

	arr := make([]byte,len(b1)+len(b2))
	for i := len(arr)-1; i >= 0; i-- {
		arr[i] = '0' + byte(re % 10)
		re /= 10
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] - '0' != 0 {
			arr = arr[i:]
			fmt.Println("cut",i)
			break
		}
	}

	return string(arr)
}

func powTen(x int) int {
	re := 1
	for x > 0 {
		re *= 10
		x--
	}
	return re
}

type obj struct {
	index int
	val int
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	re := make(map[int]bool)
	mp := make(map[int]int)

	for _,v := range nums1 {
		mp[v] = 1
	}

	for _,v := range nums2 {
		val, ok := mp[v]
		if ok && val != 2 {
			re[v] = true
		} else {
			mp[v] = 2
		}
	}

	for _,v := range nums3 {
		val, ok := mp[v]
		if ok && val != 3 {
			re[v] = true
		} else {
			mp[v] = 3
		}
	}

	keys := make([]int, len(re))

	i := 0
	for k := range re {
		keys[i] = k
		i++
	}

	return keys
}

func minOperations(grid [][]int, x int) int {

	mod := grid[0][0]
	vals := make([]int,0)
	for _,v := range grid {
		vals = append(vals,v...)
		for _, innerV := range v {

			if (innerV - mod) % x != 0 {
				return -1
			}
		}
	}

	sort.Ints(vals)

	ans := 0
	n := len(vals)
	for _, v := range vals {
		ans += getAbs(v-vals[n/2]) / x
	}

	return ans
}

func getAbs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func getMin(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func getMax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
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
	//nums = *nums[r+1:]
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


