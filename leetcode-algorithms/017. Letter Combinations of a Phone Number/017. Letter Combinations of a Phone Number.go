func letterCombinations(digits string) []string {
	if len(digits) == 0 || len(digits) > 4 {
		return []string{}
	}
	digitsMap := [10]string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
	res := make([]string, 0)
	recursion("", digits, 0, digitsMap, &res)
	return res
}

func recursion(tmpStr string, digits string, index int, digitsMap [10]string, res *[]string) {
	if len(tmpStr) == len(digits) {
		*res = append(*res, tmpStr)
		return
	}

	tmpNum, _ := strconv.Atoi(string(digits[index]))
	letters := digitsMap[tmpNum]

	for i := 0; i < len(letters); i++ {
		tmpStr += string(letters[i])
		recursion(tmpStr, digits, index+1, digitsMap, res)
		tmpStr = tmpStr[:len(tmpStr)-1]
	}
}