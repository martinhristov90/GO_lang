package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID   int
	Name string
	DOB  time.Time
	Salary int 
}

func main() {

	var Marti Employee

	Marti.ID = 123
	//Marti.DOB = time.Date(1990,time.October, 11, 10, 00, 00 ,00 ,00 , time.LoadLocation("Europe/Sofia"))
	homeTimeZone, err := time.LoadLocation("Europe/Sofia")

	if err != nil {
		fmt.Errorf("Cannot determine location %v\n", err)
	}
	// DOB needs time.Time struct
	Marti.DOB = time.Date(1990, time.October, 11, 10, 00, 00, 00, homeTimeZone)

	fmt.Printf("%v\n", Marti.DOB)

	// Manipulating a field of a struct with a pointer
	p := &Marti.Salary

	*p = 1000

	fmt.Println(Marti.Salary)

	

}

func fire_an_employee(emp *Employee) {
	emp.Salary = 0 //fired
}