package main

import "fmt"

func main()  {
	fmt.Println(subset([]int{1, 2, 3}))
}

/**
* A backtracking algorithm to find the subsets from the choices
* Each choice is a natural unique number, it is used exactly once
* @param {[]int} choices The natural unique choices
* @return {[][]int} The subsets from the choices
*/
func subset(choices []int) [][]int {
	var result [][]int

	backtrack([]int{}, choices, 0, &result)

	return result
}

func backtrack(list, choices []int, start int, result *[][]int)  {
	for i := start; i < len(choices) ; i++ {
		var newList []int
		newList = copy(newList, list)
		newList = append(newList, choices[i])

		*result = append(*result, newList)

		backtrack(newList, choices, i + 1, result)
	}
}

func copy(dst, src []int) []int {
	for _, v := range src {
		dst = append(dst, v)
	}
	return dst
}