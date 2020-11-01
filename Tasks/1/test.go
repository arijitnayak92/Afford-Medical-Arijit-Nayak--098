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
var lastNum int = 3

func fib(n int) int {
	if n >= 0 {
		if _, ok := cache[n]; ok {
			return cache[n]
		} else {
			for i := lastNum; i <= n; i++ {
				cache[i] = cache[i-1] + cache[i-2]
			}
			lastNum = n
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
