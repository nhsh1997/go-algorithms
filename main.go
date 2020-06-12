package main

import "fmt"
import "github.com/nhsh1997/go-algorithms/sorting/bubble-sort"

func main()  {
	var indices [2]int
	var given_nums []int = []int{11, 2, 7, 15}
	var target int = 17
	N := len(given_nums)
	fmt.Println("Original list: ", given_nums)

	bubble_sort.Sort(given_nums, true)


	fmt.Println("Sorted list: ", given_nums)

	for i := 0; i < N - 1; i ++  {
		if given_nums[i] <= target {
			for j := i + 1; j < N ; j++ {
				if (given_nums[i] + given_nums[j]) == target {
					indices[0], indices[1] = i, j
					break
				}
			}
		}
	}

	fmt.Println("Indices: ", indices)
}
