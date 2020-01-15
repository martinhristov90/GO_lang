package main

import (
	"fmt"
	"time"
)

func main() {

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute

	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
	fmt.Printf("%T %[1]v\n", 30*time.Minute.Hours())
	fmt.Printf("%T %[1]v\n",time.Now())
	fmt.Printf("%T %[1]v\n",time.Now().Unix())


	var x time.Duration = 10
	// time.Duration and time.Time are different types for handling time and duration
	fmt.Printf("%T %[1]v\n",x)
}
