package main
// Shows how many times Add() was called with particular []string
// Useful for non-compatible types as keys of the map
import (
	"fmt"
)

//m := map[string]int{} not working outside funcs
var m = make(map[string]int)

func main() {
	
	str := []string{"Marti"}
	Add(str)
	Add(str)
	Add(str)
	Add(str)
	fmt.Printf("%v\n",str)
	fmt.Printf("%v\n",m[fmt.Sprintf("%q",str)]) // same as below
	fmt.Printf("%v\n",m[k(str)]) // same as above, should print 4
}

func k(list []string) string {
	return fmt.Sprintf("%q",list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count (list []string) int {
	return m[k(list)]
}