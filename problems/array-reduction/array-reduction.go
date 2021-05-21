package main

import "fmt"

func main()  {
	fmt.Println(reductionCost([]int32{1, 2, 3, 4, 4}))
}


func reductionCost(num []int32) int32 {
	permutationNums := permutationNotUnique(num)
	fmt.Println(permutationNums)

	var result int32
	result = calculateCost(permutationNums[0])

	for i := 1; i < len(permutationNums); i++ {
		total := calculateCost(permutationNums[i])
		if total < result {
			result = total
		}
	}

	return result
}

func calculateCost(list []int32) int32  {
	var totals []int32
	var total int32
	for len(list) > 0  {
		total += list[0]
		totals = append(totals, total)
		list = list[1:]
	}

	var result int32

	for i := 1; i < len(totals); i++  {
		result += totals[i]
	}

	return result
}

func permutationNotUnique(choices []int32) [][]int32  {
	var result [][]int32
	backtrack([]int32{}, choices, len(choices), &result)
	return result
}

func backtrack(list, remainingChoices []int32, subsetLen int , result *[][]int32)  {
	for i := 0; i < len(remainingChoices) ; i++ {
		var newList []int32
		newList = copy(newList, list)
		newList = append(newList, remainingChoices[i])

		if len(newList) == subsetLen {
			*result = append(*result, newList)
		}

		var arr []int32
		arr = copy(arr, remainingChoices)

		arr = slice(arr, i)

		backtrack(newList, arr, subsetLen, result)
	}
}

func copy(dst, src []int32) []int32 {
	for _, v := range src {
		dst = append(dst, v)
	}
	return dst
}

func slice(arr []int32, i int) []int32  {
	// Remove the element at index i from a.
	arr[i] = arr[len(arr)-1] // Copy last element to index i.
	arr[len(arr)-1] = 0   // Erase last element (write zero value).
	arr = arr[:len(arr)-1]   // Truncate slice.

	return arr
}