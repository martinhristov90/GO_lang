package main

import (
	"fmt"
	"strings"
)

func main() {

	// Modifing indexes, starting from 1
	months := [...]string{
	1: "Jan", 2: "Feb", 3: "Mar", 
	4:"Apr",5: "May", 6: "Jun", 
	7: "Jul", 8: "Aug", 9: "Sep", 
	10: "Oct", 11: "Nov", 12: "Dec"}

	fmt.Println(months[0]) // Prints empty string
	fmt.Println(months[1]) // Jan
	fmt.Println(cap(months)) // 13 , it starts from 0 though 12, even i use 12
	fmt.Println(len(months)) // 13 , same
	fmt.Println(months[3:4]) // Same as Python, last index non-inclusive

	fmt.Println(strings.Repeat("-",20))

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(Q2) // [Apr May Jun]
	fmt.Println(summer) // [Jun Jul Aug]
	
	// Check for common elements in two slices
	for _, elemQ2 := range Q2 {
		for _, elemSummer := range summer {
			if elemQ2 == elemSummer {
				fmt.Println("There is overlapping month of Q2 and the summer:",elemQ2)
			}
		}
	}

	endlessSummer := summer[:5] // The zeroth element of slice summer is "Jun" , starts the counting from there
	fmt.Println(endlessSummer) // [Jun Jul Aug Sep Oct]
}