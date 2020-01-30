package main

// Exercise from https://medium.com/rungo/structures-in-go-76377cc106a2

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	salary              int
	fullTime            bool
}

func main() {

	var ross Employee

	fmt.Printf("%v\n", ross) // Printing zero value for the struct above

	//marti := Employee{
	//	firstName: "Martin",
	//	lastName: "Hristov",
	//	salary: 3000,
	//	fullTime: true,
	//} // Initializing a struct
	marti := Employee{"Martin", "Hristov", 5000, false} //same as above, no field definition
	fmt.Printf("%v\n", marti)

	// Anonymous structs

	monica := struct {
		firstName, lastName string
	}{
		firstName: "Monica",
		lastName:  "Hristova",
	}

	fmt.Printf("%T\n", monica)

	// Pointer to a struct

	georgi := &Employee{
		firstName: "Georgi",
		lastName:  "Berchev",
		salary:    700,
		fullTime:  false,
	}

	fmt.Printf("Georgi's salary is : %v\n", (*georgi).salary)
	// We are using parenthesis around the pointer dereferencing syntax in the program above so that compiler doesn’t get confused between (*ross).firstName and *(ross.firstName).
	fmt.Printf("Georgi's salary is : %v\n", georgi.salary) // Implicit dereference, same as above

	nikolay := &Employee{firstName: "Nikolay", salary: 550} // Setting not all fields of a struct, other will take their default values, for example string - "", int - 0

	fmt.Printf("Nikolay's salary %v\n", (*nikolay).salary)

}
