package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 { // a method
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// It is valid, but not recommended to have regular functions and methods with same name
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point // Slice of Point type

func (path Path) Distance() float64 { // this is Path.Distance even it uses Point.Distance, they are completely separate
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}

	return sum
}

func main() {

	var my_point Point = Point{2, 3}

	fmt.Printf("%g is the distance\n", my_point.Distance(Point{4, 5}))
	fmt.Printf("%g is the distance\n", Distance(my_point, Point{4, 5})) // Both are the same, method vs regular func

	perim := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}

	fmt.Println(perim.Distance()) // Calling Distance method of type Path

}
