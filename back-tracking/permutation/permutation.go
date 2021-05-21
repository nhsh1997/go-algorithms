package main

import "fmt"

func main()  {
	fmt.Println(permutation([]int{1, 2, 3}))
}


func permutation(choices []int) [][]int  {
	var result [][]int
	backtrack(choices, []int{}, 0, &result)
	return result
}

func backtrack(choices, list []int, start int, result *[][]int) {
	//reach the full length
	if len(list) == len(choices) {
		fmt.Println(list)
		*result = append(*result, list)
	} else {
		//from the start candidate to the last one
		for i := start; i < len(choices) ; i++  {
			if !includes(list, choices[i]) {
				//init new list
				var newList []int
				newList = copy(newList, list)
				newList = append(newList, choices[i])
				backtrack(choices, newList, start, result)
				newList = newList[:(len(newList) - 1)]
			}
		}
	}
}

/*
* includes: The function check that choice is in the list
* @param list The array is checked
* @param choice The value need to check
*/
func includes(list []int, choice int) bool {
	for _, value := range list {
		if value == choice {
			return true
		}
	}
	return false
}

func copy(dst, src []int) []int {
	for _, v := range src {
		dst = append(dst, v)
	}
	return dst
}

