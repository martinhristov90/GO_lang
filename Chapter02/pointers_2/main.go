package main

import (
	"fmt"
)

func main() {
	x := 2
	y := &x
	fmt.Println("value of X is: ", num_incr(&x)) // calling num_incr with the address of x by using &x
	fmt.Println("value of X after num_incr: ", x)
	fmt.Println("value of Y after num_incr: ", *y)
}

func num_incr(p *int) int {
	*p = *p + 1
	return *p
}
