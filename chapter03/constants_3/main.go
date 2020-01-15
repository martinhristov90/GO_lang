package main
// testing import to local namespace
import (. "fmt")

func main() {

	type Weekday int

	const (
		Monday Weekday = iota
		Tuesday = iota
		Wednesday = iota
	)

	Printf("%v\n",Wednesday)

}