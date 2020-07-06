package climb_stairs

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairs(n - 1) + climbStairs(n - 2)
}


func climbStairs1(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	first := 1
	second := 2
	for i := 3; i <= n; i ++ {
		third := first + second
		first = second
		second = third
	}
	return second
}