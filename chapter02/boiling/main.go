//prints the boiling point of the water

package main

import (
	"fmt"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling point = %gF degrees or %gC degrees\n", f, c)
}
