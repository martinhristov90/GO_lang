package main

import (
	"fmt"
)

func main() {
	// This will allocate a random memory address to hold type int variable, it has no name, just pointer to the memory area
	p := new(int)
	// This will create a pointer named a of type *int, by default it points to `nil`
	var a *int

	*p = 2
	a = p

	fmt.Printf("the value to which the pointer P points to is: %v \n", *p)
	fmt.Printf("the address to which P points to is : %v \n", p)
	fmt.Printf("the address the pointer P itself: %v \n", &p)

	fmt.Println("-------------------------------------------")

	fmt.Printf("the value to which the pointer A points to is: %v \n", *a)
	fmt.Printf("the address to which A points to is : %v \n", a)
	fmt.Printf("the address the pointer A itself: %v \n", &a)

}
