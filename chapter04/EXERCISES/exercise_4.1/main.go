//“Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256 hashes. (See PopCount from Section 2.6.2.)”

// A.K.A finding hamming distance of c1 and c2

package main

import (
	"crypto/sha256"
	"fmt"
)

func hamming_distance(c1, c2 [32]byte) int {

	if len(c1) != len(c2) {
		fmt.Printf("Strings of both hashes should be equal, e.g 32 bytes\n")
	}

	diff_bits := 0

	bit_mask := [...]byte{1, 2, 4, 8, 16, 32, 64, 128}
	// 0000 0000 first index [0]
	// 0000 0010
	// 0000 0100
	// 0000 1000
	// 0001 0000
	// 0010 0000
	// 0100 0000
	// 1000 0000 last index [7]

	// comparing byte by byte
	for i := 0; i < len(c1); i++ {
		// getting one byte from c1 and one byte from c2
		c1_byte := c1[i]
		c2_byte := c2[i]
		// Comparing bits of both bytes using mask
		for j := 0; j < 8; j++ {
			// Using new mask for each iteration and comparing bitwise, depending on j 
			if (c1_byte & bit_mask[j]) != (c2_byte & bit_mask[j]) {
				diff_bits++
			}

		}

	}

	return diff_bits

}

func main() {

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("Hex-decimal value of c1 : %x\n", c1) //2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	fmt.Printf("Hex-decimal value of c2 : %x\n", c2) //4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	fmt.Printf("Is c1 equal to c2 : %v\n", c1 == c2) //false
	fmt.Printf("Type of c1 %T\n", c1)                //[32]uint8

	fmt.Printf("Hamming distance of two Sha256 sums %v\n", hamming_distance(c1, c2))

}



