package main

import "fmt"

func main(){

	arr := []int{1, 100, 9, 5, 8, 10, 15, 2}
	arr = mergeSort(arr)
	fmt.Println(arr)
}

func merge(L, R []int) []int {
	result := make([]int, len(L) + len(R))

	i := 0
	for len(L) > 0 && len(R) > 0  {
		if L[0] < R[0] {
			result[i] = L[0]
			L = L[1:]
		} else {
			result[i] = R[0]
			R = R[1:]
		}
		i++
	}
	for j := 0; j < len(L) ; j++  {
		result[i] = L[j]
		i++
	}

	for j := 0; j < len(R) ; j++ {
		result[i] = R[j]
		i++
	}

	return result
}


func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr)/2
	var L, R []int

	L = copy(arr[:mid], L)
	R = copy(arr[mid:], R)

	return merge(mergeSort(L), mergeSort(R))
}

func copy(dst, src []int) []int {
	for _, v := range src {
		dst = append(dst, v)
	}
	return dst
}