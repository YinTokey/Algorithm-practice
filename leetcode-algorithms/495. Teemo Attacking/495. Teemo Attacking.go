func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	if len(timeSeries) == 1 {
		return duration
	}
	timeSeries = append(timeSeries, timeSeries[len(timeSeries)-1]+duration)
	rs, s := 0, -1
	for i := 0; i < len(timeSeries)-1; i++ {
		if timeSeries[i+1]-timeSeries[i] < duration {
			if s == -1 {
				s = i
			}
			continue
		}
		if s != -1 {
			rs += timeSeries[i] - timeSeries[s] + duration
			s = -1
		} else {
			rs += duration
		}
	}
	return rs
}
