package main

import "fmt"

func main()  {
	fmt.Println(permutationCount([]int{1, 2, 3, 4}))
}

/*
 * Dynamic programming algorithm to calculate permutation of a choices
 * Each choice is an unique natural number, and it is used exactly once
 * @param {[]int} A choice array of unique natural number
 * @return {int} The permutation count
*/
func permutationCount(choices []int) int {

	dp := map[int]int{}
	dp[0] = 1

	for i := 1; i <= len(choices) ; i++  {
		dp[i] = i * dp[i - 1]
	}

	return dp[len(choices) - 1]
}
