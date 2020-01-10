package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, 世界"

	for len(str) > 0 {
		// This functin probably is returning the size and first rune it encounters, for example if string "hello" is passed to it, it returns the info only for "h" letter(rune)
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%c %v\n", r, size)

		str = str[size:]

	}
}