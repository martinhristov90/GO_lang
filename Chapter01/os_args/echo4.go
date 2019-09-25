// echo with range func
// echo program that prints the command used to run it
package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("The name of the command is :", path.Base(os.Args[0]))
}
