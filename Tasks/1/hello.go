package main

import "fmt"

//imported the package fmt

func checkMultiplyAbort(m, n [][]int) [][]int {
	modFactor := 1000000009                                                // taking all the possible case at very first level then dividing the level of the code
	w := (m[0][0] * n[0][0] % modFactor) + (m[0][1] * n[1][0] % modFactor) // tracing inside the application with dfs one then one and same way exiting the pair match
	x := (m[0][0] * n[0][1] % modFactor) + (m[0][1] * n[1][1] % modFactor) // moving cursor for colum wise indexing as it will stop at intial statement as we are passing max of 3 index
	y := (m[1][0] * n[0][0] % modFactor) + (m[1][1] * n[1][0] % modFactor)
	z := (m[1][0] * n[0][1] % modFactor) + (m[1][1] * n[1][1] % modFactor)

	return [][]int{{w % modFactor, x % modFactor}, {y % modFactor, z % modFactor}}
}

func getFibo(num int) [][]int {
	m := [][]int{{1, 1}, {1, 0}} // 2d array for for the base case for 2 index as we are passing a single digit . So it will check only consequitive cases only
	if num <= 1 {
		return m
	}
	n := getFibo(num / 2)
	secondaryOutput := checkMultiplyAbort(n, n)

	if num%2 != 0 {
		secondaryOutput = checkMultiplyAbort(secondaryOutput, m)
	}

	return secondaryOutput
}

func doOperation(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	result := getFibo(n - 1)
	// retreving the initial result for the standalone index as both of the case are reduced down to 0th index limit
	return result[0][0]
}

func main() {
	// var myslice = []int64{1, 2, 3}
	// myslice = append(myslice, 9, 5, 6)
	// fmt.Println("Slice Is: ", myslice[0:4])
	// fmt.Println(myslice)
	// fmt.Print(cap(myslice))

	/*Map Type*/
	// var mymap = make(map[int]string)
	// mymap[2] = "Arijit"
	// fmt.Print(mymap)

	fmt.Println(doOperation(9))
}
