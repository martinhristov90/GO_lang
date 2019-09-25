package main

import (
	"fmt"
	// The package should be imported in GOPATH
	"tempconv"
)

var temp tempconv.Celsius

var outsideTemp tempconv.Celsius = 20

var tempK tempconv.Kelvin = 200

func main() {
	fmt.Println(temp)
	fmt.Println(tempconv.CToF(outsideTemp))
	fmt.Printf("Температурата на замръзване е %g°C\n", tempconv.AbsolouteZeroC)
	// Exercise 2.1
	fmt.Printf("Type of tempK is %T\n", tempK)
	fmt.Printf("%g Kelvins is equal %g Celsius\n", tempK, tempconv.KToC(tempK))
}
