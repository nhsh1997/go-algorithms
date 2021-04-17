package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(isHappy(3))
}

func isHappy(n int) bool {
	if n < 1 || n > math.MaxInt32 || n == 145 {
		return false
	}

	var digits []int

	if n >= 10 {
		for n > 0 {
			digit := n % 10
			digits = append(digits, digit)
			n = n / 10
		}
	} else {
		digits = append(digits, n)
		digits = append(digits, 0)
	}

	total := 0

	for _, digit := range digits {
		total += digit * digit
	}

	if total != 1 {
		return isHappy(total)
	}

	return true
}
