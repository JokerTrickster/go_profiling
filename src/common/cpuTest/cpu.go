package cpuTest

func IncreaseInt() {
	i := 0
	for {
		i = increase1000(i)
		i = increase2000(i)
	}
}

func IncreaseIntGoroutine() {
	go func() {
		i := 0
		for {
			i = increase1000(i)
			i = increase2000(i)
		}
	}()
}

func increase1000(n int) int {
	for n := 0; n < 1000; n++ {
		n = n + 1
	}
	return n
}

func increase2000(n int) int {
	for n := 0; n < 1000; n++ {
		n = n + 1
	}
	return n
}

func FibRecursive(n int) uint64 {
    // Function Creation to accept integer till which the Fibonacci series runs
    if n == 0 {
        return 0
        // Base case for the recursive call
    } else if n == 1 {
        return 1
        // Base case for the recursive call
    } else {
        return FibRecursive(n-1) + FibRecursive(n-2)
        // Recursive call for finding the Fibonacci number
    }
}
