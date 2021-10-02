func addBinary(a string, b string) string {

	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	if lenA > lenB {
		b = genZero(lenA-lenB) + b
	} else if lenB > lenA {
		a = genZero(lenB-lenA) + a
	}

	i := n - 1
	k := n
	res := make([]string, n+1)
	c := 0 // 进位

	for i >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		bi, _ := strconv.Atoi(string(b[i]))
		res[k] = strconv.Itoa((ai + bi + c) % 2)
		c = (ai + bi + c) / 2
		i--
		k--
	}
	if c > 0 {
		res[0] = strconv.Itoa(c)
	}

	return strings.Join(res, "")
}

func genZero(n int) string {
	s := ""
	for n > 0 {
		s += "0"
		n--
	}
	return s
}