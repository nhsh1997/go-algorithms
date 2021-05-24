package main

import "fmt"

func main()  {
	fmt.Println(getBonuses([]int{1, 2, 3, 2, 3, 5, 1}))
	fmt.Println(getBonusesDP([]int{1, 2, 3, 2, 3, 5, 1}))
}

func getBonuses(performances []int) []int{
	bonuses := make([]int, len(performances))
	for i := 0; i < len(bonuses); i++  {
		bonuses[i] = 1
	}
	for i := 1; i < len(performances); i++  {
		if performances[i] > performances[i-1]{
			bonuses[i]++
		}
	}
	for i := len(performances) - 1 ; i >= 1; i--  {
		if performances[i-1] > performances[i]{
			bonuses[i-1]++
		}
	}
	return bonuses
}
