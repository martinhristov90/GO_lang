package main

import (
	"flag"
	"fmt"
)

type Bar float64
type PSI float64

func PSIToBar(pressure_psi PSI) Bar {
	return Bar(pressure_psi * 0.0689475728)
}

func (str Bar) String() string {
	return fmt.Sprintf("%g Ba(Bar)", str)
}

type BarFlag struct {
	Bar
}

func (press *BarFlag) Set(s string) error {

	var unit string
	var value int
	//fmt.Printf("%q\n",s)

	_, err := fmt.Sscanf(s, "%d%s", &value, &unit) // Why it does not work with %f for floating point with 45PSI ?
	if err != nil {
		return fmt.Errorf("Error while parsing arguments %v\n", err)
	}

	//fmt.Println("Value",value)
	//fmt.Println("Unit:",unit)

	switch unit {
	case "Ba", "Bar":
		press.Bar = Bar(value)
		return nil
	case "PSI", "psi":
		press.Bar = PSIToBar(PSI(value))
		return nil
	}

	return fmt.Errorf("Cannot handle temp %q ", s)
}

func PressureFlag(name string, value Bar, usage string) *Bar {
	f := BarFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Bar
}

var pressure = PressureFlag("pressure", 20.0, "the pressure")

func main() {

	flag.Parse()
	fmt.Println(*pressure)

}
