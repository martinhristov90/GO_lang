package main

import (
	"fmt"
)

var temp byte = 0xFF

func main() {
	
	temp&= 0x0F
	
	fmt.Printf("printing temp var in dec %d",temp)
	
}
