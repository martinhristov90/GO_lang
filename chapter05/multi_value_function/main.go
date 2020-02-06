package main

import (
	"fmt"
)

func main() {
	// initalizing
	a, b := 2, 3
	// Calling the function by reference
	double_log(&a, &b)
}

func double(a, b int) (int, int) {
	// doubles a and b
	return a * 2, b * 2

}

func double_log(a *int, b *int) (int, int) {
	fmt.Println(*a, *b)   // print values of a,b
	return double(*a, *b) // Function can be used as return statement, the return statement of double and double_log should match
}
