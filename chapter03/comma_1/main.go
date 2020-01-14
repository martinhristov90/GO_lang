package main

import (
	"fmt"
	"bytes"
)

func comma(s string) string {
	// Declaring working buffer to add elements to
	var buf bytes.Buffer
	// Iterating over sting s , discarding values _
	for i,_ := range s {
		// Copying each byte(rune) from s to buf
		buf.WriteRune(rune(s[i]))
		// If the index of the byte(rune) deviled by 3 has remainder 2 and it is not the last element ( it will insert comma at the very end if && i != len(s) - 1 is missing), insert comma in buf
		if i % 3 == 2 && i != len(s) - 1 {
			buf.WriteRune(',')
		}
	}
	// Convert the slice of bytes to string and return it
	return buf.String()
}

func main() {

	fmt.Printf("Result: %v \n",comma("123456789123213123123"))
	// go run main.go
	// Result: 123,456,789,123,213,123,123 

}