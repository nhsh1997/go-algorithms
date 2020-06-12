package two_sum_solutions

func BruteForce(nums []int, target int) []int {
	indices := make([]int, 2)

	N := len(nums)

	for i := 0; i < N - 1; i++{
		for j := i + 1; j < N; j++ {
			if nums[i] + nums[j] == target {
				indices[0], indices[1] = i, j
				break
			}
		}
	}

	return indices
}