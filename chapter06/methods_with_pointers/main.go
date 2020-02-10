package main

import(
	"fmt"
)

type Point struct {
	X,Y int
}

func (p *Point) Cordinates () string {
	// Prints the cordinates of the point

	return fmt.Sprintf("X cord : %d, Y cord: %d",p.X,p.Y)
}

func main(){

	tochka := Point{1,2}
	ptrTochka := &tochka
	fmt.Println((*ptrTochka).Cordinates())
	fmt.Println(ptrTochka.Cordinates()) // ^ both are the same, the compiler does &ptrTocka implicitly


}

