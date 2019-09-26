package main

import (
	"fmt"
)

func main() {
	x := 1 // equal to "var x int = 1 "
	p := &x

	fmt.Println("Value to which P points to is", *p)
	fmt.Println("Address to which P points to is", p)

	*p = 2
	fmt.Println("x equal to :", x)
}
