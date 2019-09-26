package main

import (
	"fmt"
)

func f() {}

func main() {
	f := 1
	fmt.Println(f)
	// the local f var is going to shadown the function f, when reference a name, the look up starts from inside out.
}
