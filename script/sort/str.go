package main

import "fmt"

func main() {
	dp := make([]int,6,1)
	fmt.Println(dp)

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


