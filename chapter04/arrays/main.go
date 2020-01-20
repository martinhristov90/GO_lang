// Arrays are fixed size and homogeneous (have same type of elements, for example all elements of particular array are int)

// homogeneous - same type of elements
// heterogeneous - different types of elements gathered together

package main

import "fmt"

func main() {

	var a [3]int // Array of three elements of type int

	fmt.Println(a[0]) // Printing the first element

	fmt.Println(a[len(a)-1]) // Printing last element

	// Lets print the indexes and their values
	for i,v := range a {
		fmt.Printf("Index %d - Value %v\n",i,v)
	}

	// Print only the values, discard the indexes

	for _,v := range a {
		fmt.Printf("Value : %v\n",v)
	}
}
