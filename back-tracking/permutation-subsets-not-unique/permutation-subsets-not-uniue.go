package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println(permutationSubsetsNotUnique([]int{1, 2, 1, 1, 3}))
}

func permutationSubsetsNotUnique(choices []int) [][]int {
	var result [][]int

	// sort choices and the filter duplicated choices
	sort.Ints(choices)
	fmt.Println(choices)

	backtrack([]int{}, choices, &result)

	return result
}

func backtrack(list, remainingChoices[]int, result *[][]int)  {
	// loop through all remaining choices
	for i := 0; i < len(remainingChoices); i++ {
		//try the candidate
		var newList []int

		newList = copy(newList, list)

		newList = append(list, remainingChoices[i])

		*result = append(*result, newList)

		var arr []int
		arr = copy(arr, remainingChoices)

		arr = slice(arr, i)

		backtrack(newList, arr, result)

		for i + 1 < len(remainingChoices) && remainingChoices[i] == remainingChoices[i + 1] {
			i++
		}
	}
}

func copy(dst, src []int) []int {
	for _, v := range src {
		dst = append(dst, v)
	}
	return dst
}

func slice(arr []int, i int) []int  {
	return append(arr[:i], arr[(i + 1):]...)
}