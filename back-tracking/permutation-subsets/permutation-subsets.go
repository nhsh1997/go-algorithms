package main

import "fmt"

func main()  {
	fmt.Println(permutationSubsets([]int{1, 3, 3}))

}

/**
** A backtracking algorithm to find the permutation subsets from the choices
** Each number choice is a unique natural number, and it is used exactly once.
** @param {[]int} The choice array of unique natural numbers
** @return {[][]int} The permutation subsets from the choices
 */
func permutationSubsets(choices []int) [][]int {
	var result [][]int

	backtrack([]int{}, choices, &result)

	return result
}

func backtrack(list, remainingChoices []int, result *[][]int)  {
	for i := 0; i < len(remainingChoices) ; i++ {
		var newList []int
		newList = copy(newList, list)
		newList = append(newList, remainingChoices[i])

		*result = append(*result, newList)

		var arr []int
		arr = copy(arr, remainingChoices)

		arr = slice(arr, i)

		backtrack(newList, arr, result)
	}
}

func copy(dst, src []int) []int {
	for _, v := range src {
		dst = append(dst, v)
	}
	return dst
}

func slice(arr []int, i int) []int  {
	// Remove the element at index i from a.
	arr[i] = arr[len(arr)-1] // Copy last element to index i.
	arr[len(arr)-1] = 0   // Erase last element (write zero value).
	arr = arr[:len(arr)-1]   // Truncate slice.

	return arr
}