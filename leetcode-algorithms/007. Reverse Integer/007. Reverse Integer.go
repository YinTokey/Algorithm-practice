import "math"

func reverse(x int) int {

	re := 0
	flag := false
	if x < 0 {
		flag = true
		x = -x
	}
	for x > 0 {
		re = re*10 + x%10
		x /= 10
	}

	if re > math.MaxInt32 || re < math.MinInt32 {
		return 0
	}

	if flag == true {
		return -re
	} else {
		return re
	}
}