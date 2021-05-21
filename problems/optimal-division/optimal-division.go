package main

import (
	"fmt"
	"math"
)

func main()  {
	nums := []int{1000,100,10,2}
	fmt.Println(optimalDivision(nums))
}

type T struct {
	minVal int `json:"min_val"`
	maxVal int `json:"max_val"`
	minStr string `json:"min_str"`
	maxStr string `json:"max_str"`
}
func optimalDivision(nums []int) string {
	t := optimal(nums, 0, len(nums) - 1, "")
	return t.maxStr
}

func optimal(nums []int, start int, end int, response string) T {
	var t T
	if start == end {
		t.maxVal = nums[start]
		t.minVal = nums[start]
		t.minStr = fmt.Sprintf("%d", nums[start])
		t.maxStr = fmt.Sprintf("%d", nums[start])
		return t
	}

	t.minVal = math.MaxInt64
	t.maxVal = math.MinInt64

	for i := start; i < end ; i++  {
		left := optimal(nums, start, i , "")
		right := optimal(nums, i + 1, end, "")
		if t.minVal > left.minVal / right.maxVal {
			t.maxVal = left.minVal / right.maxVal
			if i + 1 != end {
				t.minStr = left.minStr + "/" + "(" + right.maxStr + ")"
			} else {
				t.minStr = left.minStr + "/" + "" + right.maxStr + ""
			}
		}

		if t.maxVal < left.maxVal / right.minVal {
			t.maxVal = left.maxVal / right.minVal
			if i + 1 != end {
				t.maxStr = left.maxStr + "/" + "(" + right.minStr + ")"
			} else {
				t.maxStr = left.maxStr + "/" + "" + right.minStr + ""
			}
		}
	}

	return t
}

