package main

import (
	"fmt"
)

func main() {
	fmt.Println("Value of x from function test is :", *test())
}

func test() *int {
	x := 23
	return &x
}
