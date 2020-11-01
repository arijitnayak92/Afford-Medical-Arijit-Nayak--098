package main

import (
	"fmt"
	"time"
)

var cache = make(map[int]int)

func fib(n int) int {
	if n >= 0 {
		cache[0], cache[1], cache[2] = 0, 1, 1
		if n == 0 {
			return 0
		} else if n <= 2 {
			return 1
		} else if val, ok := cache[n]; ok {
			return val
		}
		// val := fib(n-1) + fib(n-2)
		// cache[n] = val
		// return val

		for i := 3; i <= n; i++ {
			cache[i] = cache[i-1] + cache[i-2]
		}

		return cache[n]
	} else {
		return -1
	}
}

func main() {
	num := 9
	start := time.Now()
	res := fib(num)
	duration := time.Since(start)
	fmt.Printf("Fibanocci of %d is = %d, execuation time %v\n", num, res, duration)
}
