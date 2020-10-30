package main

import "fmt"

func fibO(n int) int {
	fir_asset := 0
	second_asset := 1
	if n == 0 {
		return fir_asset
	} else if n <= 2 {
		return second_asset
	}
	for i := 2; i <= n; i++ {
		second_asset = second_asset + fir_asset
		fir_asset = second_asset - fir_asset
	}
	return second_asset
}

func main() {
	num := 9
	res := fibO(num)
	fmt.Print(res)
}
