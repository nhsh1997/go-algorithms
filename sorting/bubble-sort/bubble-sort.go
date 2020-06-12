package bubble_sort

import "fmt"

func Sort(list []int, ascending bool) {
	var N int = len(list)
	var count int

	for i := 0; i < N; i++ {
		pass(list, ascending)
		count++
	}
	fmt.Println("Number of iterations: ", count)
}