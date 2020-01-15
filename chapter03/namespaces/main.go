package main

import (
	. "fmt"
)
// going to crash badly, overlapping namespace
func Println(a string) string {
	return "hello"
}

func main() {

	Println("HELLO WORLD")

}