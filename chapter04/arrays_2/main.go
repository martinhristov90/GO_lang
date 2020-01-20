package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	BGN
)

func main() {

	symbol := [...]string{USD: "$", EUR: "€", BGN: "LEV"}

	fmt.Printf("%v\n", symbol[2])
	// Same as above
	fmt.Printf("%v\n", symbol[BGN])

}
