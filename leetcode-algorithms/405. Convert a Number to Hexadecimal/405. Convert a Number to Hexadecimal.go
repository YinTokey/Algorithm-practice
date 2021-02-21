func toHex(num int) string {
	if num < 0 {
		num += 4294967296
	}
	if num == 0 {
		return "0"
	}
	res := []int32{}
	hash := []int32{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	for num != 0 {
		temp := num % 16
		num = num / 16
		res = append([]int32{hash[temp]}, res...)
	}
	return string(res)
}
