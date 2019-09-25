package main

import (
	"fmt"
)

func main() {

	if f := 0; true {
		// both are true, so we go here
		fmt.Printf("True %d\n", f)
	}

	// Trying to use f outside of if stanza, is going to result in an error.
	// Error
	fmt.Printf("Printing f %d\n", f)
}
