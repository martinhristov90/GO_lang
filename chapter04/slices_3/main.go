package main

import "fmt"

func main() {

	a := []string{"Marti", "Georgi", "Nikolay"}

	b := a[1:]

	fmt.Println(b)
	// b[0] is a[1] changeing the same underlying array
	b[0] = "Georgiman"

	fmt.Println(a)
}
