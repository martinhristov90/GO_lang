package main

import "fmt"

func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}


func main() {

	a := [...]int{1,2,3,4,5}

	reverse(a[:]) // Slice needs to be passed to the reverse func, not [5]int array

	fmt.Println(a)

	// Unline arrays of same size, slices are not comparable with ==
}