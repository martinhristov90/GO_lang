package main

import (
	"fmt"
)

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}

func main() {
	var buffer [30]byte

	slice := buffer[10:20]

	fmt.Println("Before: len(slice) =", len(slice)) // length of the slice before call the reduce function
	newSlice := SubtractOneFromLength(slice)        // Copy of the slice header (struct) is passed to the function, not the original, but the underlying array is the same
	fmt.Println("After:  len(slice) =", len(slice))
	fmt.Println("Slice itself ", slice)
	fmt.Println("After:  len(newSlice) =", len(newSlice))
	fmt.Println("newSlice itself ", newSlice) // it actually reduces the underlying array by 1 from the right, by reducing its length

}
