package main

import "fmt"

func main()  {
	fmt.Println(fibonacci(4))
}

func fibonacci(n int) int {
	if n < 0 {
		return 0
	}

	dp := map[int]int{}
	dp[0] = 0
	dp[1] = 1

	return findResult(n, dp)
}

func findResult(i int, dp map[int]int) int {
	_, ok := dp[i]
	if !ok {
		dp[i] = findResult(i - 1, dp) + findResult(i - 2, dp)
	}
	return dp[i]
}