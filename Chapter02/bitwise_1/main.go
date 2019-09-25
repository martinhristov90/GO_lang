package main

import (
	"fmt"
)

var test uint8 = 0xAC

func main() {
	fmt.Printf("Variable test has a type of %T\n", test)
	fmt.Printf("Binary content of variable test is : %08b\n", test)
	fmt.Printf("Decimal representation of variable test is : %d\n", test)

	// Appling AND mask
	test = test & 0xF0
	// this can be written short hand test &= 0xF0

	fmt.Printf("Binary content of variable test is : %08b after appling 0xF0 AND mask\n", test)
	fmt.Printf("Decimal representation of variable test is : %d after appling 0xF0 AND mask\n", test)

}
