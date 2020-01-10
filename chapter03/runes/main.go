package main

import (
	"fmt"
	//"unicode/utf8"
)

func main() {

	str := "Hello, 世界"

	for i,r := range(str) {
		fmt.Printf("%d,%q\n",i,r)
	}

}