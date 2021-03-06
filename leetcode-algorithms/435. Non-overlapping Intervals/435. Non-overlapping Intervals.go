import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	count := 0
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
		} else {
			if intervals[i][1] < end {
				end = intervals[i][1]
			}
			count++
		}
	}
	return count
}

// new

// 【灵魂】计算不重叠区间
func intervalSchedule(intvs [][]int) int {
	if len(intvs) == 0 {
		return 0
	}

	// 按照end升序排列
	sort.Slice(intvs, func(i, j int) bool {
		return intvs[i][1] < intvs[j][1]
	})

	// 至少有一个区间不相交
	count := 1
	// 排序过后第一个区间就是x
	xEnd := intvs[0][1]

	for _, v := range intvs {
		start := v[0]
		// 开始的数字大于等于上一个结束的，所以说明是不重叠的
		if start >= xEnd {
			count++
			// 更新xEnd为下一个无重叠的end
			xEnd = v[1]
		}
	}

	return count
}

func eraseOverlapIntervals(intervals [][]int) int {
	// 已经找到了不重叠区间的个数，所以总的减去这个就得到了要剔除的个数
	return len(intervals) - intervalSchedule(intervals)
}

