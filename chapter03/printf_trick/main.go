package main

import (
	"fmt"
)

func main() {
	o := 0666
	x := int64(0xdeadbeef)

	fmt.Printf("%d %[1]o %#[1]o\n",o)
	fmt.Printf("%d %#[1]x %#[1]X\n",x)
}