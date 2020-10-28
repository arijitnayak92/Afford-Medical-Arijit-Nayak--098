package main

import "fmt"

//imported the package fmt

func getFibo(n int) int {
	if n <= 1 {
		return n
	} else {
		return getFibo(n-1) + getFibo(n-2)
	}
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

	fmt.Println(getFibo(9))
}
