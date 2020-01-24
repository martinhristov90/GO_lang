package main

import "fmt"

func main() {

	//a := make(map[int]string) // both are equicalent
	a := map[int]string{}

	a[1] = "Marti"
	a[2] = "Kremi"

	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	for name, age := range ages {
		fmt.Printf("%s is %d of age\n",name,age)
	}
}
