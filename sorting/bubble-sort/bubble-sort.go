package bubble_sort

import "fmt"

func sort(list []int, ascending bool) {
	var N int = len(list)
	var count int

	for i := 1; i < N; i++ {
		pass(list, ascending)
		count++
	}
	fmt.Println("Number of iterations: ", count)
}