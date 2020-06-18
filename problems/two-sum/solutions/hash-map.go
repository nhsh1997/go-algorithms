package two_sum_solutions

func HashMap(nums []int, target int) []int {
	indices := make([]int, 2)
	mapNums := make(map[int]int)
	N := len(nums)

	for index, value := range nums {
		mapNums[value] = index
	}

	for i := 0; i < N; i++ {
		complement := target - nums[i]
		if mapNums[complement] != i {
			complementIndex, ok := mapNums[complement]
			if ok {
				indices[0] = i
				indices[1] = complementIndex
				break
			}
		}
	}
	return indices
}

