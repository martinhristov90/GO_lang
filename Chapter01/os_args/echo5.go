// echo with range func
// echo program that prints the arguments and their numbers
package main

import (
	"fmt"
	"os"
)

func main() {
	for arg_num, arg := range os.Args[1:] {
		fmt.Print("Arg num: ", arg_num+1, " - ")
		fmt.Print("Arg content: ", arg, "\n")
	}
}
