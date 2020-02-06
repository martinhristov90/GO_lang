package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("Sums is %d\n", sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Printf("Sums is %d\n", sum(a...)) // when the arguments are already in a slice the function can be called
	// by placing ellipsis after the slice

}

func sum(vals ...int) int { // .. called elipsis

	var sums int

	for _, value := range vals { // Inside the function body vals ( type ...int ) is actually a slice of ints
		sums += value
	}
	return sums
}
