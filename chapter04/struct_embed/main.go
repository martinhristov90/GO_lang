package main

import (
	"fmt"
)

// Embedded struct
type engine struct {
	engine_type string
	capacity    int
}

type car struct {
	make        string
	model       string
	engine_info engine
}

func main() {

	my_car := &car{
		make:  "Ford",
		model: "Escort",
		engine_info: engine{ // important for composite struct literals
			engine_type: "Gasoline",
			capacity:    1800,
		},
	}

	display_make_engine_type(my_car)

}

// Using fields from embedded struct
func display_make_engine_type(car *car) {
	fmt.Printf("Make : %v\nEngine_type : %v\n", car.make, (*car).engine_info.engine_type)
	// car.make is the same as (*car).make, go does this for us
}
