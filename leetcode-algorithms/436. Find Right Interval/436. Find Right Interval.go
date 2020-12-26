import "sort"

func findRightInterval(intervals [][]int) []int {

	res := make([]int, len(intervals))
	nodes := make([]int, len(intervals))
	for i := range nodes {
		nodes[i] = i
	}

	sort.Slice(nodes, func(i, j int) bool {
		return intervals[nodes[i]][0] < intervals[nodes[j]][0]
	})

	for i := range nodes {
		res[nodes[i]] = -1

		l, r := i, len(nodes)
		for l < r {
			mid := (l + r) >> 1
			if intervals[nodes[mid]][0] >= intervals[nodes[i]][1] {
				r = mid
			} else {
				l = mid + 1
			}
		}

		if l > i && l < len(nodes) {
			res[nodes[i]] = nodes[l]
		}
	}

	return res
}

