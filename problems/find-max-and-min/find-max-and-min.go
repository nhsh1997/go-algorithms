package main

import "fmt"

func main()  {
	fmt.Println(FindMaxAndMin([]int{1, 2, 5, 9, 1, 4, 10}))
}


func FindMaxAndMin(nums []int) (int, int) {
	if len(nums) < 0 {
		return -1, -1
	}
	var min, max int
	min = nums[0]
	max = nums[0]
	for i := 0; i < len(nums); i++  {
		if nums[i] < min {
			min = nums[i]
		}  else if nums[i] > max {
			max = nums[i]
		}
	}
	return min, max
}