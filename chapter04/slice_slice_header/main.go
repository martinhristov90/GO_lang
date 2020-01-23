package main

import "fmt"

func main() {
	var buffer [150]byte

    slice := buffer[10:20]

    for i := 0; i < len(slice); i++ {
        slice[i] = byte(i)
    }
    fmt.Println("before", slice)
    AddOneToEachElement(slice)
    fmt.Println("after", slice)
    fmt.Println("after_buffer", buffer[10:20]) // the underlying array is the same
}

func AddOneToEachElement(slice []byte) {
    for i := range slice {
        slice[i]++
    }
}