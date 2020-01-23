package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[0:2] // 1,2

	fmt.Println(slice2) //[1 2]

	slice2[0] = 300 // slice2 modifies the original slice1, not a copy

	fmt.Println(slice2) //[300 2]
	fmt.Println(slice1) //[300 2 3 4 5]

	//[1 2]
	//[300 2]
	//[300 2 3 4 5]

}
