package fibonacci

func fib (n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n - 1) + fib(n - 2)
}

func fib1 (n int) int {
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
	for i := 3; i < n; i ++ {
		third := first + second
		first = second
		second = third
	}
	return second
}
