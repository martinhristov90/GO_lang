package main

import (
	"fmt"
	"time"
)

func main() {

	go count("marti")
	go count("kremi")

	fmt.Scanf("fix") // Wrong wait to wait for goroutines to finish
}

func count(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
		time.Sleep(1 * time.Second)
	}
}
