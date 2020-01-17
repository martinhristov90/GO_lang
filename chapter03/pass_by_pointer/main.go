package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func changeName(p *Person) {
	p.firstName = "Scott"
	fmt.Printf("The address of p struct, %p\n",p)

	
}

func main() {

	human := Person{
		firstName: "Martin",
		lastName:  "Hristov",
	}

	changeName(&human)
	
	fmt.Printf("The address of human struct, %p\n",&human)

	fmt.Println(human)

}
