package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g) // print interface
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	r := rect{3, 4}
	c := circle{5}

	measure(r) // type rect
	measure(c) // type circle

}
