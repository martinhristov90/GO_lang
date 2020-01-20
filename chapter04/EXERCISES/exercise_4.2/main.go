// Example ./main -sha_type sha384 DATA
// ./main DATA // going to use sha256 by default

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	// Getting the data that should be hashed, and the type of the hash
	sha_type := flag.String("sha_type", "sha256", "Defining type of sha sum")

	//Parsing the flags
	flag.Parse()
	// Getting what is left after the flags are parsed, flag.Args() returns slice
	sha_hash := flag.Args()

	// Converting it to []byte slice, that is what Sum functions want, just the zeroth element
	sha_hash_byte := []byte(sha_hash[0])

	switch *sha_type {
	// Printing sha384 of the data provided
	case "sha384":
		fmt.Printf("Sum384 of the data is %x\n", sha512.Sum384(sha_hash_byte))
	// Printing sha512 of the data provided
	case "sha512":
		fmt.Printf("Sum512 of the data is %x\n", sha512.Sum512(sha_hash_byte))
	// By default sha256 is used
	default:
		fmt.Printf("Sum256 of the data is %x\n", sha256.Sum256(sha_hash_byte))

	}

}

// TODO:
// Implement error handling
