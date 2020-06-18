package two_sum_solutions


func HashMap1(nums []int, target int) []int {
	mapNums := make(map[int]int)

	for index, value := range nums {
		search, ok := mapNums[value]
		if ok {
			return []int{index, search}
		} else {
			complement := target - value
			mapNums[complement] = index
		}
	}

	return nil
}