package main

import "fmt"

func main()  {
	matrix := make([][]int, 2)
	matrix[0] = []int{1, 2}
	matrix [1] = []int{1, 4}
	fmt.Println(minimalPath(matrix))
}

func minimalPath(matrix [][]int) int {
	// save the 2d length
	colLen := len(matrix)
	rowLen := len(matrix[0])

	// initialize 2d array for memoization
	// dp[i][j] is the sum of the minimal path from the top left corner to the bottom right corner to point[i][j]
	dp := make([][]int, colLen)
	for i := 0; i < colLen; i++  {
		dp[i] = make([]int, rowLen)
	}

	return findResult(matrix, colLen, rowLen, dp, colLen - 1, rowLen - 1)
}

func findResult(matrix [][]int, colLen, rowLen int, dp [][]int, start, end int) int {

	if start == 0 && end == 0 {
		dp[start][end] = matrix[start][end]
	} else if start == 0 {
		dp[start][end] = findResult(matrix, colLen, rowLen, dp, start, end - 1) + matrix[start][end]
	} else if end == 0 {
		dp[start][end] = findResult(matrix, colLen, rowLen, dp, start - 1, end) + matrix[start][end]
	} else {
		dp[start][end] = min(findResult(matrix, colLen, rowLen, dp, start - 1, end), findResult(matrix, colLen, rowLen, dp, start, end - 1)) + matrix[start][end]
	}

	return dp[start][end]
}

func min( a, b int) int {
	if a > b {
		return b
	}
	return a
}