package main

import (
	"fmt"
)
func incr(p *int) int {
	*p++
	return *p
}

func main(){
	v := 1 
	// calling incr with the address of v - &v
	// indirect referencing
	incr(&v)
	fmt.Println(v)
}