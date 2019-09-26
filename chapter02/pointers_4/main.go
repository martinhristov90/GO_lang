package main

import (
	"fmt"
)

func main() {
	a := 1
	var pa *int
	pa = &a

	fmt.Printf("pointer pa of type %T with value of %v\n", pa, pa)
}
