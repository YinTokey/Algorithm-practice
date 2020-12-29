func canIWin(maxChoosableInteger int, desiredTotal int) bool {

	if (maxChoosableInteger * (maxChoosableInteger + 1) / 2) < desiredTotal {
		return false
	}

	memo := make(map[int]bool)

	var dfs func(used int, curTotal int) bool
	dfs = func(used int, curTotal int) bool {

		if val, ok := memo[used]; ok {
			return val
		}

		for i := 1; i <= maxChoosableInteger; i++ {
			cur := (1 << i)
			if cur&used == 0 {
				if i >= curTotal || !dfs(used|cur, curTotal-i) {
					memo[used] = true
					return true
				}
			}
		}

		memo[used] = false
		return false
	}

	return dfs(0, desiredTotal)
}

