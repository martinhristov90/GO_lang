package main

import (
	"crypto/sha256"
	"fmt"
)

// Slices, in this case []byte slices are passed by referrence
func capitalize(str []byte) []byte {
	str[0] = 'X'
	return str
}

func main() {

	str := []byte{'x'}
	// passing by reference
	capitalize(str)
	// change value of str by passing it to func as reference

	c1 := sha256.Sum256(str)
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("Hex-decimal value of c1 : %x\n", c1) //4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	fmt.Printf("Hex-decimal value of c2 : %x\n", c2) //4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	fmt.Printf("Is c1 equal to c2 : %v\n", c1 == c2) //true
	fmt.Printf("Type of c1 %T\n", c1) //[32]uint8

}
