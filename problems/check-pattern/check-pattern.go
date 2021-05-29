package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(Pattern("abbab", "dog,cat,cat,dog,cat"))
}


func Pattern(pattern string, input string ) bool {
	patternChars := strings.Split(pattern, "")
	inputElements := strings.Split(input, ",")

	if len(patternChars) != len(inputElements) {
		return false
	}

	patternMap := make(map[string][]int)
	inputMap := make(map[string][]int)

	patternInputMap := make(map[string]string)
	inputPatternMap := make(map[string]string)

	for i := 0; i < len(patternChars); i++ {
		_, ok := patternInputMap[patternChars[i]]
		if !ok {
			patternInputMap[patternChars[i]] = inputElements[i]
		}
		_, ok1 := inputPatternMap[inputElements[i]]

		if !ok1 {
			inputPatternMap[inputElements[i]] = patternChars[i]
		}
	}

	if len(patternInputMap) != len(inputPatternMap) {
		return false
	}

	for i, char := range patternChars  {
		_, ok := patternMap[char]
		if !ok {
			var arr []int
			arr = append(arr, i)
			patternMap[char] = arr
		} else {
			patternMap[char] = append(patternMap[char], i)
		}
	}

	for i, char := range inputElements  {
		_, ok := inputMap[char]
		if !ok {
			var arr []int
			arr = append(arr, i)
			inputMap[char] = arr
		} else {
			inputMap[char] = append(inputMap[char], i)
		}
	}

	for k, v := range patternInputMap  {
		if !compareSlice(patternMap[k], inputMap[v]){
			return false
		}
	}

	return true
}

func compareSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++  {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

