package main

import "fmt"

func comp_slices(a, b []int) bool {
	if len(a) != len(b) {
		fmt.Println("Len of the slices should be the same")
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true

}

func main() {

	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5, 6}


	fmt.Println(comp_slices(a, b))

}
