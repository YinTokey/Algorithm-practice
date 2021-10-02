package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := "11"
	b := "1"
	re := add(a,b)
	fmt.Println(re)
}

func add(a, b string) string {
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	if lenA > lenB {
		b = genZero(lenA-lenB) + b
	} else if lenB > lenA {
		a = genZero(lenB-lenA) + a
	}

	i := n - 1
	k := n
	res := make([]string,n+1)
	c := 0 // 进位

	for i >= 0 {
		ai,_ := strconv.Atoi(string(a[i]))
		bi,_ := strconv.Atoi(string(b[i]))
		res[k] = strconv.Itoa((ai+bi+c) % 2)
		c = (ai + bi + c) / 2
		i--
		k--
	}
	if c > 0 {
		res[0] = strconv.Itoa(c)
	}

	return strings.Join(res,"")

}

func genZero(n int) string {
	s := ""
	for n > 0 {
		s += "0"
		n--
	}
	return s
}


func pack(we, v []int,W int) {
	// 1. 老样子，二维dp初始化
	dp := make([][]int,len(we))
	for i,_ := range dp {
		dp[i] = make([]int,W+1)
		// 2.1 dp 局部初始化（第0 列），直接从这里搞定
		dp[i][0] = 0
	}
	// 2.2 局部初始第一行
	for j,_ := range dp[0] {
		if j > 0 {
			dp[0][j] = v[0]
		}
	}


	// 3. 两层遍历， dp 公式处理。 第1行，第1列都局部初始化完了，从1开始遍历
	for i := 1; i < len(we); i++ {
		for j := 1; j <= W; j++ {
			if j < we[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(v[i]+dp[i-1][j-we[i]],dp[i-1][j])
			}
		}
	}

	fmt.Println(dp[len(we)-1][W])

}

func max(x, y int ) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func KMP(str string, pattern string) bool {

	next := getNext(pattern)

	i, j:= 0,0

	for i < len(str) && j < len(pattern) {
		if str[i] == pattern[j] {
			i++
			j++
		} else if j != 0 {
			j = next[j-1]
		} else {
			i++
		}
		fmt.Println(j)
	}

	return j == len(pattern)

}

// https://www.youtube.com/watch?v=GTJr8OvyEVQ
func getNext(pattern string) []int {

	next := make([]int,len(pattern))
	next[0] = 0
	j := 0

	for i := 1; i < len(pattern); {
		if pattern[i] == pattern[j] {
			next[i] = j+1
			j++
			i++
		} else if j != 0 {
			j = next[j-1]
		} else {
			next[i] = 0
			i++
		}
	}

	return next
}

func isLetter(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
		return true
	} else {
		return false
	}
}

func transSmall(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		c += 'a'-'A'
	}
	return c
}



func min(x, y int) int  {
	if x < y {
		return x
	} else {
		return y
	}
}