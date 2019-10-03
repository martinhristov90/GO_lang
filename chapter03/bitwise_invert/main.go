package main

import (
	"fmt"
)

var temp uint8 = 128

func main() {
	fmt.Printf("%08b",^temp)
}
