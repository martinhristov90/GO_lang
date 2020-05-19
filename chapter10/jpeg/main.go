package main

// How to run cat go.png | ./main > go.jpeg

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // Important, not goign to work without it, it registers the PNG format (p.288 Golang book)
	"io"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin,os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr,"jpeg: %v\n", err)
		os.Exit(1)
	}

}

func toJPEG(in io.Reader, out io.Writer) error {
	img ,kind, err := image.Decode(in)
	if err != nil{
		return err
	} 
	opts := &jpeg.Options{
		Quality: 95,
	}

	fmt.Fprintf(os.Stderr, "input format =", kind)
	return jpeg.Encode(out,img, opts)
}