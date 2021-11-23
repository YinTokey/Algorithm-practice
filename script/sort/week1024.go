package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	n := 1
	fmt.Println(nextBeautifulNumber(n))

}




func nextBeautifulNumber(n int) int {
	n++
	for n <= 1000000 {
		arr := numToArr(n)
		for idx, v  := range arr {
			if !isBalance(n,v) {
				break
			}
			if idx == len(arr)-1 {
				return n
			}
		}
		n++
	}
	return 0
}

func isBalance(num, c int) bool {
	count := 0
	arr := numToArr(num)
	for _,v := range arr {
		if v == c {
			count++
		}
	}

	return count == c
}

func numToArr(num int) []int {
	re := make([]int,0)
	for num > 0 {
		a := num%10
		re = append(re,a)
		num /= 10
	}
	return re
}

func countValidWords(sentence string) int {
	count := 0
	arr := strings.Fields(sentence)
	for _,v := range arr {
		vali := valid(v)
		fmt.Println(vali)
		if vali {
			count++
		}
	}

	fmt.Println(count)
	return count
}

func valid(str string) bool {
	b := []byte(str)
	connectCount := 0
	for idx,v := range b {
		// 数字
		_, err := strconv.Atoi(string(v))
		if err == nil {
			return false
		}

		// 连接符
		if v == '-' {
			connectCount++
			if connectCount > 1 {
				return false
			}
			if idx == 0 || idx == len(b)-1 {
				return false
			}
			innerArr := strings.Split(str,"-")
			for _,innerV := range innerArr {
				innerB := []byte(innerV)
				for inidx, ib := range innerB {
					if !isLetter(ib) {
						if (ib == '!' || ib == '.' || ib == ',') && inidx == len(innerB)-1{

						} else {
							return false
						}

					}
				}
			}

		}

		// 标点
		if v == '!' || v == '.' || v == ',' {
			if idx != len(b)-1 {
				return false
			}
 		}

	}

	return true

}



func isLetter(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
		return true
	} else {
		return false
	}
}
