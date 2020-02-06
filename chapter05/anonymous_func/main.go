package main

import (
	"fmt"
)

func main() {

	f := square()

	for i := 0; i <= 4; i++ {
		fmt.Println(f())
	}

}

func square() (func() int) { //returning an anonymous function
	var x int // this has state, it does not disappear when goes out of scope
	return func() int { 
		x++
		return x * x
	}
}
