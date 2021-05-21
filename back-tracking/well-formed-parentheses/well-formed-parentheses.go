package main

import (
	"fmt"
)

func main()  {
	fmt.Println(wellFormedParentheses(2))
}

func wellFormedParentheses(target int) []string {
	var result []string

	backtrack("", target, 0, 0, &result)

	return result
}

func backtrack( current string, target, open, close int, result *[]string)  {
	// reach the target opening and closing parenthesis count
	if open == target && close == target {
		*result = append(*result, current)
	// still has room to grow
	} else {
		// able to add more opening parentheses
		if open < target {
			backtrack(fmt.Sprintf("%s(", current), target, open + 1, close, result)
		}

		if close < open {
			backtrack(fmt.Sprintf("%s)", current), target, open, close + 1, result)
		}
	}
}
