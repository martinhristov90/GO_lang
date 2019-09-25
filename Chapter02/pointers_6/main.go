package main

import (
	"fmt"
)

var p = new(uint8)

func main() {
	
	*p = 200
	fmt.Printf("%08b",*p)
}
