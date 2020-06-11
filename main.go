package main

import "fmt"
import "github.com/nhsh1997/go-algorithms/sorting/bubble-sort"

func main()  {
	var numbers []int = []int{33, 91, 76, 8, 22}
	var ascending bool = true

	fmt.Println("Original list: ", numbers)
	bubble_sort.Sort(numbers, ascending)
	fmt.Println("Sorted list: ", numbers)
}
