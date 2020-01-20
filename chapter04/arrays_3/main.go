package main

import "fmt"

func translate_second(str_arr []string) {
	//slices are passed by reference
	str_arr[1] = "две"

}

func main() {

	txt := []string{"one", "two", "three"}

	translate_second(txt)

	fmt.Printf("%v\n", txt)

}
