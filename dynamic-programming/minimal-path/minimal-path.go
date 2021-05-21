package main

import "fmt"

func main()  {
	matrix := make([][]int, 2)
	matrix[0] = []int{1, 2}
	matrix [1] = []int{3, 4}
	fmt.Println(minimalPath(matrix))
}

/*
 * A dynamic programming algorithm to find the minimal path from the top left corner
 * to the bottom right corner based on the matrix values
 * @param {[][]int} The 2d array with weighted path value
 * @return {int} The sum of the minimal path
*/
func minimalPath(matrix [][]int) int {
	// save the 2d array lengths
	colLen := len(matrix)
	rowLen := len(matrix[0])

	// initialize the 2d array for memoization
	// dp[i][i] is the sum of the minimal path from the top left to the bottom right to point[i][j]
	dp := make([][]int, colLen)
	for i := 0; i < colLen; i++  {
		dp[i] = make([]int, rowLen)
	}

	for i := 0; i < colLen; i++ {
		for j := 0; j < rowLen; j++ {
			// at the top left corner, the minimal path is its own value
			if j == 0 && i == 0 {
				dp[i][j] = matrix[i][j]

				// at the start row, the minimal path is the left point's minimal path
				// plus it own value
			} else if i == 0 {
				dp[i][j] = dp[i][j - 1] + matrix[i][j]

				//at the start column, the minimal path is the above point's minimal path
				// plus it own value
			} else if j == 0{
				dp[i][j] = dp[i - 1][j] + matrix[i][j]

			} else {
				dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + matrix[i][j]
			}
		}
	}
	return dp[colLen - 1][ rowLen - 1]
}

func min( a, b int) int {
	if a > b {
		return b
	}
	return a
}

