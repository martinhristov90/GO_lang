package main

import (
	"fmt"
)

func main() {

	fmt.Println(aTimesb(2, 3))
}

func aTimesb(a, b int) (c int) {

	if a == 0 && b == 0 {
		return
	}
	c = a * b
	return // bare return with named return arguments
}
