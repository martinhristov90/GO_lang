package main

// interesting example using pointer for calculating area of square
import (
	"fmt"
)

func square(x *float64) float64 {
	*x = *x * *x
	return *x
}
func main() {
	x := 1.5
	fmt.Println("test", square(&x))
	// should print 2.25
}
