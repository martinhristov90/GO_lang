package tempconv

import (
	"fmt"
)

func init(){
	fmt.Println("Initializing tempconv package")
}


type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsolouteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
	AbsolouteZeroK Kelvin = 0

)


func (c Celsius) String() string { return fmt.Sprintf("%g°C",c)}
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F",f)}
// Exercise 2.1
func (k Kelvin) String() string { return fmt.Sprintf("%g°K",k)}
