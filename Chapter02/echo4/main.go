package main

import (
	"fmt"
	"flag"
	"strings"
)
// n and sep are pointers
var n = flag.Bool("n",false,"omit newline char the end")
var sep = flag.String("s", " ","separator")

func main() {
	flag.Parse()
	// to get the value of sep, *sep needs to be used "sep" is a pointer.
	fmt.Print(strings.Join(flag.Args(), *sep ))
	if !*n {
		//if n has no value, print empty line
		fmt.Println()
	}
}