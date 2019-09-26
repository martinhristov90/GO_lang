package main

// Determines if a number is odd or even using AND bitwise operator
import (
	"fmt"
	"math/rand"
)

func main() {

	//	rand.Seed(1234)

	var num int = rand.Int()
	fmt.Printf("Randomly generated number is : %d\n", num)
	num &= 1 // can be shorten to num&1

	if num == 1 {
		fmt.Println("The number is odd")
	} else {
		fmt.Println("The number is even")
	}

}
