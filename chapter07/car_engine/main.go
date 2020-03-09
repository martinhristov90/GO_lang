package main

import (
	"fmt"
)

type engine struct {
	model           string
	capacity, power int
}

type car struct {
	engine
	maker, model string
}

func (eng *engine) printModel() {
	fmt.Printf("The model of the engine is %s\n", eng.model)
}

func (car *car) printModel() {
	fmt.Printf("The model of the car is %s\n", car.model)

}

type Imodel interface {
	printModel()
}

func myModel(I Imodel) {
	I.printModel()
}

func main() {

	var my_engine = engine{"b18c4", 1800, 169}

	var my_car = car{
		engine: my_engine,
		maker:  "Honda",
		model:  "Civic",
	}

	myModel(&my_car)    // it requires a pointer to a object of car, because of func (car *car)
	myModel(&my_engine) // uses printModel() of engine type

	var w Imodel

	fmt.Printf("%t, %v , %T \n",w == nil,w,w)

	w = &my_engine

	fmt.Printf("%t, %v , %T \n",w == nil,w,w)

	w = &my_car

	fmt.Printf("%t, %v , %T \n",w == nil,w,w)

}
