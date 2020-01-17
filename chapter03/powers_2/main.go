package main

import (
	"fmt"
)
// Shift by 1 
const (
	firstBit = 1 << (iota)
	secondBit 
	thirdBit
	forthBit
	fifthBit
	sixthBit
	seventhBit
	eighthBit
)

func main() {
	// Binary representation with padding of 0s and decimal representation
	fmt.Printf("Binary representation : %08b Decimal : %[1]d\n",seventhBit)

}