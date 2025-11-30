package task1

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for k, v := range nums {
		rightNum := target - v
		if index, exists := sumMap[rightNum]; exists {
			return []int{index, k}
		}

		sumMap[v] = k
	}

	return nil
}
