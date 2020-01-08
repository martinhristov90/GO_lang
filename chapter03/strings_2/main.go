package main

import "fmt"

func main() {

	s := "hello world" //escape chars can be used here
	s1 := `hello world` // raw string, all /r are deleted so it is the same on all OS 
	s2 := `Multiple line:
	- one
	- two
	`
	fmt.Println(s,s1,s2)
}