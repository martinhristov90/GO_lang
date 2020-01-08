package main

import "fmt"

func main() {

	s := "hello, world"

	fmt.Printf("Lenght of the string is : %d\n",len(s))
	// zeroth byte
	fmt.Println(s[0])
	//substring
	fmt.Printf("substring : %s \n",s[0:5])
	// the whole string
	fmt.Printf("the whole string : %s \n",s[:])
	// concatenating
	// starts from the comma, "hello" world is 5 bytes as ASCII = UTF-8, not true for Unicode points that take 2 bytes
	fmt.Println("goodbuy" + s[5:])
	// appending to the same string
	s += " :)" //same as s = s + " :)"
	fmt.Println(s)
	// This is not allowed s[0] = "L", results in error
	// Strings in Go are like Python strings, references.
}