package main

import (
	"fmt"
	"strings"
)

func main()  {
	input := "Hello,World"
	fmt.Println(reverse(input))
}

func reverse(input string) string {
	words := strings.Split(input, ",")
	result := ""

	for i := 0; i < len(words); i++  {
		reverseWord := reversWord(words[i])
		result = fmt.Sprintf("%s %s", result, reverseWord)
	}

	return result
}

func reversWord(input string) string {
	chars := strings.Split(input, "")
	result := ""
	n := len(chars)
	for i := n - 1; i >= 0 ; i--  {
		fmt.Println(chars[i])
		result = fmt.Sprintf("%s%s", result, chars[i])
	}
	return result
}

