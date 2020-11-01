package main

import (
	"fmt"
	"time"
)

var cache = map[int]int{
	0: 0,
	1: 1,
	2: 1,
}
var latestNum int = 1

func fib(n int) int {
	if n >= 0 {
		i := 3 // last value inside the cache
		if n == 0 {
			return 0
		} else if n <= 2 {
			return 1
		} else {
		LOOP:
			if i <= n {
				cache[i] = cache[i-1] + cache[i-2]
				i++
				goto LOOP
			}
			return cache[n]
		}
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
