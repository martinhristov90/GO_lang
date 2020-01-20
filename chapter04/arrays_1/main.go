package main

import "fmt"

func main() {

	// [...] means that array length is determined by the initial elements, in this case 3
	arr := [...]int{1,2,3}

	// print [3]int
	fmt.Printf("%T\n",arr)


}