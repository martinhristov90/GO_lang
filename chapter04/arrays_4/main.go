package main

import "fmt"

func main() {

	a := [2]int{1,2}
	b := [...]int{1,2}
	c := [2]int{1,3}

	fmt.Println( a==b, a==c, b==c)

	// d := [4]int{1,2,3,4}

	// even first and the second element are the same the types [2]int and [4]int are different, cannot compare different types
	// fmt.Println(a == d)


}