package main

import "fmt"

func main() {
	// Createing pointer, p defaults to <nil>, it can only point to addresess that contain data of int type
	var p *int
	// x has value of 2
	x := 2
	// p can only hold address of int type data, &x returns the address in memory where "2" is stored, x is just a name that points to this area in memory.
	p = &x
	// *p changes the data at the address to which it points to, now p has value (for exampe) 0xc0000a4008 and x (&x address) points to the same address. Important: addresses of x and p are different.
	*p = 3
	//printing x
	fmt.Println(x)

}
