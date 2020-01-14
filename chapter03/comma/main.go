package main

import (
	"fmt"
)
// Gets a string as 12345 and return a string like 12,345

func comma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {

fmt.Println(comma("1234567890"))

}