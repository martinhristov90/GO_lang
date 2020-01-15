package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Printf("The type of y is %T\n", y)
	fmt.Printf("The value of y is %v\n", y)

	fmt.Printf("Changing types...")

	if output, err := strconv.Atoi(y); err == nil {
		fmt.Printf("The type of y now is %T\n", output)
	}

}
