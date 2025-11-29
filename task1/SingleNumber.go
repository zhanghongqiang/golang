package task1

func singleNumber(nums []int) int {
	singleMap := make(map[int]int)
	for _, v := range nums {
		singleMap[v]++
	}

	for k, v := range singleMap {
		if v == 1 {
			return k
		}
	}

	return 0

}
