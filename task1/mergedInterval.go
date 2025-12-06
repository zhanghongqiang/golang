package task1

func mergedInterval(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	merged := make([][]int, 0)

	for _, v := range intervals {
		startIndex, endIndex := v[0], v[1]
		if len(merged) == 0 || merged[len(merged)-1][1] < startIndex {
			merged = append(merged, []int{startIndex, endIndex})
		} else {
			if merged[len(merged)-1][1] < endIndex {
				merged[len(merged)-1][1] = endIndex
			}
		}
	}
	return merged
}
