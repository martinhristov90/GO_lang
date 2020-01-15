package main

import "fmt"

func main() {
	// when constants are used in a group, types and values might be omitted, it gets inherited from the first const in the list.
	const (
		a = 1
		b
		c
	)

	fmt.Printf("Type of a is %T value is %[1]v, type of b is %T value is %[2]v, type of c is %T value is %[3]v\n", a, b, c)

}
