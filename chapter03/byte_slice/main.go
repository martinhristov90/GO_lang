package main

import "fmt"

func main() {

	var s string = "abc"
	// Convert string to slice of bytes
	b := []byte(s)
	// Change one byte
	b[1] = 'd'
	// Conv back to string
	s2  := string(b)
	// Print it
	fmt.Printf("Value of s: %v\n",s)
	// s2 is copy of s, not actual s 
	fmt.Printf("Value of s2: %v\n",s2)

}