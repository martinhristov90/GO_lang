package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 { // a method
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {

	var my_point ColoredPoint
	my_point.X = 1
	my_point.Y = 2

	fmt.Printf("X = %g\n", my_point.X)
	fmt.Printf("Distance = %g\n", my_point.Distance(Point{3, 4})) // Methods of Point (promoted methods) can be called with ColoredPoint
	// When compiling wrapper methods are generated, for example point.Distance() translates into :
	// func (p ColoredPoint) Distance (q Point) float64 {
	//	return p.Point.Distance(q)
	//}
	// mypoint can use methods explicitly declared, methods from ColoredPoint and methods from Point

	v := Point.Distance // Method assigned to a var

	fmt.Println(v(Point{1, 2}, Point{3, 4})) // The method Point.Distance, requires "receiver" and the point to measure the distance
}
