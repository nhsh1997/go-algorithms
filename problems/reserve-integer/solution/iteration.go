package reserve_integer_solution

//#Problem: https://leetcode.com/problems/reverse-integer/solution/

func reverse(x int) int {
	result := 0
	current := x
	flag := false
	if x < 0 {
		flag = true
		current = -x
	}
	for current > 0 {
		result *= 10
		result += current % 10
		current /= 10
	}

	if !isSignedInteger(result) ||!isSignedInteger(-result){
		return 0
	}

	if flag {
		return -result
	}
	return result
}

func isSignedInteger(n int) bool {
	const MaxUint = ^uint32(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	if n > MaxInt || n < MinInt {
		return false
	}
	return true
}
