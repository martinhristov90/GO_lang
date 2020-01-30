package main

// Passing struct with pointer

import (
	"fmt"
)

type Point struct {
	x, y int
}

func main() {
	var dot1 = Point{2,3} // long struct declaration
	dot := &Point{2, 3}
	fmt.Printf("Double point %v\n", double(dot))
	fmt.Printf("Dot point %v\n", *dot) // New Point is returned, dot remains untouched
	fmt.Printf("Compare two structs of same type %t\n", *dot == dot1) // one as pointer, other as value
}

func double(point *Point) Point {

	return Point{point.x * 2, point.y * 2}
}
