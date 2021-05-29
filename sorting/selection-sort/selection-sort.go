package main

import "fmt"

func main()  {
	arr := []int{1, 2, 8, 8, 0, 2, 1, 0}
	arr = selectionSort(arr)
	fmt.Println(arr)
}

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n - 1 ; i++  {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min]{
				min = j
			}
		}
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
	return arr
}

