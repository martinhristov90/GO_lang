package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("Hex-decimal value of c1 : %x\n", c1) //2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	fmt.Printf("Hex-decimal value of c2 : %x\n", c2) //4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	fmt.Printf("Is c1 equal to c2 : %v\n", c1 == c2) //false
	fmt.Printf("Type of c1 %T\n", c1) //[32]uint8

}
