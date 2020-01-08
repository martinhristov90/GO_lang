package main

import "fmt"

func main() {

	var a bool = true

	b := btoi(a)

	fmt.Println(b)

}

func btoi(b bool) int {

	if b {
		return 1
	}
	return 0
}
