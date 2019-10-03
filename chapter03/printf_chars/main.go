package main

import (
	"fmt"
)

func main() {
	ascii := 'a'
	// c to print a char
	// q for quoting
	// [1] to use the first verb
	fmt.Printf("%d %[1]c %[1]q\n",ascii)
}