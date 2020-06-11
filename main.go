package main

import "fmt"
import "sorting/bubble-sort/bubble-sort.go"

func main()  {
	var numbers []int = []int{33, 91, 76, 8, 22}
	var ascending bool = true

	fmt.Println("Original list: ", numbers)
	bubble_sort.sort(numbers, ascending)
	fmt.Println("Sorted list: ", numbers)
}
