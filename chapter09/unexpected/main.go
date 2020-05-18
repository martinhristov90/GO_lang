package main

import (
	"fmt"
)

func main() {
	var x []int

	go func() {
		x = make([]int, 100000)
	}()

	go func() {
		x = make([]int, 10)
	}()


	x[99999] = 1 // unxpected behaviour

	fmt.Println(x[99999])

}
