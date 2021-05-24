package main

import "fmt"

func main()  {

	fmt.Println(maxSubarray([]int{-2, -1, -3, 4, -1, 2, 1, -5, - 10, 4, 10}))

}

func maxSubarray(nums []int) int {
	subsets := findSubsets(nums)
	fmt.Println(subsets)
	maxValue := sum(subsets[0])
	maxSubset := subsets[0]
	for i := 1; i < len(subsets); i++{
		sum := sum(subsets[i])
		if maxValue < sum {
			maxValue = sum
			maxSubset = subsets[i]
		}
	}
	fmt.Println(maxSubset)
	return maxValue
}

func sum(list []int) int {
	total := 0
	for _, v := range list{
		total += v
	}
	return total
}

func findSubsets(nums []int) [][]int {
	var subsets [][]int
	for i := 0; i < len(nums); i++ {
		var list []int
		list = append(list, nums[i])
		for j := i + 1; j < len(nums); j++  {
			var newList []int
			list = append(list, nums[j])
			newList = copy(list, newList)
			subsets = append(subsets, newList)
		}
	}
	return subsets
}

func copy(src, dst []int) []int{
	for _, value := range src {
		dst = append(dst, value)
	}
	return dst
}

func slice(arr []int, i int) []int  {
	// Remove the element at index i from a.
	arr[i] = arr[len(arr)-1] // Copy last element to index i.
	arr[len(arr)-1] = 0   // Erase last element (write zero value).
	arr = arr[:len(arr)-1]   // Truncate slice.

	return arr
}